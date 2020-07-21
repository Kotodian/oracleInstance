package service

import (
	"database/sql"
	"oracleInstance/database/dao"
	"oracleInstance/database/model"
	"oracleInstance/database/model/dba"
	"oracleInstance/util/logger"
	"oracleInstance/util/pager"
	"oracleInstance/util/xorm"
)

// 分页查询
func FindInstancePager(instance *model.OracleInstance, pager *pager.Pager) (*pager.Pager, error) {
	session := xorm.NewSession()
	defer session.Close()

	pager, err := dao.FindOracleInstancePager(session, instance, pager)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	list := pager.Data.([]*model.OracleInstance)
	var instances []model.Instance
	for i, _ := range list {
		db, err := NewOracle(list[i])
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		instance := &model.Instance{}
		instances = dba.Instance(db, instance)
		list[i].Number = instances[0].Number
		list[i].HostName = instances[0].HostName
		list[i].Status = instances[0].Status
		list[i].DatabaseStatus = instances[0].DatabaseStatus
		list[i].Type = instances[0].Role
		list[i].Version = instances[0].Version
		list[i].StartTime = instances[0].StartTime.Unix()
		db.Close()
	}
	pager.Data = list
	return pager, nil
}
func FindOracleInstanceList() ([]*model.OracleInstance, error) {
	session := xorm.NewSession()
	defer session.Close()
	list, err := dao.OracleInstanceList(session)
	if err != nil {
		return nil, err
	}
	for i, _ := range list {
		db, err := NewOracle(list[i])
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		instance := &model.Instance{}
		instances := dba.Instance(db, instance)
		list[i].Number = instances[0].Number
		list[i].HostName = instances[0].HostName
		list[i].Status = instances[0].Status
		list[i].DatabaseStatus = instances[0].DatabaseStatus
		list[i].Type = instances[0].Role
		list[i].Version = instances[0].Version
		list[i].StartTime = instances[0].StartTime.Unix()
		db.Close()
	}
	return list, nil
}

// 插入
func InsertOracleInstance(instance *model.AddOracleInstance) error {
	session := xorm.NewSession()
	defer session.Close()

	orclInstance := &model.OracleInstance{
		Host:     instance.Host,
		UserName: instance.UserName,
		Password: instance.Password,
		Dbname:   instance.Dbname,
	}

	if err := dao.HostIsExit(session, instance); err != nil {
		logger.Error(err)
		return err
	}

	err := dao.InsertOracleInstance(session, orclInstance)

	if err != nil {
		logger.Error(err)
		return err
	}

	//todo: 添加定时任务
	return nil
}

