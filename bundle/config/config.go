package configfx

import (
	"go.uber.org/fx"

	"scaffold/pkg/config"
	"scaffold/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(config.Load),
	fx.Provide(NewAnnotatedConfig),
)

type Result struct {
	fx.Out

	FxVerbose  bool
	Logger     logger.Config
	ServerAddr string `name:"server_addr"`
}

func NewAnnotatedConfig(cfg *config.Config) Result {
	return Result{
		FxVerbose:  cfg.FxVerbose,
		Logger:     cfg.Logger,
		ServerAddr: cfg.ServerAddr,
	}
}
