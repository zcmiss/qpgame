package xorm

type AdminLogs struct {
	Id      int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	AdminId string `xorm:"not null default '' VARCHAR(110)"`
	Type    string `xorm:"not null default '' comment('操作类型') VARCHAR(255)"`
	Node    string `xorm:"not null default '' comment('操作节点ID') VARCHAR(11)"`
	Content string `xorm:"not null default '' comment('操作内容') VARCHAR(255)"`
	Created int    `xorm:"default 0 comment('操作时间') INT(11)"`
}
