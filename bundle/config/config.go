package configfx

import (
	"go.uber.org/fx"
	"golang.org/x/oauth2"

	"scaffold/pkg/config"
	"scaffold/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(config.Load),
	fx.Provide(NewAnnotatedConfig),
)

type Result struct {
	fx.Out

	FxVerbose   bool
	Logger      logger.Config
	ServerAddr  string         `name:"server_addr"`
	GoogleOAuth *oauth2.Config `name:"google_oauth"`
	GithubOAuth *oauth2.Config `name:"github_oauth"`
}

func NewAnnotatedConfig(cfg *config.Config) Result {
	return Result{
		FxVerbose:   cfg.FxVerbose,
		Logger:      cfg.Logger,
		ServerAddr:  cfg.ServerAddr,
		GoogleOAuth: cfg.GoogleOAuth,
		GithubOAuth: cfg.GithubOAuth,
	}
}
