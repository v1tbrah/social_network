package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *API) newRouter() (router *chi.Mux) {
	router = chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/ping", a.ping)

	return router
}
