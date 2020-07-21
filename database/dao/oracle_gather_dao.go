package dao

import (
	"oracleInstance/database/model"
	"time"
	"xorm.io/xorm"
)

func FindOracleGatherByInstanceId(session *xorm.Session, id int64) (*model.OracleGather, error) {
	session.Table("tb_oracle_gather")

	var gather model.OracleGather
	exit, err := session.Where("instance_id = ?", id).Get(&gather)
	if err != nil {
		return nil, err
	}
	if !exit {
		return nil, nil
	}
	return &gather, nil
}

func InsertOracleGather(session *xorm.Session, gather *model.OracleGather) error {
	gather.Time = time.Now().Unix()
	if _, err := session.Table("tb_oracle_gather").Insert(gather); err != nil {
		return err
	}
	return nil
}

func DeleteOracleGatherByInstanceId(session *xorm.Session, id int64) error {
	if _, err := session.Table("tb_oracle_gather").Where("tb_oracle_gather.instance_id = ?", id).Delete(&model.OracleGather{}); err != nil {
		return err
	}
	return nil
}
