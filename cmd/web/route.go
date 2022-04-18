package main

import (
	"net/http"

	"github.com/Philip-21/temp/pkg/config"
	"github.com/Philip-21/temp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter() //defining a router
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Use(middleware.Recoverer) // middleware.Recoverer absorb panics and prints the stack trace,panic occurs when a program cannot function
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
