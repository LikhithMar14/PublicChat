package main

import (
	"log"
	"net/http"
	"time"

	"github.com/LikhithMar14/social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}
type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}
type config struct {
	addr string
	db   dbConfig
	env  string
}

func (app *application) mount() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
		})

		r.Route("/posts/{id}", func(r chi.Router) {
			r.Get("/", app.getPostHandler)
		})
	})
	return r
}
func (app *application) run(r http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      r,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Starting server on %s", srv.Addr)
	return srv.ListenAndServe()
}
