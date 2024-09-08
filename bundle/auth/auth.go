package authfx

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	"scaffold/internal/auth/github"
	"scaffold/internal/auth/google"
)

var Auth = fx.Options(
	fx.Provide(fx.Annotate(google.NewHandler, fx.ParamTags(`name:"google_oauth"`))),
	fx.Provide(fx.Annotate(github.NewHandler, fx.ParamTags(`name:"github_oauth"`))),
	fx.Invoke(func(
		router chi.Router,
		googleHandler *google.Handler,
		githubHandler *github.Handler,
	) {
		googleHandler.AttachOn(router)
		githubHandler.AttachOn(router)
	}),
)
