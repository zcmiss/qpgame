package xorm

type Configs struct {
	Id      int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Name    string `xorm:"not null default '' comment('设置项名称') VARCHAR(100)"`
	Value   string `xorm:"not null comment('设置参数值') JSON"`
	Mark    string `xorm:"not null default '' comment('记号说明(该字段只作数据库字段应用说明,不作程序使用)') VARCHAR(100)"`
	Updated int    `xorm:"not null default 0 comment('更新时间') INT(11)"`
}
