package model

type OracleGather struct {
	Id                               int64   `xorm:"id" json:"id"`
	BytesSentViaSQLNetToClient       float64 `xorm:"bytes_sent_via_sql_net_to_client" json:"bytesSentViaSqlNetToClient"`
	BytesReceivedViaSQLNetFromClient float64 `xorm:"bytes_received_via_sql_net_from_client" json:"bytesReceivedViaSqlNetFromClient"`
	DualTime                         int64   `xorm:"dual_time" json:"dualTime"`
	Iops                             float64 `xorm:"iops" json:"iops"`
	Mbps                             float64 `xorm:"mbps" json:"mbps"`
	UseTotalPGA                      float64 `xorm:"use_total_PGA" json:"useTotalPGA"`
	SharePoolSize                    float64 `xorm:"share_pool_size" json:"sharePoolSize"`
	SqlPinHitRatio                   float64 `xorm:"sql_pin_hit_ratio" json:"sqlPinHitRatio"`
	BufferHit                        float64 `xorm:"buffer_hit" json:"bufferHit"`
	SortMemory                       float64 `xorm:"sort_memory" json:"sortMemory"`
	ParseCountHard                   float64 `xorm:"parse_count_hard" json:"parseCountHard"`
	RedoBufferAllocationRetries      float64 `xorm:"redo_buffer_allocation_retries" json:"redoBufferAllocationRetries"`
	UserInactiveSessions             int     `xorm:"user_inactive_sessions" json:"userInactiveSessions"`
	UserActiveSessions               int     `xorm:"user_active_sessions" json:"userActiveSessions"`
	ExecuteCount                     int     `xorm:"execute_count" json:"executeCount"`
	BackgroundSessions               int     `xorm:"background_sessions" json:"backgroundSessions"`
	SqlNetRoundtripsTfromClient      float64 `xorm:"sql_net_roundtrips_tfrom_client" json:"sqlNetRoundtripsTFromClient"`
	Time                             int64   `xorm:"time" json:"time"`
	InstanceId                       int64   `xorm:"instance_id" json:"instanceId"`
}
