package xorm

type AppVersions struct {
	Id          int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Version     string `xorm:"not null default '' comment('版本号') VARCHAR(10)"`
	Status      int    `xorm:"not null default 0 comment('版本状态') TINYINT(2)"`
	Description string `xorm:"not null comment('版本说明') TEXT"`
	Created     int    `xorm:"not null default 0 comment('添加时间') INT(10)"`
	Link        string `xorm:"not null default '' comment('下载地址') VARCHAR(150)"`
	PackageType int    `xorm:"not null default 1 comment('包类型:1全量包,2增量包') TINYINT(1)"`
	AppType     int    `xorm:"not null default 1 comment('APP类型:1安卓,2ios') TINYINT(1)"`
	UpdateType  int    `xorm:"not null default 1 comment('更新类型，1强制更新，2提示更新，3不提示更新') TINYINT(1)"`
	Updated     int    `xorm:"not null default 0 comment('修改时间') INT(10)"`
}
