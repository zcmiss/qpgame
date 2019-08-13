package xorm

type VipLevels struct {
	Id                 int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Level              int    `xorm:"not null default 0 comment('等级例子(1-10)') INT(10)"`
	Name               string `xorm:"not null default '' comment('等级名称') VARCHAR(30)"`
	ValidBetMin        int    `xorm:"not null default 0 comment('有效投注金额区间起点单位万') INT(10)"`
	ValidBetMax        int    `xorm:"not null default 0 comment('有效投注金额区间封顶单位万') INT(10)"`
	UpgradeAmount      int    `xorm:"not null default 0 comment('晋级礼金') INT(10)"`
	WeeklyAmount       int    `xorm:"not null default 0 comment('周礼金') INT(10)"`
	MonthAmount        int    `xorm:"not null default 0 comment('月俸禄') INT(10)"`
	UpgradeAmountTotal int    `xorm:"not null default 0 comment('累计晋级礼金') INT(10)"`
	HasDepositSpeed    int    `xorm:"not null default 0 comment('存款加速通道(0不支持,1支持)') TINYINT(1)"`
	HasOwnService      int    `xorm:"not null default 0 comment('专属客服经理(0没有,1有)') TINYINT(1)"`
	WashCode           string `xorm:"not null default 0.00 comment('洗码率') DECIMAL(12,3)"`
}
