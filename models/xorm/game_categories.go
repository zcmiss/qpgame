package xorm

type GameCategories struct {
	CategoryLevel  int    `json:"category_level" xorm:"not null comment('分类层级(最多三次)') TINYINT(3)"`
	Created        int    `json:"created" xorm:"not null default 0 comment('创建时间') INT(11)"`
	Id             int    `json:"id" xorm:"not null pk autoincr comment('游戏分类id') INT(11)"`
	Img            string `json:"img" xorm:"comment('2级分类游戏图片') VARCHAR(500)"`
	Name           string `json:"name" xorm:"not null comment('游戏分类名称或者游戏名称') VARCHAR(50)"`
	ParentId       int    `json:"parent_id" xorm:"not null default 0 comment('上级游戏分类id') INT(11)"`
	Seq            int    `json:"seq" xorm:"not null default 0 comment('分类排序') INT(11)"`
	Status         int    `json:"status" xorm:"not null comment('分类状态,0不可用 1可用') TINYINT(3)"`
	PlatformStatus int    `json:"platformStatus" xorm:"not null comment('游戏平台状态') TINYINT(3)"`
	Rate           string `xorm:"not null default '0.0055' comment('返水比例') VARCHAR(8)"`
	BtnSelectedImg string `json:"selectImg" xorm:"comment('2级分类游戏选中图片') VARCHAR(500)"`
	BtnImg         string `json:"selectImg" xorm:"comment('2级分类游戏按钮图片') VARCHAR(500)"`
}
