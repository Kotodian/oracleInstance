package dba

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"oracleInstance/database/model"
	"oracleInstance/util/logger"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var db *sql.DB

//const  (
//	host  = "192.168.1.67"
//	port = 1521
//	username = "system"
//	password = "oracle"
//	dbname = "test"
//)

//远程连接oracle
func InitSql(host string, port int, username, password, dbname string) (*sql.DB, error) {
	osqlInfo := fmt.Sprintf("%s/%s@%s:%d/%s", username, password, host, port, dbname)
	DB, err := sql.Open("godror", osqlInfo)
	if err != nil {
		return nil, err
	}
	err = DB.Ping()
	if err != nil {
		return nil, err
	}
	db = DB
	return db, nil
}

// 数据库连接实例信息
func Instance(db *sql.DB, value *model.Instance) []model.Instance {
	var res []model.Instance
	stmt, err := db.Prepare("select  * from v$instance")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	//fmt.Println(rows.Columns())
	defer stmt.Close()
	for rows.Next() {
		//fmt.Println(rows)
		err := rows.Scan(&value.Number, &value.Name,
			&value.HostName, &value.Version, &value.StartTime,
			&value.Status, &value.Parallel, &value.Thread, &value.Archiver, &value.LogSwitchWait,
			&value.Logins, &value.ShutDownPending, &value.DatabaseStatus,
			&value.Role, &value.ActiveState, &value.Blocked)

		if err != nil {
			panic(err)
		}
		res = append(res, *value)
		//fmt.Println(value)
	}
	return res
}

// 数据库系统信息
func Osstat(db *sql.DB, os *model.Osstat) []model.Osstat {
	osList := []model.Osstat{}
	stmt, err := db.Prepare("select  * from v$osstat")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()

	if err != nil {
		panic(err)
	}
	//fmt.Println(rows.Columns())

	defer stmt.Close()

	for rows.Next() {
		err := rows.Scan(&os.StatName, &os.Value, &os.Id, &os.Comments, &os.Cumulative)
		if err != nil {
			panic(err)
		}
		//fmt.Println(os)
		osList = append(osList, *os)
	}
	return osList
}

//数据库信息
func Database(db *sql.DB, database *model.Database) []model.Database {
	var databases []model.Database
	stmt, err := db.Prepare("select name,platform_name,created from v$database")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	//fmt.Println(rows.Columns())
	for rows.Next() {
		//fmt.Println(rows)
		rows.Scan(&database.Name, &database.PlatformName, &database.CreateTime)
		databases = append(databases, *database)
	}

	return databases
}

// 数据库实例pga信息
func PgaStat(db *sql.DB, pgaStat *model.PgaStat) []model.PgaStat {
	pgaStatList := []model.PgaStat{}
	stmt, err := db.Prepare("select  * from v$pgastat")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rows.Scan(&pgaStat.Name, &pgaStat.Value, &pgaStat.Unit)
		pgaStatList = append(pgaStatList, *pgaStat)
	}
	return pgaStatList
	//fmt.Println(rows.Columns())
}

// sga是空的
// 数据库实例 sga 信息
// sga 所有服务进程和后台进程共享
func Sga(db *sql.DB, sga *model.Sga) []model.Sga {
	var sgaList []model.Sga
	stmt, err := db.Prepare("select  * from v$sga")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	fmt.Println(rows.Columns())
	defer stmt.Close()
	for rows.Next() {
		rows.Scan(&sga.Name, &sga.Value)
		sgaList = append(sgaList, *sga)
	}
	return sgaList
}

func SgaStat(db *sql.DB, sgaStat *model.Sgastat) []model.Sgastat {
	var sgaStatList []model.Sgastat
	stmt, err := db.Prepare("select  pool,sum(bytes) from v$sgastat where pool in ('java pool','large pool','shared pool') group by pool")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for rows.Next() {
		err := rows.Scan(&sgaStat.Pool, &sgaStat.SumBytes)
		if err != nil {
			panic(err)
		}
		sgaStatList = append(sgaStatList, *sgaStat)
	}
	return sgaStatList
	//fmt.Println(rows.Columns())
}

