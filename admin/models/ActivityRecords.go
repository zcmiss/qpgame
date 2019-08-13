package models

import (
	"qpgame/admin/common"
)

// 模型
type ActivityRecords struct{}

// 表名称
func (self *ActivityRecords) GetTableName(ctx *Context) string {
	return "activity_records"
}

// 得到所有记录-分页
func (self *ActivityRecords) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, activity_id, remark, state, operator, "+
			"applied, created, updated",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"activity_id": "=",
				"state":       "=",
			})
			return append(queries, getQueryFieldByTime(ctx, "applied", "start_time", "end_time"))
		},
		func(ctx *Context, row *map[string]string) {
			platform := (*ctx).Params().Get("platform")
			processDatetime(&[]string{"created", "applied", "updated"}, row)
			processOptions("state", &statusTypes, row)
			(*row)["activity_name"] = common.GetActivityName(platform, (*row)["activity_id"])
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, nil)
}

// 得到记录详情
func (self *ActivityRecords) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *ActivityRecords) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *ActivityRecords) Delete(ctx *Context) error {
	return denyDelete()
}
