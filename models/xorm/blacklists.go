package xorm

type Blacklists struct {
	Id     int `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	UserId int `xorm:"not null default 0 comment('用户id') INT(11)"`
}
