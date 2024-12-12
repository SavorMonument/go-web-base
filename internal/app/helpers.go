package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

var LOG = log.New(os.Stdout, "APP.INFO\t", log.Ldate|log.Ltime)
var ELOG = log.New(os.Stderr, "APP.ERROR\t", log.Ldate|log.Ltime)

func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	ELOG.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
