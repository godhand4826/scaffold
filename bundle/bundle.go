package bundlefx

import (
	"go.uber.org/fx"

	authfx "scaffold/bundle/auth"
	configfx "scaffold/bundle/config"
	examplefx "scaffold/bundle/example"
	handlerfx "scaffold/bundle/handler"
	logfx "scaffold/bundle/log"
	oauthfx "scaffold/bundle/oauth"
	restfulfx "scaffold/bundle/restful"
)

var Bundle = fx.Options(
	configfx.Module,
	logfx.Module,
	handlerfx.Module,
	restfulfx.Module,

	examplefx.Module,
	authfx.Module,
	oauthfx.Module,
)
