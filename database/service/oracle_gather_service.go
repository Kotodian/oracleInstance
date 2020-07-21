package service

import (
	"oracleInstance/database/dao"
	"oracleInstance/database/model"
	"oracleInstance/util/xorm"
)

func FindOracleGatherByInstanceId(id int64) (*model.OracleGather, error) {
	session := xorm.NewSession()
	defer session.Close()
	return dao.FindOracleGatherByInstanceId(session, id)
}

func deleteOracleGatherByInstanceId(id int64) error {
	session := xorm.NewSession()
	defer session.Close()
	return dao.DeleteOracleGatherByInstanceId(session, id)
}

func insertOracleGatherByInstanceId(Gather *model.OracleGather) error {
	session := xorm.NewSession()
	defer session.Close()
	return dao.InsertOracleGather(session, Gather)
}
