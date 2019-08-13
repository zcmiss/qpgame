package models

import (
	"strconv"
)

// 模型
type WashCodeInfos struct{}

// 表名称
func (self *WashCodeInfos) GetTableName(ctx *Context) string {
	return "wash_code_infos"
}

// 得到所有记录-分页
func (self *WashCodeInfos) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, record_id, type_name, type_id, game_name, bet_amount, amount",
		func(ctx *Context) []string { //获取查询条件
			var conditions []string                              //查询条件数组
			if id, err := (*ctx).URLParamInt("id"); err == nil { //按编号查询
				conditions = append(conditions, "id = "+strconv.Itoa(id))
			}
			return conditions
		}, nil, nil)
}

// 得到记录详情
func (self *WashCodeInfos) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *WashCodeInfos) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *WashCodeInfos) Delete(ctx *Context) error {
	return denyDelete()
}
