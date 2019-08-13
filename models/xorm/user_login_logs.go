package xorm

type UserLoginLogs struct {
	Id         int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	UserId     int    `xorm:"not null comment('用户编号') INT(11)"`
	LoginTime  int    `xorm:"not null default 0 comment('登录时间') INT(10)"`
	Ip         string `xorm:"not null default '' comment('ip') CHAR(50)"`
	Addr       string `xorm:"not null default '' comment('地址') VARCHAR(150)"`
	LogoutTime int    `xorm:"not null default 0 comment('退出时间') INT(10)"`
	LoginFrom  string `xorm:"not null default '' comment('登陆来源(ios、android)') VARCHAR(20)"`
}
