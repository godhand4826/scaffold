package oauthfx

import (
	"go.uber.org/fx"

	"scaffold/internal/oauth"
	"scaffold/internal/oauth/github"
	"scaffold/internal/oauth/google"
	"scaffold/internal/oauth/repo"
	"scaffold/pkg/restful"
)

var Module = fx.Options(
	fx.Provide(oauth.New),
	fx.Provide(fx.Annotate(
		repo.New,
		fx.As(new(oauth.Repo)),
	)),
	fx.Provide(fx.Annotate(
		google.NewRouteHandler,
		fx.As(new(restful.RouteHandler)),
		fx.ParamTags(`name:"google_oauth"`),
		fx.ResultTags(`group:"route_handler"`),
	)),
	fx.Provide(fx.Annotate(
		github.NewRouteHandler,
		fx.As(new(restful.RouteHandler)),
		fx.ParamTags(`name:"github_oauth"`),
		fx.ResultTags(`group:"route_handler"`),
	)),
)
