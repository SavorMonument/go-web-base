package app

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func PrintError(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	ELOG.Output(2, trace)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	ELOG.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}
