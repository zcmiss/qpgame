package xorm

type Activities struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	Title           string `xorm:"not null default '' comment('标题') VARCHAR(50)"`
	SubTitle        string `xorm:"not null default '' comment('子标题') VARCHAR(50)"`
	Content         string `xorm:"not null comment('内容') TEXT"`
	Cover           string `xorm:"not null default '' comment('封面') VARCHAR(100)"`
	Created         int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	TimeStart       int    `xorm:"not null default 0 comment('开始时间') INT(10)"`
	TimeEnd         int    `xorm:"not null default 0 comment('结束时间') INT(10)"`
	Type            int    `xorm:"not null default 0 comment('类型') TINYINT(2)"`
	Status          int    `xorm:"not null default 0 comment('状态') TINYINT(1)"`
	Updated         int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	ActivityClassId int    `xorm:"not null default 0 comment('活动分类（关联活动分类表）') INT(10)"`
	IsHomeShow      int    `xorm:"not null default 0 comment('是否首页弹出显示(0为否,1为是)') TINYINT(1)"`
	TotalIpLimit    int    `xorm:"not null default 0 comment('限制IP领取总数') INT(10)"`
	DayIpLimit      int    `xorm:"not null default 0 comment('限制IP当日领取总数') INT(10)"`
	Money           string `xorm:"not null default 0.000 comment('活动奖励金额') DECIMAL(12,3)"`
	IsRepeat        int    `xorm:"not null default 0 comment('是否可以重复领取') TINYINT(1)"`
	Icon            string `xorm:"not null default '' comment('活动图标') VARCHAR(500)"`
}
