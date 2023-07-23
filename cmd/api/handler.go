package main

import (
	"errors"
	"fmt"
	"github.com/cosmicray001/go-url-shortener/internal/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	app.urlBank.UpdateHitCount(1, 1)
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
	if input.Url == "" {
		app.badRequestResponse(w, r, errors.New("url can not be empty"))
		return
	}
	shortUrl, err := app.generateShortUrl(input.Url)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	var urlBank models.UrlBank
	urlBank.ActualUrl = input.Url
	urlBank.ShortUrl = shortUrl
	err = app.urlBank.Insert(&urlBank)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusCreated, envelope{"results": urlBank}, nil)
}

func (app *application) getLongUrl(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	shortUrl := params.ByName("shortUrl")
	fmt.Fprintf(w, "short url: %s", shortUrl)
}

func (app *application) urlList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "all the url list with hit count")
}
