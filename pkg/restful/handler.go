package restful

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterHandlers(router chi.Router) {
	router.Get("/", HelloHandler)
	router.Get("/ping", PingPongHandler)
	router.Post("/echo", EchoHandler)
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("hello"))
}

func PingPongHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10MB
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	_, _ = w.Write(body)
}
