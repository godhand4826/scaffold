package bundlefx

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	"scaffold/pkg/config"
	"scaffold/pkg/jwt"
	"scaffold/pkg/pg"
	"scaffold/pkg/restful"
	"scaffold/src/auth"
	"scaffold/src/example"
	"scaffold/src/oauth"
	"scaffold/src/oauth/github"
	"scaffold/src/oauth/google"
	"scaffold/src/oauth/repo"
	"scaffold/src/web"
)

const GroupRouterHandler = `group:"route_handler"`

var Bundle = fx.Options(
	// config
	fx.Provide(config.Load),
	fx.Provide(NewAnnotatedConfig),

	// logger
	fx.Provide(NewLogger),
	fx.WithLogger(NewFxLogger),

	// database
	fx.Provide(pg.New),

	// repository
	fx.Provide(fx.Annotate(
		repo.New,
		fx.As(new(oauth.Repo)),
	)),

	// service
	fx.Provide(oauth.New),
	fx.Provide(jwt.NewJWTForger),

	// middleware
	fx.Provide(auth.NewMiddleware),

	// route handler
	fx.Provide(fx.Annotate(
		web.NewRouteHandler,
		fx.As(new(restful.RouteHandler)),
		fx.ResultTags(GroupRouterHandler),
	)),
	fx.Provide(fx.Annotate(
		example.NewRouteHandler,
		fx.As(new(restful.RouteHandler)),
		fx.ResultTags(GroupRouterHandler),
	)),
	fx.Provide(fx.Annotate(
		google.NewRouteHandler,
		fx.As(new(restful.RouteHandler)),
		fx.ParamTags(`name:"google_oauth"`),
		fx.ResultTags(GroupRouterHandler),
	)),
	fx.Provide(fx.Annotate(
		github.NewRouteHandler,
		fx.As(new(restful.RouteHandler)),
		fx.ParamTags(`name:"github_oauth"`),
		fx.ResultTags(GroupRouterHandler),
	)),

	// router
	fx.Provide(fx.Annotate(
		chi.NewRouter,
		fx.As(new(http.Handler)),
		fx.As(new(chi.Router)),
	)),
	fx.Provide(restful.ZapToRequestLoggerAdaptor),

	// server
	fx.Provide(fx.Annotate(
		restful.New,
		fx.ParamTags(`name:"server_addr"`),
	)),

	// invoke
	fx.Invoke(restful.RegisterMiddlewares),
	fx.Invoke(restful.RegisterMetricsHandler),
	fx.Invoke(restful.RegisterHealthCheckHandler),
	fx.Invoke(fx.Annotate(
		RegisterHandlers,
		fx.ParamTags(GroupRouterHandler),
	)),
	fx.Invoke(StartPostgres),
	fx.Invoke(StartServer),
)
