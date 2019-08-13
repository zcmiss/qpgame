package models

import "qpgame/admin/common"

// 模型
type WashCodeRecords struct{}

// 表名称
func (self *WashCodeRecords) GetTableName(ctx *Context) string {
	return "wash_code_records"
}

// 得到所有记录-分页
func (self *WashCodeRecords) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, total_betamount, amount, washtime, user_id",
		func(ctx *Context) []string { //获取查询条件
			var queries []string
			return append(queries, getQueryFieldByTime(ctx, "washtime", "wash_start", "wash_end"))
		},
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			processDatetime(&[]string{"washtime"}, row)
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, nil)
}

// 得到记录详情
func (self *WashCodeRecords) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			processDatetime(&[]string{"washtime"}, row)
		})
}

// 添加记录
func (self *WashCodeRecords) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *WashCodeRecords) Delete(ctx *Context) error {
	return denyDelete()
}
