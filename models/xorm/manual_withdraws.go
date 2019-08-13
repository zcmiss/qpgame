package xorm

type ManualWithdraws struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	UserId   int    `xorm:"not null default 0 comment('关联用户表id') INT(11)"`
	Order    string `xorm:"not null default '' comment('订单编号') CHAR(21)"`
	Amount   string `xorm:"not null default 0.00 comment('取款金额') DECIMAL(12,3)"`
	Quantity string `xorm:"not null default 0.00 comment('打码量,例如:人工入款时,录入了打码量;人工出款时可以抵消人工入款的打码量') DECIMAL(10,2)"`
	Item     string `xorm:"not null default '' comment('存款项目') VARCHAR(10)"`
	Comment  string `xorm:"not null default '' comment('备注') VARCHAR(255)"`
	DealTime int    `xorm:"not null default 0 comment('交易日期') INT(11)"`
	Operator string `xorm:"not null default '' comment('操作人') VARCHAR(255)"`
	State    int    `xorm:"not null default 0 comment('审核状态，0为待审核，1为审核通过，2为审核失败') TINYINT(3)"`
}
