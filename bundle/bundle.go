package bundlefx

import (
	"go.uber.org/fx"

	authfx "scaffold/bundle/auth"
	configfx "scaffold/bundle/config"
	examplefx "scaffold/bundle/example"
	handlerfx "scaffold/bundle/handler"
	logfx "scaffold/bundle/log"
	oauthfx "scaffold/bundle/oauth"
	pgfx "scaffold/bundle/pg"
	restfulfx "scaffold/bundle/restful"
	webfx "scaffold/bundle/web"
)

var Bundle = fx.Options(
	configfx.Module,
	logfx.Module,
	handlerfx.Module,
	restfulfx.Module,

	pgfx.Module,
	webfx.Module,
	examplefx.Module,
	authfx.Module,
	oauthfx.Module,
)