// 数据库主要的性能指标
func SysStat(db *sql.DB, sys *model.SysStat) []model.SysStat {
	var sysList []model.SysStat
	stmt, err := db.Prepare("select  * from v$sysstat")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	//fmt.Println(rows.Columns())
	defer stmt.Close()

	for rows.Next() {
		err := rows.Scan(&sys.Statistic, &sys.Name, &sys.Class, &sys.Value, &sys.Id)
		if err != nil {
			panic(err)
		}
		sysList = append(sysList, *sys)
	}
	return sysList
}

// 数据库dispatcher繁忙度
func Dispatcher(db *sql.DB, dis *model.Dispatcher) []model.Dispatcher {
	var disList []model.Dispatcher
	stmt, err := db.Prepare("select busy,idle from v$dispatcher")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	//fmt.Println(rows.Columns())
	defer stmt.Close()
	for rows.Next() {
		rows.Scan(&dis.Idle, &dis.Busy)
		disList = append(disList, *dis)
		//fmt.Println(rows)
	}
	return disList
}

// 数据库事件信息
func SessionEvent(db *sql.DB, event *model.SessionEvent) []model.SessionEvent {
	var eventList []model.SessionEvent
	stmt, err := db.Prepare("select event, count(*) as value from v$session_event where event in ('log file sync', 'db file scattered read', 'db file sequential read', 'direct path read', 'direct path write') group by event")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	//fmt.Println(rows.Columns())
	for rows.Next() {
		rows.Scan(&event.Event, &event.Value)
		eventList = append(eventList, *event)
	}
	return eventList
}

// 数据库会话事件信息
func SessionEventTime(db *sql.DB, event *model.SessionEvent) []model.SessionEvent {
	var eventList []model.SessionEvent
	stmt, err := db.Prepare("select event, sum(average_wait) as value from v$session_event where event in ('log file sync', 'db file scattered read', 'db file sequential read', 'direct path read', 'direct path write') group by event")
	if err != nil {
		panic(err)
	}
	//varvaluestring
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	//fmt.Println(rows.Columns())
	defer stmt.Close()

	for rows.Next() {
		rows.Scan(&event.Event, &event.Value)
		eventList = append(eventList, *event)
	}
	return eventList
}

// 数据库会话
func Session(db *sql.DB, session *model.Session) {
	db.QueryRow("select count(*) from gv$session where status = 'ACTIVE'").Scan(&session.UserActiveSessions)
	db.QueryRow("select count(*) from gv$session where status = 'INACTIVE'").Scan(&session.UserInactiveSessions)
	db.QueryRow("select count(*) from gv$session where status != 'ACTIVE' and status != 'INACTIVE'").Scan(&session.UserBackGroundSessions)
}

// 数据库进程
func Process(db *sql.DB, process *model.Process) {
	db.QueryRow("select count(*) from v$process").Scan(&process.Process)
	db.QueryRow("select count(*) from v$bgprocess").Scan(&process.BackgroundProcess)
	db.QueryRow("select count(*) from v$bgprocess where paddr <> ? and name like ?", "00", "LGW%").Scan(&process.Redo)
}

// 数据库空的
// 数据库闪回区使用情况
func FlashRecoveryArea(db *sql.DB, area *model.FlashRecoveryArea) []model.FlashRecoveryArea {
	var areaList []model.FlashRecoveryArea
	stmt, err := db.Prepare("select  * from v$flash_recovery_area_usage")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rows.Scan(&area.FileType, &area.PercentSpaceUsed, &area.PercentSpaceReclaimable, &area.NumberOfFiles)
		areaList = append(areaList, *area)
	}
	return areaList
}

//数据库归档日志
func ArchivedLog(db *sql.DB, arc *model.ArchivedLog) []model.ArchivedLog {
	var arcList []model.ArchivedLog
	stmt, err := db.Prepare("select  substr(t.NAME,1) name, ROUND(sum(t.BLOCKS*t.BLOCK_SIZE)/1024/1024) total from v$archived_log t where t.DELETED='NO' group by substr(t.NAME,1)")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()

	if err != nil {
		panic(err)
	}
	fmt.Println(rows.Columns())
	for rows.Next() {
		rows.Scan(&arc.Name, &arc.Total)
		arcList = append(arcList, *arc)
	}
	return arcList
}

