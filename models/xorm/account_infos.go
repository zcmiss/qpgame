package xorm

type AccountInfos struct {
	Id               int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	UserId           int    `xorm:"not null comment('用户id') INT(11)"`
	Amount           string `xorm:"not null default 0.000 comment('金额 正数为收入，负数为支出') DECIMAL(12,3)"`
	Balance          string `xorm:"not null default 0.000 comment('变化后的余额') DECIMAL(12,3)"`
	ChargedAmount    string `xorm:"not null default 0.000 comment('变化后充值总金额') DECIMAL(12,3)"`
	ChargedAmountOld string `xorm:"not null default 0.000 comment('变化前充值总金额') DECIMAL(12,3)"`
	OrderId          string `xorm:"not null default '' comment('变化后的余额') VARCHAR(50)"`
	Msg              string `xorm:"not null default '' comment('变化后的余额') VARCHAR(50)"`
	Type             int    `xorm:"not null default -1 comment('1.充值,2.提现,3.洗码,4.保险箱存取款,5.赠送彩金,6.优惠入款,7.代理佣金提成,8.其他,9.活动奖励,10.提现未通过返还,11.平台资金转换,12.额度转换失败返还,13.额度转换-转出,14.红包收入,15.提现退款') INT(2)"`
	Created          int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
}
