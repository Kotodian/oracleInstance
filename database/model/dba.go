package model

import "time"

type Instance struct {
	Number          int64
	Name            string
	HostName        string
	Version         string
	StartTime       time.Time
	Status          string
	Parallel        string
	Thread          int
	Archiver        string
	LogSwitchWait   string
	Logins          string
	ShutDownPending string
	DatabaseStatus  string
	Role            string
	ActiveState     string
	Blocked         string
}

type Osstat struct {
	Id         int64
	StatName   string
	Value      float64
	Comments   string
	Cumulative string // YES or NO
}

type Database struct {
	Name         string
	PlatformName string
	CreateTime   time.Time
}

type PgaStat struct {
	Name  string
	Value float64
	Unit  string
}
type Sga struct {
	Name  string
	Value float64
}

type Sgastat struct {
	Pool     string
	SumBytes float64
}
type SysStat struct {
	Id        int64
	Statistic string
	Name      string
	Class     string
	Value     float64
}

type Dispatcher struct {
	Idle int64
	Busy int
}
type SessionEvent struct {
	Event string
	Value int
}

type Session struct {
	UserActiveSessions     int
	UserInactiveSessions   int
	UserBackGroundSessions int
}

type Process struct {
	Process           int
	BackgroundProcess int
	Redo              int
}
type FlashRecoveryArea struct {
	FileType                string
	PercentSpaceUsed        int64
	PercentSpaceReclaimable int64
	NumberOfFiles           int
}

type ArchivedLog struct {
	Name  string
	Total int
}

type Server struct {
	Dedicated int
	Shared    int
	Total     uintptr
}

type Hits struct {
	BufferCache float64
	RedoBuffer  float64
	PinHit      float64
}

type TableSpace struct {
	Id      int64
	Name    string
	File    string
	Auto    string
	Total   int
	Use     float64
	Free    float64
	Percent float64
}

type EventWait struct {
	Sid           int64
	Event         string
	TotalWaits    string
	TotalTimeOuts int64
	TimeWaited    int64
	AverageWait   float64
	WaitClass     string
}

type SqlCpu struct {
	CpuTime       float64
	Executions    int64
	UserName      string
	ParsingUserID int64
	SqlId         string
	ElapsedTime   float64
	SqlText       string
}

type Lock struct {
	Sid            int64
	Serial         string
	Username       string
	ObjectId       int64
	OracleUsername string
	OSUsername     string
	Process        string
}
