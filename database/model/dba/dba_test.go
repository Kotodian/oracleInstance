package dba

import (
	"fmt"
	"oracleInstance/database/model"
	"testing"
)

func TestTestSqlCpu(t *testing.T) {
	//cpu := new(model.SqlCpu)
	//db, err := InitSql("192.168.1.67", 1521, "system", "oracle", "test")
	//if err != nil {
	//	t.Error(err)
	//}
	//TestSqlCpu(db)
}

func TestSqlCpu(t *testing.T) {
	cpu := new(model.SqlCpu)
	db, err := InitSql("192.168.1.67", 1521, "system", "oracle", "test")
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	t.Log(SqlCpu(db, cpu))
}

func TestTablespace(t *testing.T) {
	tablespace := new(model.TableSpace)
	db, err := InitSql("192.168.1.67", 1521, "system", "oracle", "test")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	list := Tablespace(db, tablespace)
	for _, space := range list {
		t.Log(space)
		fmt.Println()
	}
}

func TestParseTime(t *testing.T) {

	t.Log(ParseTime("192.168.1.67", 1521, "system", "oracle", "test"))
}

func TestSession(t *testing.T) {
	db, err := InitSql("192.168.1.67", 1521, "system", "oracle", "test")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	session := new(model.Session)
	Session(db, session)
	t.Log(session)
}
