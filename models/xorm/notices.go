package xorm

type Notices struct {
	Id      int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	UserId  int    `xorm:"not null comment('用户关联id') INT(11)"`
	Title   string `xorm:"not null comment('标题') VARCHAR(200)"`
	Content string `xorm:"not null comment('内容') TEXT"`
	Status  int    `xorm:"not null default 0 comment('状态') TINYINT(1)"`
	Created int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
}
