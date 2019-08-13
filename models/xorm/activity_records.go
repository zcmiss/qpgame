package xorm

type ActivityRecords struct {
	Id         int    `xorm:"not null pk autoincr comment('主键自增ID') INT(11)"`
	UserId     int    `xorm:"not null default 0 comment('用户ID,关联Users.id') INT(11)"`
	ActivityId int    `xorm:"not null default 0 comment('活动ID,关联Activity.id') INT(11)"`
	Remark     string `xorm:"not null default '' comment('备注') VARCHAR(100)"`
	State      int    `xorm:"not null default 0 comment('是否处理(1为已处理，0为未处理)') TINYINT(1)"`
	Operator   string `xorm:"not null default '' comment('操作者') VARCHAR(30)"`
	Applied    int    `xorm:"not null default 0 comment('申请时间') INT(11)"`
	Created    int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	Updated    int    `xorm:"not null default 0 comment('更新时间') INT(11)"`
	IpAddr     string `xorm:"not null default '' comment('参与者IP地址') VARCHAR(15)"`
}
