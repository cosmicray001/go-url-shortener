package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routers() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/api/ping", app.ping)
	router.HandlerFunc(http.MethodPost, "/api/shorten", app.createShortUrl)
	router.HandlerFunc(http.MethodGet, "/api/decode/:shortUrl", app.getLongUrl)
	router.HandlerFunc(http.MethodGet, "/api/stats", app.urlList)

	return app.logRequest(router)
}
