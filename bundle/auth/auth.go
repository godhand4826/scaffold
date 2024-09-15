package authfx

import (
	"go.uber.org/fx"

	"scaffold/internal/auth"
	"scaffold/pkg/jwt"
)

var Module = fx.Options(
	fx.Provide(jwt.NewJWTForger),
	fx.Provide(auth.NewMiddleware),
)
