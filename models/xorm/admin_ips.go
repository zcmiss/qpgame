package xorm

type AdminIps struct {
	Id int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Ip string `xorm:"not null default '' comment('可访问的ip') VARCHAR(20)"`
}
