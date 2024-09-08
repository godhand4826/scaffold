package examplefx

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	"scaffold/internal/example"
)

var Example = fx.Options(
	fx.Provide(example.NewHandler),
	fx.Invoke(func(
		router chi.Router,
		exampleHandler *example.Handler,
	) {
		exampleHandler.AttachOn(router)
	}),
)
