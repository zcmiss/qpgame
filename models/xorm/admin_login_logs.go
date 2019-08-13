package xorm

type AdminLoginLogs struct {
	Id        int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	AdminId   int    `xorm:"not null INT(11)"`
	LoginTime int    `xorm:"not null default 'CURRENT_TIMESTAMP' INT(10)"`
	Ip        string `xorm:"not null VARCHAR(20)"`
}
