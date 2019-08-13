package models

import (
	"strconv"
	"time"
)

// 模型
type PlatformAccounts struct{}

// 表名称
func (self *PlatformAccounts) GetTableName(ctx *Context) string {
	return "platform_accounts"
}

// 得到所有记录-分页
func (self *PlatformAccounts) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, plat_id, user_id, username, password, created",
		func(ctx *Context) []string { //获取查询条件
			var conditions []string                              //查询条件数组
			if id, err := (*ctx).URLParamInt("id"); err == nil { //按编号查询
				conditions = append(conditions, "id = "+strconv.Itoa(id))
			}
			return conditions
		}, nil, nil)
}

// 得到记录详情
func (self *PlatformAccounts) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *PlatformAccounts) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, nil, getSavedFunc("平台账号", "user_id"))
}

// 删除记录
func (self *PlatformAccounts) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("平台账号"))
}
