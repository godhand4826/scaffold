package authfx

import (
	"go.uber.org/fx"

	"scaffold/internal/auth/github"
	"scaffold/internal/auth/google"
	"scaffold/pkg/restful"
)

var Module = fx.Options(
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
