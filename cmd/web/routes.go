package main

import (
	"bookings-udemy/pkg/config"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	return mux
}
