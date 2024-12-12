package app

import (
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/schema"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"generic.com/internal/models"
)

var decoder = schema.NewDecoder()

func (app *Application) getUserLoginPage(w http.ResponseWriter, r *http.Request) {

	if _, ok := app.SessionManager.Get(r.Context(), "user").(*models.User); ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	files := []string{
		"./ui/templates/login/login.page.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		ServerError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		ServerError(w, err)
	}
}

func (app *Application) postUserLogin(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		ServerError(w, err)
		return
	}

	username := r.PostForm.Get("username")
	pass := r.PostForm.Get("password")

	if user, err := app.Repos.UserRepo.FindByUsername(username); err != nil {
		if err == gorm.ErrRecordNotFound {
			ELOG.Printf("User tried to login with non-existing username: %s", username)
			app.SessionManager.Put(r.Context(), "flash", "Invalid credentials")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ServerError(w, err)
		return
	} else if err := bcrypt.CompareHashAndPassword(user.Password, []byte(pass)); err != nil {
		ELOG.Printf("User tried to login with invalid password: %s", username)

		// To prevent timing attacks, we sleep for a random amount of time before redirecting
		jitter := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(5*time.Second + jitter)

		app.SessionManager.Put(r.Context(), "flash", "Invalid credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else {
		app.SessionManager.Put(r.Context(), "user", user)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) getLogout(w http.ResponseWriter, r *http.Request) {
	app.SessionManager.Remove(r.Context(), "user")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
