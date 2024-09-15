package restful

import (
	"net/http"

	"scaffold/pkg/log"
)

func InjectContextLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(log.WithLogger(r.Context(), getLogger(r)))

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
