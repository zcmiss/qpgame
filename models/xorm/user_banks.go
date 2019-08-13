package xorm

type UserBanks struct {
	Id   int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Name string `xorm:"not null default '' VARCHAR(30)"`
	Logo string `xorm:"not null default '' comment('银行图标') VARCHAR(255)"`
}
