package middleware

import (
	"github.com/labstack/echo"
)

type JwtConfig struct {
	Secret  string
	Session string
}

func JwtMiddlewareWithConfig(conf JwtConfig) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(ctx echo.Context) error {

			return next(ctx)
		}

	}
}
