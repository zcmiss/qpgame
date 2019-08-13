package xorm

type ProxyChessLevels struct {
	Id             int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Level          int    `xorm:"not null default 0 comment('等级例子(1-10)') INT(10)"`
	Name           string `xorm:"not null default '' comment('等级名称') VARCHAR(30)"`
	TeamTotalLow   int    `xorm:"not null default 0 comment('团队起始资金单位万/天') INT(10)"`
	TeamTotalLimit int    `xorm:"not null default 0 comment('团队起始资金单位万封顶单位万/天') INT(10)"`
	Commission     int    `xorm:"not null default 0 comment('万/返佣') INT(10)"`
	Created        int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
}
