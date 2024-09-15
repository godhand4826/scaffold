package auth

import (
	"net/http"
	"strings"

	"scaffold/pkg/jwt"
)

type Middleware struct {
	forger *jwt.Forger
}

func NewMiddleware(forger *jwt.Forger) *Middleware {
	return &Middleware{
		forger: forger,
	}
}

func (h *Middleware) Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			subject, err := h.forger.Verify(token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			// TODO inject to context
			_ = subject

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
