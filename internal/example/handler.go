package example

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AttachOn(router chi.Router) {
	router.Get("/", h.HelloHandler)
	router.Get("/ping", h.PingPongHandler)
	router.Post("/echo", h.EchoHandler)
}

func (*Handler) HelloHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("hello"))
}

func (*Handler) PingPongHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func (*Handler) EchoHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10MB
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	_, _ = w.Write(body)
}
