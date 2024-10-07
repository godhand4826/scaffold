package bundlefx

import (
	"go.uber.org/fx"
	"golang.org/x/oauth2"

	"scaffold/pkg/config"
	"scaffold/pkg/jwt"
	"scaffold/pkg/log"
	"scaffold/pkg/pg"
)

type Result struct {
	fx.Out

	FxVerbose   bool
	Log         log.Config
	ServerAddr  string `name:"server_addr"`
	Jwt         jwt.Config
	GoogleOAuth *oauth2.Config `name:"google_oauth"`
	GithubOAuth *oauth2.Config `name:"github_oauth"`
	Postgres    pg.Config
}

func NewAnnotatedConfig(cfg *config.Config) Result {
	return Result{
		FxVerbose:   cfg.FxVerbose,
		Log:         cfg.Log,
		ServerAddr:  cfg.ServerAddr,
		Jwt:         cfg.Jwt,
		GoogleOAuth: cfg.GoogleOAuth,
		GithubOAuth: cfg.GithubOAuth,
		Postgres:    cfg.Postgres,
	}
}