//更新
func UpdateOracleInstance(instance *model.UpdateOracleInstance) error {
	session := xorm.NewSession()
	defer session.Close()

	orclInstance := &model.OracleInstance{
		Host:     instance.Host,
		UserName: instance.UserName,
		Password: instance.Password,
		Dbname:   instance.Dbname,
	}
	// 查询ip是否存在
	if err := dao.HostIsExit(session, instance); err != nil {
		logger.Error(err)
		return err
	}

	if err := dao.UpdateOracleInstance(session, orclInstance, "host", "username", "password", "dbname"); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func DeleteOracleInstance(id int64) error {
	session := xorm.NewSession()
	defer session.Close()

	if err := dao.DeleteOracleInstance(session, id); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// 根据id去查询
func FindOracleInstanceById(id int64) (*model.OracleInstance, error) {
	session := xorm.NewSession()
	defer session.Close()

	if instance, err := dao.FindOracleInstanceById(session, id); err != nil {
		logger.Error(err)
		return nil, err
	} else {
		// 连接Oracle
		db, err := NewOracle(instance)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		defer db.Close()
		ins := &model.Instance{}
		instances := dba.Instance(db, ins)
		instance.Number = instances[0].Number
		instance.HostName = instances[0].HostName
		instance.Status = instances[0].Status
		instance.DatabaseStatus = instances[0].DatabaseStatus
		instance.Type = instances[0].Role
		instance.Version = instances[0].Version
		instance.StartTime = instances[0].StartTime.Unix()

		return instance, nil
	}
}

// dualTime
func DualTime(instance *model.OracleInstance) (int64, error) {
	time, err := dba.ParseTime(instance.Host, 1521, instance.UserName, instance.Password, instance.Dbname)
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	return time, nil
}

func Osstat(instance *model.OracleInstance) ([]model.Osstat, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	os := new(model.Osstat)
	osstats := dba.Osstat(db, os)
	return osstats, nil
}

func DataBase(instance *model.OracleInstance) ([]model.Database, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	database := new(model.Database)
	databases := dba.Database(db, database)
	return databases, nil
}

func Sysstat(instance *model.OracleInstance) ([]model.SysStat, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	sysStat := new(model.SysStat)
	stats := dba.SysStat(db, sysStat)
	return stats, nil
}

func Sga(instance *model.OracleInstance) ([]model.Sga, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	sga := new(model.Sga)
	sgas := dba.Sga(db, sga)
	return sgas, nil
}

func SgaStat(instance *model.OracleInstance) ([]model.Sgastat, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	sgaStat := new(model.Sgastat)
	sgaStats := dba.SgaStat(db, sgaStat)
	return sgaStats, nil
}

func PgaStat(instance *model.OracleInstance) ([]model.PgaStat, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	pgaStat := new(model.PgaStat)
	pgaStats := dba.PgaStat(db, pgaStat)
	return pgaStats, nil
}

func Session(instance *model.OracleInstance) (*model.Session, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	session := new(model.Session)
	dba.Session(db, session)
	return session, nil
}

func SessionEvent(instance *model.OracleInstance) ([]model.SessionEvent, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	sessionEvent := new(model.SessionEvent)
	events := dba.SessionEvent(db, sessionEvent)
	return events, nil
}

func SessionEventTime(instance *model.OracleInstance) ([]model.SessionEvent, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	sessionEvent := new(model.SessionEvent)
	time := dba.SessionEventTime(db, sessionEvent)
	return time, nil
}

func Process(instance *model.OracleInstance) (*model.Process, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	process := new(model.Process)
	dba.Process(db, process)
	return process, nil
}

func FlashRecoveryArea(instance *model.OracleInstance) ([]model.FlashRecoveryArea, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	flashRecvoeryArea := new(model.FlashRecoveryArea)
	area := dba.FlashRecoveryArea(db, flashRecvoeryArea)
	return area, nil
}

func ArchivedLog(instance *model.OracleInstance) ([]model.ArchivedLog, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	archivedLog := new(model.ArchivedLog)
	logs := dba.ArchivedLog(db, archivedLog)
	return logs, nil
}

func Server(instance *model.OracleInstance) (*model.Server, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	server := new(model.Server)
	dba.Server(db, server)
	return server, nil
}

func Hits(instance *model.OracleInstance) (*model.Hits, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	hits := new(model.Hits)
	dba.Hits(db, hits)
	return hits, nil
}

func TableSpace(instance *model.OracleInstance) ([]model.TableSpace, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	tableSpace := new(model.TableSpace)
	tableSpaces := dba.Tablespace(db, tableSpace)
	return tableSpaces, nil
}

func EventWait(instance *model.OracleInstance) ([]model.EventWait, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	eventWait := new(model.EventWait)
	eventWaits := dba.EventWait(db, eventWait)
	return eventWaits, nil
}

func SqlCpu(instance *model.OracleInstance) ([]model.SqlCpu, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	sqlCpu := new(model.SqlCpu)
	sqlCpus := dba.SqlCpu(db, sqlCpu)
	return sqlCpus, nil
}

func Locked(instance *model.OracleInstance) ([]model.Lock, error) {
	db, err := NewOracle(instance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer db.Close()
	lock := new(model.Lock)
	locks := dba.Locked(db, lock)
	return locks, nil
}

func NewOracle(instance *model.OracleInstance) (*sql.DB, error) {
	db, err := dba.InitSql(instance.Host, 1521, instance.UserName, instance.Password, instance.Dbname)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return db, nil
}
