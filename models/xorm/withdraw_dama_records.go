package xorm

type WithdrawDamaRecords struct {
	Id               int    `xorm:"not null pk autoincr INT(11)"`
	UserId           int    `xorm:"not null comment('用户ID') INT(11)"`
	Amount           string `xorm:"not null default 0.00 comment('原始金额') DECIMAL(12,3)"`
	FundType         int    `xorm:"not null default 0 comment('1.充值,3.洗码,5.赠送彩金,6.优惠入款,9.活动奖励,14.红包收入') INT(10)"`
	FinishRate       string `xorm:"not null default 1.00 comment('打码量比例,需要原始资金的多少倍才算完成') DECIMAL(12,3)"`
	Updated          int    `xorm:"not null default 0 comment('修改时间') INT(10)"`
	Created          int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	FinishedProgress string `xorm:"not null default 0.00 comment('完成的资金量，产生一次流水都要更新这个金额,也就是实际打码量') DECIMAL(12,3)"`
	FinishedNeeded   string `xorm:"not null default 0.00 comment('打满量完成金额,打码量比例乘以原始金额') DECIMAL(12,3)"`
	State            int    `xorm:"not null default 0 comment('0未完成，1已完成，2已提现，3已失效') TINYINT(1)"`
}
