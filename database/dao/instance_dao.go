package dao

import (
	"fmt"
	"oracleInstance/database/model"
	"oracleInstance/util/pager"
	"reflect"
	"xorm.io/xorm"
)

func OracleInstanceList(session *xorm.Session) ([]*model.OracleInstance, error) {
	session.Table("tb_oracle_instance")

	var instanceList []*model.OracleInstance

	if err := session.Find(&instanceList); err != nil {
		return nil, err
	}

	if instanceList == nil {
		return make([]*model.OracleInstance, 0), nil
	} else {
		return instanceList, nil
	}
}

func FindOracleInstanceList(session *xorm.Session, instance *model.OracleInstance, size int, offset int) ([]*model.OracleInstance, error) {
	session.Table("tb_oracle_instance")

	if instance.Id != 0 {
		session.Where("tb_oracle_instance.id = ?", instance.Id)
	}

	if instance.UserName != "" {
		session.Where("tb_oracle_instance.username like ?", "%"+instance.UserName+"%")
	}

	if instance.Host != "" {
		session.Where("tb_oracle_instance.host like ?", "%"+instance.Host+"%")
	}

	if instance.Dbname != "" {
		session.Where("tb_oracle_instance.dbname like ?", "%"+instance.Dbname+"%")
	}

	session.OrderBy("tb_oracle_instance.id desc").Limit(size, offset)

	var instanceList []*model.OracleInstance

	if err := session.Find(&instanceList); err != nil {
		return nil, err
	}

	if instanceList == nil {
		return make([]*model.OracleInstance, 0), nil
	} else {
		return instanceList, nil
	}
}

func FindOracleInstanceCount(session *xorm.Session, instance *model.OracleInstance) (int64, error) {
	session.Table("tb_oracle_instance")

	if instance.Id != 0 {
		session.Where("tb_oracle_instance.id = ?", instance.Id)
	}

	if instance.UserName != "" {
		session.Where("tb_oracle_instance.username like ?", "%"+instance.UserName+"%")
	}

	if instance.Host != "" {
		session.Where("tb_oracle_instance.host like ?", "%"+instance.Host+"%")
	}

	if instance.Dbname != "" {
		session.Where("tb_oracle_instance.dbname like ?", "%"+instance.Dbname+"%")
	}

	return session.Count(&model.OracleInstance{})
}

func FindOracleInstancePager(session *xorm.Session, instance *model.OracleInstance, pager *pager.Pager) (*pager.Pager, error) {
	count, err := FindOracleInstanceCount(session, instance)
	if err != nil {
		return nil, err
	}

	pager.SetTotal(int(count))

	list, err := FindOracleInstanceList(session, instance, pager.Size, pager.Offset)
	if err != nil {
		return nil, err
	}

	pager.Data = list
	return pager, nil
}

func InsertOracleInstance(session *xorm.Session, instance *model.OracleInstance) error {
	if _, err := session.Table("tb_oracle_instance").Insert(instance); err != nil {
		return err
	}
	return nil
}

func DeleteOracleInstance(session *xorm.Session, id int64) error {
	if _, err := session.Table("tb_oracle_instance").Where("tb_oracle_instance.id = ?", id).Delete(&model.OracleInstance{}); err != nil {
		return err
	}
	return nil
}

func UpdateOracleInstance(session *xorm.Session, instance *model.OracleInstance, columns ...string) error {
	session.Table("tb_oracle_instance")
	if columns != nil {
		session.Cols(columns...)
	}
	if _, err := session.Where("tb_oracle_instance.id = ?", instance.Id).Update(instance); err != nil {
		return err
	}
	return nil
}

func HostIsExit(session *xorm.Session, instance interface{}) error {
	session.Table("tb_oracle_instance")
	// 如果当前是添加的时候判断ip是否重复
	if reflect.TypeOf(instance).String() == "*model.AddOracleInstance" {
		exist, err := session.Where("tb_oracle_instance.host = ?", instance.(model.AddOracleInstance).Host).Exist(&model.AddOracleInstance{})
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("该ip:%s已存在", instance.(model.AddOracleInstance).Host)
		}
		// 如果当前是更新的时候判断ip是否重复
	} else if reflect.TypeOf(instance).String() == "*model.UpdateOracleInstance" {
		exist, err := session.Where("tb_oracle_instance.host = ?", instance.(model.UpdateOracleInstance).Host).Where("tb_oracle_instance.id != ?", instance.(model.UpdateOracleInstance).Id).Exist(&model.AddOracleInstance{})
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("该ip:%s已存在", instance.(model.UpdateOracleInstance).Host)
		}
	}
	return nil
}

func FindOracleInstanceById(session *xorm.Session, id int64) (*model.OracleInstance, error) {
	session.Table("tb_oracle_instance")
	instance := model.OracleInstance{}
	exit, err := session.Where("tb_oracle_instance.id = ?", id).Get(&instance)
	if !exit || err != nil {
		return nil, fmt.Errorf("未找到该实例")
	}
	return &instance, nil

}
