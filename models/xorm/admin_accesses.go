package xorm

type AdminAccesses struct {
	Id           int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	RoleId       int    `xorm:"not null index SMALLINT(6)"`
	NodeId       int    `xorm:"not null index SMALLINT(6)"`
	AccessMethod string `xorm:"not null TEXT"`
	Level        int    `xorm:"not null TINYINT(1)"`
	Module       int    `xorm:"not null INT(3)"`
}
