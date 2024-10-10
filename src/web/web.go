package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	_ "embed"
)

var (
	//go:embed index.html
	indexHTML []byte

	//go:embed callback.html
	callbackHTML []byte
)

type RouteHandler struct {
}

func NewRouteHandler() *RouteHandler {
	return &RouteHandler{}
}

func (h *RouteHandler) AttachOn(router chi.Router) {
	router.Get("/index", h.Index)
	router.Get("/v1/oauth/{provider}/callback", h.OAuthCallback)
}

func (h *RouteHandler) Index(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(indexHTML)
}

func (h *RouteHandler) OAuthCallback(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(callbackHTML)
}
