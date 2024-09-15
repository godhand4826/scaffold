package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	_ "embed"
)

//go:embed index.html
var indexHTML []byte

type RouteHandler struct {
}

func NewRouteHandler() *RouteHandler {
	return &RouteHandler{}
}

func (h *RouteHandler) AttachOn(router chi.Router) {
	router.Get("/index", h.Index)
}

func (h *RouteHandler) Index(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(indexHTML)
}
