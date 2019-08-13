package models

import "qpgame/admin/common"

// 模型
type Accounts struct{}

// 表名称
func (self *Accounts) GetTableName(ctx *Context) string {
	return "accounts"
}

// 得到所有记录-分页
func (self *Accounts) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, charged_amount, consumed_amount, withdraw_amount, total_bet_amount + today_bet_amount AS total_bet_amount, "+
			"balance_lucky + today_balance_lucky AS balance_lucky, balance_safe, balance_wallet, updated ",
		func(ctx *Context) []string { //获取查询条件
			var queries []string
			return append(queries, getQueryFieldByTime(ctx, "updated", "time_start", "time_end"))
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"updated"}, row)
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, nil)
}

// 得到记录详情
func (self *Accounts) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *Accounts) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *Accounts) Delete(ctx *Context) error {
	return denyDelete()
}
