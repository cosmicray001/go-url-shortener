package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routers() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/ping", app.ping)

	return app.logRequest(router)
}
