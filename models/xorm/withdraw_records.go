package xorm

type WithdrawRecords struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	UserId       int    `xorm:"not null default 0 comment('关联用户表id') index INT(10)"`
	Amount       string `xorm:"not null default 0.00 comment('提现金额') DECIMAL(12,3)"`
	OrderId      string `xorm:"not null default '' comment('提现订单号') unique VARCHAR(50)"`
	Updated      int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Status       int    `xorm:"not null default 0 comment('0 待审核,1 提现成功,2 退回出款 3 锁定') TINYINT(1)"`
	Created      int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	CardNumber   string `xorm:"not null default '' comment('银行卡号') VARCHAR(30)"`
	RealName     string `xorm:"not null default '' comment('创建时间') VARCHAR(15)"`
	BankAddress  string `xorm:"not null default '' comment('银行卡地址') VARCHAR(50)"`
	BankName     string `xorm:"not null default '' comment('银行名称') VARCHAR(30)"`
	WithdrawType string `xorm:"not null default '' comment('提现类型比如:在线提款') VARCHAR(30)"`
	Remark       string `xorm:"not null default '' comment('备注') VARCHAR(100)"`
	Operator     string `xorm:"not null default '' comment('操作者') VARCHAR(15)"`
}
