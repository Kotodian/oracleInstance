package model

type OracleLastData struct {
	Id                               int64   `xorm:"id" json:"id"`
	BytesSentViaSQLNetToClient       float64 `xorm:"bytes_sent_via_sql_net_to_client" json:"bytesSentViaSqlNetToClient"`
	BytesReceivedViaSQLNetFromClient float64 `xorm:"bytes_received_via_sql_net_from_client" json:"bytesReceivedViaSqlNetFromClient"`
	ParseCountHard                   float64 `xorm:"parse_count_hard" json:"parseCountHard"`
	ExecuteCount                     int     `xorm:"execute_count" json:"executeCount"`
	Iops                             float64 `xorm:"iops" json:"iops"`
	Mbps                             float64 `xorm:"mbps" json:"mbps"`
	SqlNetRoundTripsFromClient       float64 `xorm:"sql_net_round_trips_from_client" json:"sqlNetRoundTripsFromClient"`
	InstanceId                       int64   `xorm:"instance_id" json:"instanceId"`
}
