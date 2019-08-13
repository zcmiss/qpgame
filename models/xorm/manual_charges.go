package xorm

type ManualCharges struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	UserId   int    `xorm:"not null default 0 comment('管理用户表ID') INT(11)"`
	Order    string `xorm:"not null default '' comment('订单编号') CHAR(21)"`
	Amount   string `xorm:"not null default 0.00 comment('充值金额') DECIMAL(12,3)"`
	Benefits string `xorm:"not null default 0.00 comment('优惠金额') DECIMAL(10,2)"`
	Quantity string `xorm:"not null default 0.00 comment('打码量，0表示无需综合打码量流水审核') DECIMAL(10,2)"`
	Audit    int    `xorm:"not null default 0 comment('是否流水审核，0为否，1为是') TINYINT(3)"`
	Item     string `xorm:"not null default '' comment('存款项目') VARCHAR(10)"`
	Comment  string `xorm:"not null default '' comment('备注') VARCHAR(255)"`
	DealTime int    `xorm:"not null default 0 comment('交易日期') INT(11)"`
	Operator string `xorm:"not null default '' comment('操作人') VARCHAR(255)"`
	State    int    `xorm:"not null default 0 comment('审核状态，-1为无需审核，0为未审核，1为审核通过，2为审核作废') TINYINT(3)"`
}
