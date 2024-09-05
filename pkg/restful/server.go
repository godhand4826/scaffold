package restful

import (
	"net/http"
	"time"
)

func New(handler http.Handler, addr string) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: time.Second,
	}
}
