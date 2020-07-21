package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"oracleInstance/database/model"
	"oracleInstance/database/service"
	"oracleInstance/util/ginutil"
	"oracleInstance/util/pager"
)

var task *cron.Cron

func RegisterDatabaseHandler(router *gin.Engine) {
	dataBaseHandler := router.Group("/api/v1/database")
	//dataBaseHandler.GET("", findOracleInstancePager)
	dataBaseHandler.GET("", oracleInstanceList)
	dataBaseHandler.POST("", insertOracleInstance)
	dataBaseHandler.GET("/:id", findOracleInstanceById)
	dataBaseHandler.DELETE("", deleteOracleInstance)

	dataBaseInfoHandler := router.Group("/api/v1/database/:id")
	dataBaseInfoHandler.GET("/osstat", osstat)
	dataBaseInfoHandler.GET("/dualTime", dualTime)
	dataBaseInfoHandler.GET("/database", database)
	dataBaseInfoHandler.GET("/sysstat", sysstat)
	dataBaseInfoHandler.GET("/sga", sga)
	dataBaseInfoHandler.GET("/sgaStat", sgaStat)
	dataBaseInfoHandler.GET("/pgaStat", pgaStat)
	dataBaseInfoHandler.GET("/session", session)
	dataBaseInfoHandler.GET("/sessionEvent", sessionEvent)
	dataBaseInfoHandler.GET("/sessionEventTime", sessionEventTime)
	dataBaseInfoHandler.GET("/process", process)
	dataBaseInfoHandler.GET("/flashRecoveryArea", flashRecoveryArea)
	dataBaseInfoHandler.GET("/archivedLog", archivedLog)
	dataBaseInfoHandler.GET("/server", server)
	dataBaseInfoHandler.GET("/hits", hits)
	dataBaseInfoHandler.GET("/tableSpace", tableSpace)
	dataBaseInfoHandler.GET("/eventWait", eventWait)
	dataBaseInfoHandler.GET("/sqlCpu", sqlCpu)
	dataBaseInfoHandler.GET("/locked", locked)
	dataBaseInfoHandler.GET("/gather", oracleGather)
}

func init() {
	task = service.NewCron()
}

func RegisterSchedulerHandler() {
	service.StartJob(task)
}

func findOracleInstancePager(c *gin.Context) {
	pager := pager.NewPager(ginutil.GetIntQuery(c, "page", 1), ginutil.GetIntQuery(c, "size", 10))
	host := ginutil.GetQuery(c, "host", "")
	username := ginutil.GetQuery(c, "username", "")
	password := ginutil.GetQuery(c, "passwowrd", "")
	dbname := ginutil.GetQuery(c, "dbname", "")

	orclInstance := &model.OracleInstance{
		Host:     host,
		UserName: username,
		Password: password,
		Dbname:   dbname,
	}
	instancePager, err := service.FindInstancePager(orclInstance, pager)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, instancePager)
}

func findOracleInstanceById(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		c.JSON(200, instance)
	}
}

func insertOracleInstance(c *gin.Context) {
	instance := new(model.AddOracleInstance)
	c.ShouldBindJSON(instance)
	service.StopJob(task)
	err := service.InsertOracleInstance(instance)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		service.StartJob(task)
		c.JSON(200, "ok")
	}
}

func deleteOracleInstance(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	service.StopJob(task)
	err := service.DeleteOracleInstance(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		service.StartJob(task)
		c.JSON(200, "ok")
	}
}

func dualTime(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.DualTime(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func osstat(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.Osstat(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func sysstat(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.Sysstat(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func database(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.DataBase(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func sga(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.Sga(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func sgaStat(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.SgaStat(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func pgaStat(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.PgaStat(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func session(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.Session(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func sessionEvent(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.SessionEvent(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func sessionEventTime(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.SessionEventTime(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func process(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.Process(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func flashRecoveryArea(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.FlashRecoveryArea(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func archivedLog(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.ArchivedLog(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func server(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.Server(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func hits(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.Hits(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func tableSpace(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.TableSpace(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}

}

func eventWait(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.EventWait(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func sqlCpu(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.SqlCpu(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func locked(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.Locked(instance)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func oracleGather(c *gin.Context) {
	id := ginutil.GetInt64Param(c, "id", 0)
	instance, err := service.FindOracleInstanceById(id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		data, err := service.FindOracleGatherByInstanceId(instance.Id)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, data)
		}
	}
}

func oracleInstanceList(c *gin.Context) {
	instances, err := service.FindOracleInstanceList()
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		c.JSON(200, instances)
	}

}
