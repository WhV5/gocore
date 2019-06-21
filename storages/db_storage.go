package storages

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type DBStorage struct {
	engine *xorm.Engine
}

//NewStorage 创建 Storage
func NewDBStorage(driverName, ds string, showsql bool, idle, conns int) (*DBStorage, error) {
	engine, err := xorm.NewEngine(driverName, ds)

	if err != nil {
		return nil, err
	}

	if showsql {
		engine.SetLogLevel(core.LOG_DEBUG)

	} else {
		engine.SetLogLevel(core.LOG_OFF)
	}

	engine.SetMaxIdleConns(idle)

	engine.SetMaxOpenConns(conns)

	return &DBStorage{
		engine: engine,
	}, nil
}
