package xorm

type SilverMerchantCapitalFlows struct {
	Id               int    `xorm:"not null pk autoincr INT(11)"`
	MerchantId       int    `xorm:"not null comment('silver_merchant_users表Id') INT(11)"`
	Amount           string `xorm:"not null default 0.000 comment('金额 正数为收入，负数为支出') DECIMAL(12,3)"`
	Balance          string `xorm:"not null default 0.000 comment('变化后的余额') DECIMAL(12,3)"`
	Type             int    `xorm:"not null default -1 comment('1.额度充值,2.会员充值扣款') INT(2)"`
	Created          int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	OrderId          string `xorm:"not null default '' comment('订单') VARCHAR(65)"`
	Msg              string `xorm:"not null default '' comment('资金流描述说明') VARCHAR(100)"`
	ChargedAmount    string `xorm:"not null default 0.000 comment('资金变动之后的额度余额') DECIMAL(12,3)"`
	ChargedAmountOld string `xorm:"not null default 0.000 comment('更新之前的额度金额') DECIMAL(12,3)"`
	MemberUserId     int    `xorm:"not null default 0 comment('会员Id,只有在type=2的情况,1的话默认为0') INT(11)"`
}
