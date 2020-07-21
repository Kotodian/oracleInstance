package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"oracleInstance/database/handler"
)

func main() {
	r := gin.Default()
	handler.RegisterSchedulerHandler()
	handler.RegisterDatabaseHandler(r)
	r.Run(":9999")
	//dba.ParseTime("192.168.1.67",1521,"system","oracle","test")
}
