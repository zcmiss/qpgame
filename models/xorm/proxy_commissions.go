package xorm

type ProxyCommissions struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	UserId          int    `xorm:"not null default 0 comment('用户编号') INT(11)"`
	ParentId        int    `xorm:"not null default 0 comment('上级代理用户编号') INT(11)"`
	BetAmount       string `xorm:"not null default 0.000 comment('个人总业绩') DECIMAL(12,3)"`
	TotalAmount     string `xorm:"not null default 0.000 comment('团队总业绩') DECIMAL(12,3)"`
	Created         int    `xorm:"not null default 0 comment('数据生成时间') INT(10)"`
	Contributions   int    `xorm:"not null default 0 comment('业绩贡献人数，不包括自己') INT(5)"`
	ProxyType       int    `xorm:"not null default 0 comment('代理类型2棋牌5真人') INT(2)"`
	ProxyLevel      int    `xorm:"not null default 0 comment('代理等级') INT(2)"`
	ProxyLevelRate  string `xorm:"not null default 0.0000 comment('代理等级对应返水率') DECIMAL(12,4)"`
	ProxyLevelName  string `xorm:"not null comment('代理等级名称') VARCHAR(20)"`
	Commission      string `xorm:"not null default 0.000 comment('个人总佣金') DECIMAL(12,3)"`
	TotalCommission string `xorm:"not null default 0.000 comment('团队总佣金') DECIMAL(12,3)"`
	CreatedStr      string `xorm:"not null comment('佣金生成的时间字符串') VARCHAR(10)"`
	States          int    `xorm:"not null default 0 comment('佣金领取状态0未领1已领') TINYINT(2)"`
}
