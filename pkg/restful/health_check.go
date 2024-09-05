package restful

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterHealthCheckHandler(router *chi.Mux) {
	router.Get("/livez", func(_ http.ResponseWriter, _ *http.Request) {})
	router.Get("/readyz", func(_ http.ResponseWriter, _ *http.Request) {})
}
