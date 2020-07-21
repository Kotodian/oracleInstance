package service

import (
	"github.com/robfig/cron/v3"
	"oracleInstance/database/model"
	"oracleInstance/database/model/dba"
	"oracleInstance/util/logger"
)

func NewCron() *cron.Cron {
	return cron.New()
}

func StartJob(c *cron.Cron) {
	//c.AddFunc("0 */3 * * * *",oracleGather)
	entryID, err := c.AddFunc("@every 0h0m50s", OracleGather)
	if err != nil {
		panic(err)
	} else {
		logger.Info(entryID)
	}
	c.Start()
}

func StopJob(c *cron.Cron) {
	c.Stop()
}

func OracleGather() {
	instances, err := FindOracleInstanceList()
	if err != nil {
		logger.Error(err)
		return
	}
	if len(instances) == 0 {
		logger.Info("目前无实例")
		return
	}
	for _, ins := range instances {
		sysstat, err := Sysstat(ins)
		if err != nil {
			logger.Error(err)
			return
		}
		pgaStat, err := PgaStat(ins)
		if err != nil {
			logger.Error(err)
			return
		}
		sgaStat, err := SgaStat(ins)
		if err != nil {
			logger.Error(err)
			return
		}
		hits, err := Hits(ins)
		if err != nil {
			logger.Error(err)
			return
		}
		session, err := Session(ins)
		if err != nil {
			logger.Error(err)
			return
		}
		lastData, err := findOracleLastDataByInstanceId(ins.Id)
		if err != nil {
			logger.Error(err)
			return
		}
		item := new(model.OracleGather)
		sysStatMap := dba.SysStatTurnToMap(sysstat)
		pgaStatMap := dba.PgaStatTurnToMap(pgaStat)
		sgaStatMap := dba.SgaStatTurnToMap(sgaStat)
		if err != nil {
			logger.Error(err)
			return
		}
		if lastData != nil {
			//先删除数据库里的再添加
			item.Iops = (sysStatMap["physical writes"] + sysStatMap["physical reads"]) - lastData.Iops
			item.Mbps = (sysStatMap["physical read bytes"]+sysStatMap["physical write bytes"])*1.0/1024/1024 - lastData.Mbps
			item.BytesReceivedViaSQLNetFromClient = sysStatMap["bytes received via SQL*Net from client"] - lastData.BytesReceivedViaSQLNetFromClient
			item.ParseCountHard = sysStatMap["parse count (hard)"] - lastData.ParseCountHard
			item.SqlNetRoundtripsTfromClient = sysStatMap["SQL*Net roundtrips to/from client"] - lastData.SqlNetRoundTripsFromClient
			item.ExecuteCount = int(sysStatMap["execute count"]) - lastData.ExecuteCount
		} else {
			lastData = new(model.OracleLastData)
			item.Iops = 0
			item.Mbps = 0
			item.BytesReceivedViaSQLNetFromClient = 0
			item.BytesSentViaSQLNetToClient = 0
			item.ParseCountHard = 0
			item.SqlNetRoundtripsTfromClient = 0
			item.ExecuteCount = 0
		}
		item.DualTime, err = dba.ParseTime(ins.Host, 1521, ins.UserName, ins.Password, ins.Dbname)
		if err != nil {
			logger.Error(err)
			return
		}
		item.SortMemory = sysStatMap["sorts (memory)"] / (sysStatMap["sorts (memory)"] + sysStatMap["sorts (disk)"])
		item.UseTotalPGA = pgaStatMap["total PGA inuse"]
		item.SqlPinHitRatio = hits.PinHit
		item.BufferHit = hits.BufferCache
		item.RedoBufferAllocationRetries = hits.RedoBuffer
		item.BackgroundSessions = session.UserBackGroundSessions
		item.UserInactiveSessions = session.UserInactiveSessions
		item.UserActiveSessions = session.UserActiveSessions
		item.SharePoolSize = sgaStatMap["shared pool"]
		item.InstanceId = ins.Id

		// 保存本次的数据
		lastData.InstanceId = ins.Id
		lastData.Iops = sysStatMap["physical writes"] + sysStatMap["physical reads"]
		lastData.Mbps = (sysStatMap["physical read bytes"] + sysStatMap["physical write bytes"]) * 1.0 / 1024 / 1024
		lastData.BytesReceivedViaSQLNetFromClient = sysStatMap["bytes received via SQL*Net from client"]
		lastData.ParseCountHard = sysStatMap["bytes sent via SQL*Net to client"]
		lastData.SqlNetRoundTripsFromClient = sysStatMap["SQL*Net roundtrips to/from client"]
		lastData.ExecuteCount = int(sysStatMap["execute count"])
		err = deleteOracleLastDataByInstanceId(ins.Id)
		if err != nil {
			logger.Error(err)
			return
		}
		err = deleteOracleGatherByInstanceId(ins.Id)
		if err != nil {
			logger.Error(err)
			return
		}
		err = insertOracleLastDataByInstanceId(lastData)
		if err != nil {
			logger.Error(err)
			return
		}

		err = insertOracleGatherByInstanceId(item)
		if err != nil {
			logger.Error(err)
		}
	}
	logger.Info("执行成功")
}
