package bundlefx

import (
	"go.uber.org/fx"

	authfx "scaffold/bundle/auth"
	configfx "scaffold/bundle/config"
	examplefx "scaffold/bundle/example"
	handlerfx "scaffold/bundle/handler"
	loggerfx "scaffold/bundle/logger"
	restfulfx "scaffold/bundle/restful"
)

var Bundle = fx.Options(
	configfx.Module,
	loggerfx.Module,
	handlerfx.Module,
	restfulfx.Module,

	examplefx.Example,
	authfx.Auth,
)
