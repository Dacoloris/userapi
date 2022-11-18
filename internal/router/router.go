package router

import (
	"net/http"
	"refactoring/internal/handler"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(h *handler.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Timeout(60*time.Second),
		SetJSONHeader,
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", h.GetUsers)
				r.Post("/", h.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", h.GetUser)
					r.Patch("/", h.UpdateUserName)
					r.Delete("/", h.DeleteUser)
				})
			})
		})
	})

	return r
}

func SetJSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
