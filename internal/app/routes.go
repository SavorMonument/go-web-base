package app

import (
	"net/http"

	"github.com/justinas/alice"
	"generic.com/internal/util"
)

func (app *Application) filterHttpMethod(allowedMethods []string, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !util.Contains(allowedMethods, r.Method) {
			ELOG.Printf("Wrog method: %s in call to: %s", r.Method, r.URL.Path)
			app.notFound(w)
			return
		}
		// Retreive session
		next.ServeHTTP(w, r)
		// Save session
	})
}

func (app *Application) onlyGet(next http.Handler) http.Handler {
	return app.filterHttpMethod([]string{http.MethodGet}, next)
}

func (app *Application) onlyPost(next http.Handler) http.Handler {
	return app.filterHttpMethod([]string{http.MethodPost}, next)
}

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	// Static resources
	fileServer := http.FileServer(http.Dir("./ui/public"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	getWithLogin := alice.New(app.onlyGet, noSurf, app.requireLogin)
	// postWithLogin := alice.New(app.onlyPost, noSurf, app.requireLogin)

	mux.Handle("/login", OnlyGet(http.HandlerFunc(app.getUserLoginPage)))
	mux.Handle("/dologin", OnlyPost(http.HandlerFunc(app.postUserLogin)))
	mux.Handle("/dologout", OnlyGet(http.HandlerFunc(app.getLogout)))

	mux.Handle("/", getWithLogin.ThenFunc(app.getLanding))
	mux.Handle("/favicon.ico", http.HandlerFunc(app.getFavicon))

	return app.SessionManager.LoadAndSave(app.logRequests(mux))
}
