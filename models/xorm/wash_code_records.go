package xorm

type WashCodeRecords struct {
	Id             int    `xorm:"not null pk autoincr INT(11)"`
	TotalBetamount string `xorm:"not null default 0.00 comment('洗码量') DECIMAL(12,3)"`
	Amount         string `xorm:"not null default 0.00 comment('洗码金额') DECIMAL(12,3)"`
	Washtime       int    `xorm:"not null default 0 comment('洗码时间') INT(10)"`
	UserId         int    `xorm:"not null default 0 comment('用户编号') INT(11)"`
}
