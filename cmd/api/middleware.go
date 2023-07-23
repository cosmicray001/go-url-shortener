package main

import (
	"net/http"
	"time"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Now().Sub(startTime)
		app.infoLog.Printf("%s - %s %s %s time takes: %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI(), duration)
	})
}
