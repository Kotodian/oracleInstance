package dao

import (
	"oracleInstance/database/model"
	"xorm.io/xorm"
)

func FindOracleLastDataByInstanceId(session *xorm.Session, id int64) (*model.OracleLastData, error) {
	session.Table("tb_oracle_last_data")

	var lastData model.OracleLastData
	exit, err := session.Where("instance_id = ?", id).Get(&lastData)
	if err != nil {
		return nil, err
	}
	if !exit {
		return nil, nil
	}
	return &lastData, nil
}

func InsertOracleLastData(session *xorm.Session, lastdata *model.OracleLastData) error {
	if _, err := session.Table("tb_oracle_last_data").Insert(lastdata); err != nil {
		return err
	}
	return nil
}

func DeleteOracleLastDataByInstanceId(session *xorm.Session, id int64) error {
	if _, err := session.Table("tb_oracle_last_data").Where("tb_oracle_last_data.instance_id = ?", id).Delete(&model.OracleLastData{}); err != nil {
		return err
	}
	return nil
}
