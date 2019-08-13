package xorm

type AccountStatistics struct {
	Id                   int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Ymd                  int    `xorm:"not null default 0 comment('统计数据的年月日,如20180317,则统计的是20180317那天的数据') INT(10)"`
	RegisterNew          int    `xorm:"not null default 0 comment('新增注册人数') INT(10)"`
	EffectiveNew         int    `xorm:"not null default 0 comment('新增有效会员数') INT(10)"`
	FirstCharge          int    `xorm:"not null default 0 comment('首次充值(存入)人数') INT(10)"`
	FirstChargeAmount    string `xorm:"not null default 0.000 comment('首次充值(存入)金额') DECIMAL(20,3)"`
	FirstWithdraw        int    `xorm:"not null default 0 comment('首次提现(出款)人数') INT(10)"`
	FirstWithdrawAmount  string `xorm:"not null default 0.000 comment('首次提现(出款)金额') DECIMAL(20,3)"`
	ChargeUsers          int    `xorm:"not null default 0 comment('用户充值(线上充值)人数') INT(10)"`
	ChargeAmount         string `xorm:"not null default 0.000 comment('用户充值(线上充值)金额') DECIMAL(20,3)"`
	ManualCharge         int    `xorm:"not null default 0 comment('人工入款人数') INT(10)"`
	ManualChargeAmount   string `xorm:"not null default 0.000 comment('人工入款金额') DECIMAL(20,3)"`
	CompanyCharge        int    `xorm:"not null default 0 comment('公司入款人数') INT(10)"`
	CompanyChargeAmount  string `xorm:"not null default 0.000 comment('公司入款金额') DECIMAL(20,3)"`
	WithdrawUsers        int    `xorm:"not null default 0 comment('用户提现(线上申请提现)人数') INT(10)"`
	WithdrawAmount       string `xorm:"not null default 0.000 comment('用户提现(线上申请提现)金额') DECIMAL(20,3)"`
	ManualWithdraw       int    `xorm:"not null default 0 comment('人工出款人数') INT(10)"`
	ManualWithdrawAmount string `xorm:"not null default 0.000 comment('人工出款金额') DECIMAL(20,3)"`
	BetUsers             int    `xorm:"not null default 0 comment('投注人数') INT(10)"`
	BetAmount            string `xorm:"not null default 0.000 comment('总下注额(总投注额)') DECIMAL(20,3)"`
	WinningAmount        string `xorm:"default 0.000 comment('后台首页.中奖金额') DECIMAL(20,3)"`
	ActiveAmount         string `xorm:"default 0.000 comment('后台首页.活动优惠(活动奖金)') DECIMAL(20,3)"`
	BetRefundAmount      string `xorm:"default 0.000 comment('后台首页.盈亏-投注返利金额') DECIMAL(20,3)"`
	WithdrawDeduct       int    `xorm:"default 0 comment('账户出入款汇总.出款扣除人数') INT(10)"`
	WithdrawDeductAmount string `xorm:"default 0.000 comment('账户出入款汇总.出款扣除金额') DECIMAL(20,3)"`
	ChargeBenefits       int    `xorm:"default 0 comment('账户出入款汇总.给予优惠人数') INT(10)"`
	ChargeBenefitsAmount string `xorm:"default 0.000 comment('账户出入款汇总.给予优惠金额') DECIMAL(20,3)"`
	ProxyKickback        int    `xorm:"default 0 comment('账户出入款汇总.给予反水人数') INT(10)"`
	ProxyKickbackAmount  string `xorm:"default 0.000 comment('账户出入款汇总.给予反水金额') DECIMAL(20,3)"`
}
