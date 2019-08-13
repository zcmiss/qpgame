package xorm

type ActivityClasses struct {
	Id      int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Name    string `xorm:"not null comment('活动分类名称') VARCHAR(50)"`
	Status  int    `xorm:"not null comment('分类状态,0不可用 1可用') TINYINT(3)"`
	Seq		int    `xorm:"not null default 0 comment('分类排序') INT(11)"`
	Created int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	Updated int    `xorm:"not null default 0 comment('更新时间') INT(11)"`
}
