package restful

import "github.com/go-chi/chi/v5"

type RouteHandler interface {
	AttachOn(chi.Router)
}
