package main

import (
	"encoding/json"
	"net/http"

	"github.com/mikelv92/urlshortener/internal/models"
	"github.com/mikelv92/urlshortener/internal/urlgen"
)

type CreateMappingBody struct {
	LongURL string
}

func (app *Application) createMappingHandler(w http.ResponseWriter, r *http.Request) {
	var createMappingBody CreateMappingBody
	err := json.NewDecoder(r.Body).Decode(&createMappingBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var shortURL string
	for {
		shortURL = urlgen.GenerateURL()
		exists, err := app.URLMappingModel.Exists(shortURL)
		if err != nil {
			app.Logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !exists {
			break
		}
	}

	longURL := createMappingBody.LongURL
	err = app.URLMappingModel.Insert(shortURL, longURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := models.URLMapping{
		LongURL:  longURL,
		ShortURL: shortURL,
	}

	json.NewEncoder(w).Encode(result)
}

func (app *Application) visitURL(w http.ResponseWriter, r *http.Request) {
	shortURLEncoding := r.PathValue("shortURLEncoding")

	urlMapping, err := app.URLMappingModel.Find(shortURLEncoding)
	if err != nil {
		if err == models.ErrNoRecord {
			http.Error(w, models.ErrNoRecord.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, urlMapping.LongURL, http.StatusSeeOther)
}
