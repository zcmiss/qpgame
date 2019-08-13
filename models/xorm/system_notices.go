package xorm

type SystemNotices struct {
	Id      int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Title   string `xorm:"not null comment('标题') VARCHAR(255)"`
	Content string `xorm:"not null comment('公告内容') TEXT"`
	Status  int    `xorm:"not null default 0 comment('状态') TINYINT(1)"`
	Created int    `xorm:"not null default 0 comment('状态') INT(10)"`
}
