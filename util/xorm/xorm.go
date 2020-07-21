package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	"oracleInstance/util/logger"
	"time"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func init() {
	dataSource := "root:7kK@164173520@tcp(175.24.36.109:3306)/oracle?charset=utf8"
	engine, err := newEngine(dataSource)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(false)
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(50)
	engine.SetConnMaxLifetime(time.Duration(300) * time.Second)

	//err = engine.Sync2(new(model.OracleInstance),new(model.OracleGather),new(model.OracleLastData))
	//if err != nil {
	//	panic(err)
	//}

}

func newEngine(dataSourceName string) (*xorm.Engine, error) {
	var err error
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return engine, nil
}

func Engine() *xorm.Engine {
	return engine
}

func NewSession() *xorm.Session {
	return Engine().NewSession()
}
