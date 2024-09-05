package bundlefx

import (
	"go.uber.org/fx"

	configfx "scaffold/bundle/config"
	handlerfx "scaffold/bundle/handler"
	loggerfx "scaffold/bundle/logger"
	restfulfx "scaffold/bundle/restful"
)

var Bundle = fx.Options(
	configfx.Module,
	loggerfx.Module,
	handlerfx.Module,
	restfulfx.Module,
)
