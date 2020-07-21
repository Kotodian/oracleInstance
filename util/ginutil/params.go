package ginutil

import (
	"github.com/gin-gonic/gin"
	"oracleInstance/util/logger"
	"strconv"
)

func GetQuery(c *gin.Context, key string, defaultVal string) string {
	value, ext := c.GetQuery(key)
	if !ext {
		return defaultVal
	}
	return value
}

func GetIntQuery(c *gin.Context, key string, defaultVal int) int {
	value, ext := c.GetQuery(key)
	if !ext {
		return defaultVal
	}
	atoi, err := strconv.Atoi(value)
	if err != nil {
		logger.Error(err)
		return defaultVal
	}
	return atoi
}

func GetInt64Query(c *gin.Context, key string, defaultVal int64) int64 {
	value, ext := c.GetQuery(key)
	if !ext {
		return defaultVal
	}
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		logger.Error(err)
		return defaultVal
	}
	return v
}

func GetInt64Param(c *gin.Context, key string, defaultVal int64) int64 {
	param := c.Param(key)
	if len(param) == 0 {
		return defaultVal
	}
	v, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Error(err)
		return defaultVal
	}
	return v
}
