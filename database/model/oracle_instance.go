package model

type OracleInstance struct {
	Id             int64  `xorm:"id",json:"id"`             //id
	Host           string `xorm:"host" json:"host"`         //主机
	HostName       string `xorm:"-" json:"host_name"`       //主机名
	UserName       string `xorm:"username" json:"username"` //用户名
	Password       string `xorm:"password" json:"-"`        //密码
	Dbname         string `xorm:"dbname" json:"-"`
	Number         int64  `xorm:"-" json:"number"`         //编号
	Status         string `xorm:"-" json:"status"`         //状态
	DatabaseStatus string `xorm:"-" json:"databaseStatus"` //数据库状态
	Type           string `xorm:"-" json:"type"`           //类型
	Version        string `xorm:"-" json:"version"`        //版本
	StartTime      int64  `xorm:"-" json:"startTime"`      //创建时间
}

type AddOracleInstance struct {
	Host     string `xorm:"host" json:"host"`
	UserName string `xorm:"username" json:"username"`
	Password string `xorm:"password" json:"password"`
	Dbname   string `xorm:"dbname" json:"dbname"`
}

type UpdateOracleInstance struct {
	Id       int64  `xorm:"id" json:"id"`
	Host     string `xorm:"host" json:"host"`
	UserName string `xorm:"username" json:"username"`
	Password string `xorm:"password" json:"password"`
	Dbname   string `xorm:"dbname" json:"dbname"`
}
