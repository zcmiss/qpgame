package xorm

type AdminNodes struct {
	Id       int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Name     string `xorm:"not null comment('节点名称') VARCHAR(20)"`
	Route    string `xorm:"not null comment('路由') VARCHAR(100)"`
	Title    string `xorm:"comment('说明') VARCHAR(50)"`
	Method   string `xorm:"not null comment('方法') TEXT"`
	Status   int    `xorm:"default 0 comment('状态') TINYINT(1)"`
	Remark   string `xorm:"comment('备注') VARCHAR(255)"`
	Seq		 int    `xorm:"comment('排序') SMALLINT(6)"`
	ParentId int    `xorm:"not null comment('上级编号') SMALLINT(6)"`
	Level    int    `xorm:"not null comment('级别') TINYINT(1)"`
	Type     string `xorm:"not null comment('类型') VARCHAR(20)"`
}
