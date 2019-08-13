package xorm

type SystemStatistics struct {
	Id                int `xorm:"not null pk autoincr comment('编号') INT(11)"`
	Ymd               int `xorm:"not null default 0 comment('统计日期') INT(10)"`
	Charge            int `xorm:"not null default 0 comment('充值总额') DECIMAL(14, 4)"`
	Withdraw          int `xorm:"not null default 0 comment('提现总额') DECIMAL(14, 4)"`
	Deductions        int `xorm:"not null default 0 comment('扣除总额') DECIMAL(14, 4)"`
	BetAmount         int `xorm:"not null default 0 comment('下注总额') DECIMAL(14, 4)"`
	BetCount          int `xorm:"not null default 0 comment('下注总数') INT(10)"`
	ChargeCount       int `xorm:"not null default 0 comment('充值次数') INT(10)"`
	ChargeUserCount   int `xorm:"not null default 0 comment('充值人数') INT(10)"`
	WithdrawCount     int `xorm:"not null default 0 comment('提现次数') INT(10)"`
	WithdrawUserCount int `xorm:"not null default 0 comment('提现人数') INT(10)"`
	SaleRatio         int `xorm:"not null default 0 comment('销售返点') INT(10)"`
	Winning           int `xorm:"not null default 0 comment('中奖金额') DECIMAL(14, 4)"`
	ProxyRatio        int `xorm:"not null default 0 comment('代理返点') DECIMAL(14, 4)"`
	Active            int `xorm:"not null default 0 comment('活动奖励') DECIMAL(14, 4)"`
	UserWin           int `xorm:"not null default 0 comment('团队盈亏') DECIMAL(14, 4)"`
	GiveWin           int `xorm:"not null default 0 comment('派彩损益') DECIMAL(14, 4)"`
	RealWin           int `xorm:"not null default 0 comment('实际盈亏') DECIMAL(14, 4)"`
	Profit            int `xorm:"not null default 0 comment('利润') INT(10)"`
	RegUser           int `xorm:"not null default 0 comment('新增注册人数') INT(10)"`
	BetNew            int `xorm:"not null default 0 comment('新增投注总额') INT(10)"`
	DepositUser       int `xorm:"not null default 0 comment('新增存款会员') INT(10)"`
	FirstCharge       int `xorm:"not null default 0 comment('首充人数') INT(10)"`
	FirstChargeAmount int `xorm:"not null default 0 comment('首充金额') DECIMAL(14, 4)"`
	ProxyCount        int `xorm:"not null default 0 comment('代理总数') INT(10)"`
	DownlineBetUser   int `xorm:"not null default 0 comment('代理下级总数') INT(10)"`
	MemberUser        int `xorm:"not null default 0 comment('有效会员总数') INT(10)"`
	Created           int `xorm:"not null default 0 comment('统计时间') INT(10)"`
}
