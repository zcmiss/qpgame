package xorm

type ChargeTypes struct {
	Id            int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Name          string `xorm:"comment('充值名称') VARCHAR(30)"`
	Remark        string `xorm:"comment('充值说明') VARCHAR(512)"`
	State         int    `xorm:"default 0 comment('状态(0禁用,1启用)') TINYINT(1)"`
	ChargeNumbers string `xorm:"not null default '' comment('充值金额选项例子:(50,100,300,800,1000,2000,3000,5000)') VARCHAR(150)"`
	Created       int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	Updated       int    `xorm:"not null default 0 comment('最后更新时间') INT(10)"`
	Logo          string `xorm:"not null default '' comment('保存在s3上的地址') VARCHAR(255)"`
	//Updated       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后更新时间') TIMESTAMP"`
	Priority int `xorm:"not null default 1 comment('类型排序') TINYINT(3)"`
}
