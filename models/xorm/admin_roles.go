package xorm

type AdminRoles struct {
	Id       int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Name     string `xorm:"not null comment('名称') VARCHAR(20)"`
	ParentId int    `xorm:"comment('父级编号') SMALLINT(6)"`
	Status   int    `xorm:"comment('状态') TINYINT(1)"`
	Remark   string `xorm:"comment('备注') VARCHAR(255)"`
}
