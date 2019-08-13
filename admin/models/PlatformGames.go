package models

import (
	"qpgame/admin/common"
)

// 模型
type PlatformGames struct{}

// 表名称
func (self *PlatformGames) GetTableName(ctx *Context) string {
	return "platform_games"
}

// 得到所有记录-分页
func (self *PlatformGames) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"game_categorie_id, game_url, gamecode, gt, id, img, ishot, isnew, name, plat_id, service_code, isrecommend, ishidden ",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"game_categorie_id": "=",
				"ishot":             "=",
				"isnew":             "=",
				"plat_id":           "=",
				"name":              "%",
			})
			hasSmallImg, smallErr := (*ctx).URLParamInt("has_img") //关于圆形图的判断
			if smallErr == nil {
				if hasSmallImg == 1 {
					queries = append(queries, "small_img <> ''")
				} else {
					queries = append(queries, "small_img = ''")
				}
			}
			return queries
		},
		func(ctx *Context, row *map[string]string) {
			platform := (*ctx).Params().Get("platform")
			(*row)["category_name"] = common.GetGameCategoryName(platform, (*row)["game_categorie_id"])
			(*row)["platform_name"] = common.GetGamePlatformName(platform, (*row)["plat_id"])
		}, nil)
}

// 得到记录详情
func (self *PlatformGames) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", func(ctx *Context, row *map[string]string) {
		platform := (*ctx).Params().Get("platform")
		(*row)["category_name"] = common.GetGameCategoryName(platform, (*row)["game_categorie_id"])
		(*row)["platform_name"] = common.GetGamePlatformName(platform, (*row)["plat_id"])
	})
}

// 添加记录
func (self *PlatformGames) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil, nil, nil, getSavedFunc("平台游戏", "name"))
}

// 删除记录
func (self *PlatformGames) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("平台游戏"))
}
