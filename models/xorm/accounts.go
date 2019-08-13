package xorm

type Accounts struct {
	Id                int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	UserId            int    `xorm:"not null default 0 comment('用户编号') INT(11)"`
	ChargedAmount     string `xorm:"not null default 0.000 comment('充值总金额') DECIMAL(12,3)"`
	ConsumedAmount    string `xorm:"not null default 0.000 comment('消费总金额') DECIMAL(12,3)"`
	WithdrawAmount    string `xorm:"not null default 0.000 comment('提现总金额') DECIMAL(12,3)"`
	TodayBetAmount    string `xorm:"not null default 0.000 comment('当日累计打码量') DECIMAL(12,3)"`
	TotalBetAmount    string `xorm:"not null default 0.000 comment('累计打码量') DECIMAL(12,3)"`
	WashCodeAmount    string `xorm:"not null default 0.000 comment('洗码总金额') DECIMAL(12,3)"`
	ProxyAmount       string `xorm:"not null default 0.000 comment('代理佣金总金额') DECIMAL(12,3)"`
	TodayBalanceLucky string `xorm:"not null default 0.000 comment('当日总中奖金额') DECIMAL(12,3)"`
	BalanceLucky      string `xorm:"not null default 0.000 comment('总中奖金额') DECIMAL(12,3)"`
	BalanceSafe       string `xorm:"not null default 0.000 comment('保险箱余额') DECIMAL(12,3)"`
	BalanceWallet     string `xorm:"not null default 0.000 comment('钱包余额') DECIMAL(12,3)"`
	Updated           int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
}
