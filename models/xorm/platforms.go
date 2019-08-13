package xorm

type Platforms struct {
	Id           int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Name         string `xorm:"not null default '' comment('平台名') VARCHAR(30)"`
	Code         string `xorm:"not null default '' comment('平台代号') VARCHAR(20)"`
	Status       int    `xorm:"not null default 0 comment('平台状态，0不可用，1可用') TINYINT(3)"`
	Logo         string `xorm:"not null default '' comment('平台logo') VARCHAR(255)"`
	IndexLogo    string `xorm:"not null default '' comment('平台首页logo,区别平台logo,有大小尺寸之分') VARCHAR(255)"`
	Content      string `xorm:"not null default '' comment('平台介绍') VARCHAR(255)"`
	Seq          int    `xorm:"not null default 0 comment('排序') INT(11)"`
	ApiCfgFields string `xorm:"not null default '' comment('正式平台配置里所需字段的键值和名称') VARCHAR(255)"`
	ApiCfgValue  string `xorm:"comment('正式平台配置') TEXT"`
}
