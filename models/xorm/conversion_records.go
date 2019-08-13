package xorm

type ConversionRecords struct {
	Id         int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	UserId     int    `xorm:"not null default 0 comment('用户ID') INT(11)"`
	PlatformId int    `xorm:"not null default 0 comment('第三方平台ID（关联tPPlatform表）') INT(11)"`
	Type       int    `xorm:"not null default 1 comment('转换类型 1向平台转入，2从平台转出') TINYINT(3)"`
	AppOrderId string `xorm:"not null default '' comment('本平台订单号，确保唯一，关联balanceLog的orderID') VARCHAR(25)"`
	OrderId    string `xorm:"not null default '' comment('第三方订单号') VARCHAR(80)"`
	Amount     string `xorm:"not null default 0.00 comment('上分金额') DECIMAL(10,2)"`
	Status     int    `xorm:"not null default 0 comment('订单状态，0处理中，1成功，2失败') TINYINT(3)"`
	Created    int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	TpRemain   string `xorm:"not null default 0.00 comment('第三方平台转账后金额') DECIMAL(10,2)"`
}
