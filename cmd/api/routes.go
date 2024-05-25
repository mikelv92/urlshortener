package main

import "net/http"

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /create", app.createMappingHandler)
	mux.HandleFunc("GET /{shortURLEncoding}", app.visitURL)

	return commonHeaders(app.LogRequest(mux))
}