//数据库服务信息
func Server(db *sql.DB, server *model.Server) {
	db.QueryRow("select count(Server) from v$session where Server = ?", "DEDICATED").Scan(&server.Dedicated)
	db.QueryRow("select count(Server) from v$session where Server = ?", "SHARED").Scan(&server.Shared)
	db.QueryRow("select count(Server) from v$session").Scan(&server.Total)
}

// sql解析时间
func ParseTime(host string, port int, username, password, dbname string) (int64, error) {
	dbPath := fmt.Sprintf("%s/%s@%s:%d/%s", username, password, host, port, dbname)
	currentPath, _ := os.Getwd()
	base_dir := projectPath(currentPath)
	sqlPath := fmt.Sprintf("@%sdual.sql", base_dir)
	//sqlPath := "@dual.sql"
	//fmt.Println(sqlPath)
	cmd := exec.Command("sqlplus", "-S", dbPath, sqlPath)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	//fmt.Println(out.String())
	temp := strings.Split(out.String(), ":")
	temp2 := strings.Split(temp[len(temp)-1], ".")
	res := strings.Split(temp2[len(temp2)-1], "\r\n")
	r, err := strconv.ParseInt(res[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	return r, nil
}

//命中率
func Hits(db *sql.DB, hits *model.Hits) {
	db.QueryRow("select  100*round(1-((physical.value - direct.value - lobs.value) / logical.value),4) \"Buffer Cache Hit Ratio\" from v$sysstat physical,v$sysstat direct,v$sysstat lobs,v$sysstat logical where physical.name = 'physical reads' and direct.name = 'physical reads direct' and lobs.name = 'physical reads direct (lob)' and logical.name = 'session logical reads'").Scan(&hits.BufferCache)
	db.QueryRow("select 100*round(s2.value/(select sum(s1.value) from v$sysstat s1 where name in ('redo buffer allocation retries','redo entries')),2) FAILING_PER from v$sysstat s2 where s2.name = 'redo buffer allocation retries'").Scan(&hits.RedoBuffer)
	db.QueryRow("select round(((sum(pinhits) / sum(pins)) * 100),2)\"PinHitRatio\" from v$librarycache ").Scan(&hits.PinHit)
}

// 数据库表空间
func Tablespace(db *sql.DB, space *model.TableSpace) []model.TableSpace {
	var spaceList []model.TableSpace
	stmt, err := db.Prepare("select b.file_id,b.tablespace_name,b.file_name,b.AUTOEXTENSIBLE, ROUND(b.bytes/1024/1024,2), ROUND((b.bytes-sum(nvl(a.bytes,0)))/1024/1024,2), ROUND(sum(nvl(a.bytes,0))/1024/1024,2), ROUND((b.bytes-sum(nvl(a.bytes,0)))/(b.bytes),3)*100 from dba_free_space a,dba_data_files b where a.file_id=b.file_id group by b.tablespace_name,b.file_name,b.file_id,b.bytes,b.AUTOEXTENSIBLE order by b.tablespace_name")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(rows.Columns())
	for rows.Next() {
		rows.Scan(&space.Id, &space.Name, &space.File, &space.Auto, &space.Total, &space.Use, &space.Free, &space.Percent)
		//fmt.Println(rows)
		spaceList = append(spaceList, *space)
	}
	return spaceList
}

// 事件等待
func EventWait(db *sql.DB, event *model.EventWait) []model.EventWait {
	var eventList []model.EventWait
	stmt, err := db.Prepare("select  sid,event,total_waits,total_timeouts,time_waited,AVERAGE_WAIT,wait_class from v$session_event")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()

	if err != nil {
		panic(err)
	}
	//fmt.Println(rows.Columns())
	defer stmt.Close()

	for rows.Next() {
		err := rows.Scan(&event.Sid, &event.Event, &event.TotalWaits, &event.TotalTimeOuts, &event.TimeWaited, &event.AverageWait, &event.WaitClass)
		if err != nil {
			panic(err)
		}
		eventList = append(eventList, *event)
	}
	return eventList
}

//消耗cpu最多的sql
func SqlCpu(db *sql.DB, cpu *model.SqlCpu) []model.SqlCpu {
	var cpuList []model.SqlCpu
	stmt, err := db.Prepare("select * from (select  round(CPU_TIME/1000000,1),executions,username,PARSING_USER_ID,sql_id,round(ELAPSED_TIME/1000000,1),sql_text from v$sql,dba_users where user_id=PARSING_USER_ID order by CPU_TIME/1000000 desc) where rownum <= 30")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()

	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for rows.Next() {
		rows.Scan(&cpu.CpuTime, &cpu.Executions, &cpu.UserName, &cpu.ParsingUserID, &cpu.SqlId, &cpu.ElapsedTime, &cpu.SqlText)
		cpuList = append(cpuList, *cpu)
	}
	return cpuList
}

//func TestSqlCpu(db *sql.DB,cpu *model.SqlCpu)  []model.SqlCpu {
//	var cpuList []model.SqlCpu
//	stmt, err := db.Prepare("select * from (select  round(CPU_TIME/1000000,1),executions,username,PARSING_USER_ID,sql_id,round(ELAPSED_TIME/1000000,1),sql_text from v$sql,dba_users where user_id=PARSING_USER_ID order by CPU_TIME/1000000 desc) where rownum <= 30")
//	if err != nil {
//		panic(err)
//	}
//	//var value string
//	rows, err := stmt.Query()
//
//	if err != nil {
//		panic(err)
//	}
//	defer stmt.Close()
//	for rows.Next() {
//		fmt.Println(rows)
//	}
//	return cpuList
//}

// 查询死锁
func Locked(db *sql.DB, lock *model.Lock) []model.Lock {
	var lockList []model.Lock
	stmt, err := db.Prepare("select S.SID,s.SERIAL#,s.username,l.OBJECT_ID,l.ORACLE_USERNAME,l.OS_USER_NAME,l.PROCESS from V$LOCKED_OBJECT l,V$SESSION S where l.SESSION_ID=S.SID AND s.username IS NOT NULL ORDER BY s.SID")
	if err != nil {
		panic(err)
	}
	//var value string
	rows, err := stmt.Query()

	if err != nil {
		panic(err)
	}
	//fmt.Println(rows.Columns())
	for rows.Next() {
		err := rows.Scan(&lock.Sid, &lock.Serial, &lock.ObjectId, &lock.OracleUsername, &lock.OSUsername, &lock.Process)
		if err != nil {
			panic(err)
		}
		lockList = append(lockList, *lock)
	}
	return lockList

}

// 杀死死锁
func KillLocked(db *sql.DB, lock *model.Lock) {
	db.Exec("alter system kill session ?,?", lock.Sid, lock.Serial)
}

func SysStatTurnToMap(sysStat []model.SysStat) map[string]float64 {
	sysStatMap := make(map[string]float64)

	for _, v := range sysStat {
		sysStatMap[v.Name] = v.Value
	}
	return sysStatMap
}

func PgaStatTurnToMap(pgaStat []model.PgaStat) map[string]float64 {
	pgaStatMap := make(map[string]float64)

	for _, v := range pgaStat {
		pgaStatMap[v.Name] = v.Value
	}
	return pgaStatMap
}

func SgaStatTurnToMap(sgastat []model.Sgastat) map[string]float64 {
	sgaStatMap := make(map[string]float64)

	for _, v := range sgastat {
		sgaStatMap[v.Pool] = v.SumBytes
	}

	return sgaStatMap
}

func projectPath(dir string) string {
	current, file := filepath.Split(dir)

	if current == "" {
		return ""
	}

	if file == "database" {
		return current
	}
	current = current[:len(current)-1]

	return projectPath(current)

}
