package middleware

import (
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

//RBAC 菜单权限控制
type RBACMiddlewareConf struct {
	Engine *xorm.Engine
}

//RBACMiddlewareWithConf 权限控制器
func RBACMiddlewareWithConf(conf *RBACMiddlewareConf) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(ctx echo.Context) error {

			return nil
		}
	}
}
