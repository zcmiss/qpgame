package xorm

type UserFeedbacks struct {
	Id       int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	UserId   int    `xorm:"not null comment('用户ID') INT(11)"`
	Suggest  string `xorm:"not null default '' comment('建议内容') VARCHAR(400)"`
	ImageUrl string `xorm:"comment('图片地址') VARCHAR(500)"`
	Created  int    `xorm:"comment('创建时间') INT(11)"`
	Version  string `xorm:"comment('版本号') VARCHAR(100)"`
}
