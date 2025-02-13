package main

import (
    "net/http"
)

func (app *application) routes() *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
    router.HandleFunc("POST /v1/movies", app.createMovieHandler)
    router.HandleFunc("GET /v1/movies/{id}", app.showMovieHandler)

    return router
}
