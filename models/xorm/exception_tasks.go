package xorm

type ExceptionTasks struct {
	Id          int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	UserId      int    `xorm:"not null comment('用户主键') INT(11)"`
	PlatId      int    `xorm:"not null comment('平台主键') INT(11)"`
	TaskContent string `xorm:"not null comment('需要执行的任务内容') JSON"`
	Flag        int    `xorm:"not null comment('1为第三方游戏平台正常而事物提交失败的任务') TINYINT(1)"`
	Created     int    `xorm:"not null comment('创建时间') INT(10)"`
}
