package restful

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func RegisterMiddlewares(r chi.Router, lf middleware.LogFormatter) {
	r.Use(Metrics())
	r.Use(middleware.RequestID, SetRequestIDHeader())
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestLogger(lf), InjectContextLogger())
	r.Use(middleware.Recoverer)
	r.Use(SetCorsHeader())
	r.Use(middleware.NoCache)
	r.Use(middleware.SetHeader("X-Content-Type-Options", "nosniff"))
	r.Use(middleware.SetHeader("X-Frame-Options", "deny"))
}

func SetRequestIDHeader() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(middleware.RequestIDHeader, middleware.GetReqID(r.Context()))
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func SetCorsHeader() func(http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler
}
