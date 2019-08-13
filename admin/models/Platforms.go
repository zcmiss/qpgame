package models

import (
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
)

// 模型
type Platforms struct{}

// 表名称
func (self *Platforms) GetTableName(ctx *Context) string {
	return "platforms"
}

// 得到所有记录-分页
func (self *Platforms) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, name, code, status, logo, index_logo, content, seq",
		func(ctx *Context) []string { //获取查询条件
			return getQueryFields(ctx, &map[string]string{
				"code":   "%",
				"status": "=",
				"name":   "%",
			})
		},
		func(ctx *Context, row *map[string]string) {
			processOptionsFor("status", "status_name", &map[string]string{
				"0": "锁定",
				"1": "正常",
				"2": "维护中",
				"3": "敬请期待",
			}, row)
		}, nil)
}

// 得到记录详情
func (self *Platforms) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *Platforms) Save(ctx *Context) (int64, error) {
	result, err := saveRecord(ctx, self, nil, nil, nil, getSavedFunc("第三方平台", "name"))
	if err == nil{
		platform := (*ctx).Params().Get("platform")
		platforms := make([]xorm.Platforms, 0)
		errr := models.MyEngine[platform].Find(&platforms)
		if errr == nil {
			ramcache.ReloadGamePlatformApiConfig(platform, platforms)
		}
	}
	return result, err
}

// 删除记录
func (self *Platforms) Delete(ctx *Context) error {
	err := deleteRecord(ctx, self, nil, getDeletedFunc("第三方平台"))
	if err == nil{
		platform := (*ctx).Params().Get("platform")
		platforms := make([]xorm.Platforms, 0)
		errr := models.MyEngine[platform].Find(&platforms)
		if errr == nil {
			ramcache.ReloadGamePlatformApiConfig(platform, platforms)
		}
	}
	return err
}
