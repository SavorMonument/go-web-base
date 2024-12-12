package app

import (
	"net/http"
	"text/template"

	"generic.com/internal/models"
)

func (app *Application) getFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/public/img/favicon.png")
}

func (app *Application) getLanding(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/templates/landing.page.tmpl",
		"./ui/templates/base.layout.tmpl",
		"./ui/templates/footer.partial.tmpl",
		"./ui/templates/middle.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	user := app.SessionManager.Get(r.Context(), "user").(models.User)
	// ctpForms, err := app.ComputeCTPFormAvailablility(user)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	data := make(map[string]interface{})
	data["user"] = user
	err = ts.Execute(w, app.AddDefaultData(r, data))
	if err != nil {
		app.serverError(w, err)
	}
}
