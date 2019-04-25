package middleware

import (
	"github.com/labstack/echo"
	"github.com/pssauron/gocore/jwt"
	"github.com/pssauron/gocore/thc/errs"
)

type (
	//jwt 配置文件
	JwtConfig struct {
		Secret   string
		Session  string
		AuthFrom string //default:cookie
	}

	Extractor func(echo.Context) string
)

func JwtMiddlewareWithConfig(conf JwtConfig) echo.MiddlewareFunc {

	if conf.Secret == "" {
		conf.Secret = "pssauron"
	}

	if conf.Session == "" {
		conf.Session = "token"
	}

	if conf.AuthFrom == "" {
		conf.AuthFrom = "cookie"
	}

	extractor := SessionExtractorFromCookie()

	switch conf.AuthFrom {
	case "query":
		extractor = SessionExtractorFromQuery()
	case "header":
		extractor = SessionExtractorFromHeader()

	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(ctx echo.Context) error {

			token := extractor(ctx)

			if token == "" {
				return errs.ServiceError{ErrCode: "100401", ErrMessage: "用户未授权"}
			}

			data, err := jwt.ValidateToken(conf.Secret, token)

			if err != nil {
				return errs.ServiceError{ErrCode: "100401", ErrMessage: "用户未授权或授权已过去"}
			}

			ctx.Set(conf.Session, data)

			return next(ctx)
		}

	}

}

func SessionExtractorFromQuery() Extractor {
	return func(ctx echo.Context) string {
		return ctx.Param("token")
	}

}

func SessionExtractorFromHeader() Extractor {

	return func(ctx echo.Context) string {
		return ctx.Request().Header.Get(echo.HeaderAuthorization)
	}
}

func SessionExtractorFromCookie() Extractor {

	return func(ctx echo.Context) string {
		c, err := ctx.Cookie("token")
		if err != nil {
			return ""
		}

		return c.Value
	}
}
