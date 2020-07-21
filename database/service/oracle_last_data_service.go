package service

import (
	"oracleInstance/database/dao"
	"oracleInstance/database/model"
	"oracleInstance/util/xorm"
)

func findOracleLastDataByInstanceId(id int64) (*model.OracleLastData, error) {
	session := xorm.NewSession()
	defer session.Close()
	return dao.FindOracleLastDataByInstanceId(session, id)
}

func deleteOracleLastDataByInstanceId(id int64) error {
	session := xorm.NewSession()
	defer session.Close()
	return dao.DeleteOracleLastDataByInstanceId(session, id)
}

func insertOracleLastDataByInstanceId(lastData *model.OracleLastData) error {
	session := xorm.NewSession()
	defer session.Close()
	return dao.InsertOracleLastData(session, lastData)
}
