package main

import (
	"net/http"

	"github.com/fellgar246/go-templates-app/pkg/config"
	"github.com/fellgar246/go-templates-app/pkg/handlers"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
