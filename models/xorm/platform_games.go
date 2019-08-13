package xorm

type PlatformGames struct {
	GameCategorieId int    `json:"game_categorie_id" xorm:"not null comment('游戏分类ID') INT(11)"`
	GameUrl         string `json:"game_url" xorm:"not null default '' comment('游戏地址') VARCHAR(500)"`
	Gamecode        string `json:"gamecode" xorm:"not null default '' comment('游戏编码') VARCHAR(30)"`
	Gt              string `json:"gt" xorm:"not null default '' comment('游戏分类') VARCHAR(10)"`
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Img             string `json:"img" xorm:"not null default '' comment('方形游戏图片资源') VARCHAR(500)"`
	Ishot           int    `json:"ishot" xorm:"not null default 0 comment('是否热门') TINYINT(2)"`
	Isnew           int    `json:"isnew" xorm:"not null default 0 comment('是否新游戏') TINYINT(2)"`
	Isrecommend     int    `json:"isrecommend" xorm:"not null default 0 comment('是否推荐游戏') TINYINT(2)"`
	Ishidden        int    `json:"ishidden" xorm:"not null default 0 comment('是否隐藏') TINYINT(2)"`
	Name            string `json:"name" xorm:"not null comment('游戏名称') VARCHAR(40)"`
	PlatId          int    `json:"plat_id" xorm:"comment('平台编号') INT(11)"`
	ServiceCode     string `json:"service_code" xorm:"not null default 0 comment('游戏id') VARCHAR(40)"`
	SmallImg        string `json:"small_img" xorm:"not null default '' comment('圆形游戏图片资源') VARCHAR(500)"`
	PlatformStatus  int    `json:"platformStatus" xorm:"not null comment('游戏平台状态') TINYINT(3)"`
}
