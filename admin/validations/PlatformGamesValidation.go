package validations

import (
	"qpgame/common/mvc"
)

type PlatformGamesValidation struct{}

// 添加/修改动作数据验证
func (self PlatformGamesValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("game_categorie_id", "游戏分类ID必须为数字").
		//StringLength("game_url", "游戏地址长度在2-500之间", 2, 500).
		//StringLength("gamecode", "游戏编码长度在1-100之间", 1, 100).
		//StringLength("gt", "游戏分类长度在2-10之间", 2, 10).
		//StringLength("img", "方形游戏图片资源长度在2-500之间", 2, 500).
		IsNumeric("ishot", "是否热门必须为数字").
		IsNumeric("isnew", "是否新游戏必须为数字").
		IsNumeric("isrecommend", "是否推荐游戏必须为数字").
		StringLength("name", "游戏名称长度在2-40之间", 2, 40).
		IsNumeric("plat_id", "平台编号必须为数字").
		//StringLength("service_code", "游戏编码长度在2-40之间", 2, 40).
		//StringLength("small_img", "圆形游戏图片资源长度在2-500之间", 2, 500).
		Validate()
}
