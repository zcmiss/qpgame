package xorm

type RedpacketSystems struct {
	Id            int    `xorm:"not null pk autoincr INT(10)"`
	Type          int    `xorm:"not null comment('种类(1节日红包,2.每日幸运红包)目前暂时只支持这俩种') INT(5)"`
	Money         string `xorm:"not null default 0.00 comment('总金额') DECIMAL(10,2)"`
	SentMoney     string `xorm:"default 0.00 comment('已派送金额') DECIMAL(10,2)"`
	SentCount     int    `xorm:"not null default 0 comment('红包派送数量') INT(5)"`
	Total         int    `xorm:"not null default 1 comment('红包总数量，如果是0就是没有红包个数限制,派完为止') INT(5)"`
	Created       int    `xorm:"not null comment('创建时间') INT(10)"`
	Status        int    `xorm:"not null default 1 comment('1正常，2已完结,3.关闭') SMALLINT(1)"`
	EndTime       int    `xorm:"not null comment('有效期') INT(10)"`
	StartTime     int    `xorm:"default 0 comment('开始时间') INT(11)"`
	Message       string `xorm:"not null default '' comment('红包消息') VARCHAR(255)"`
	CalculateType int    `xorm:"not null default 1 comment('红包发放方式(1.随机红包,2.固定金额红包)') INT(11)"`
}
