package app

import (
	"net/http"

	"github.com/justinas/nosurf"
	"generic.com/internal/models"
	"generic.com/internal/util"
)

func (app *Application) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, ok := app.SessionManager.Get(r.Context(), "user").(*models.User); ok {
			LOG.Printf("%s - %s - %s %s %s", r.RemoteAddr, user.Username, r.Proto, r.Method, r.URL)
		} else {
			LOG.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL)
		}
		next.ServeHTTP(w, r)
	})
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})
	csrfHandler.SetFailureHandler(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ELOG.Printf("CSRF check failed: %s", nosurf.Reason(r))
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}))
	return csrfHandler
}

func (app *Application) requireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, ok := app.SessionManager.Get(r.Context(), "user").(models.User)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func FilterHttpMethod(allowedMethods []string, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !util.Contains(allowedMethods, r.Method) {
			ELOG.Printf("Wrog method: %s in call to: %s", r.Method, r.URL.Path)
			NotFound(w)
			return
		}
		// Retreive session
		next.ServeHTTP(w, r)
		// Save session
	})
}

func OnlyGet(next http.Handler) http.Handler {
	return FilterHttpMethod([]string{http.MethodGet}, next)
}

func OnlyPost(next http.Handler) http.Handler {
	return FilterHttpMethod([]string{http.MethodPost}, next)
}

func OnlyDelete(next http.Handler) http.Handler {
	return FilterHttpMethod([]string{http.MethodDelete}, next)
}

// func (app *Application) filterRole(allowedRoles []models.Role, next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		user := app.SessionManager.Get(r.Context(), "user").(*models.User)

// 		if util.Contains(allowedRoles, user.Role) {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		ELOG.Printf("User: %s does not have the required roles: %v", user.Username, allowedRoles)
// 		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
// 	})
// }
