package context

import (
	"github.com/labstack/echo"
	"github.com/pssauron/gocore/storages"
)

type DContext struct {
	echo.Context

	//异常信息
	Error error

	//数据库引擎
	Engine *storages.DBStorage
}
