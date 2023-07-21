package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func (app *application) createShortUrl(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Url string `json:"url"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) getLongUrl(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	shortUrl := params.ByName("shortUrl")
	fmt.Fprintf(w, "short url: %s", shortUrl)
}

func (app *application) urlList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "all the url list with hit count")
}
