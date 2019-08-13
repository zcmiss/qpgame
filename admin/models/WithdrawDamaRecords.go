package models

import (
	"qpgame/admin/common"
	"qpgame/config"
	"strconv"
)

// 模型
type WithdrawDamaRecords struct{}

// 表名称
func (self *WithdrawDamaRecords) GetTableName(ctx *Context) string {
	return "withdraw_dama_records"
}

// 得到所有记录-分页
func (self *WithdrawDamaRecords) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, amount, fund_type, finish_rate, updated, created, finished_progress, finished_needed, state",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"fund_type": "=",
				"state":     "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return queries
		},
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			processDatetime(&[]string{"created", "updated"}, row)
			platform := (*ctx).Params().Get("platform")
			(*row)["user_name"] = common.GetUserName(platform, (*row)["user_id"])
			typeId, _ := strconv.Atoi((*row)["fund_type"])
			(*row)["fund_type_name"] = config.GetFundChangeInfoByTypeId(typeId)
		}, nil)
}

// 得到记录详情
func (self *WithdrawDamaRecords) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			processDatetime(&[]string{"created", "updated"}, row)
		})
}

// 添加记录
func (self *WithdrawDamaRecords) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *WithdrawDamaRecords) Delete(ctx *Context) error {
	return denyDelete()
}
