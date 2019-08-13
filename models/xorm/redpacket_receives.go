package xorm

type RedpacketReceives struct {
	Id          int    `xorm:"not null pk autoincr INT(10)"`
	RedpacketId int    `xorm:"not null comment('红包id') INT(10)"`
	UserId      int    `xorm:"not null comment('抢到红包的用户') index INT(10)"`
	Money       string `xorm:"not null comment('抢到的金额') DECIMAL(10,2)"`
	Created     int    `xorm:"not null comment('抢红包时间') INT(10)"`
	RedType     int    `xorm:"not null comment('红包类型') INT(10)"`
	IsGet       int    `xorm:"not null comment('该红包是否已领取状态(0未领取,1已领取，2已失效)') TINYINT(4)"`
}
