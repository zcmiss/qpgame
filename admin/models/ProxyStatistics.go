package models

import (
	"qpgame/admin/common"
	"strconv"
	"time"
)

// 模型
type ProxyStatistics struct{}

// 表名称
func (self *ProxyStatistics) GetTableName(ctx *Context) string {
	return "proxy_statistics"
}

// 得到所有记录-分页
func (self *ProxyStatistics) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, parent_id, level, ymd, charge, withdraw, deductions, "+
			"bet_count, charge_count, charge_user_count, withdraw_count, withdraw_user_count, sale_ratio, winning, "+
			"active, user_win, give_win, real_win, profit, reg_user, bet_new, "+
			"first_charge, first_charge_amount, proxy_count, downline_bet_user, member_user, bet_amount, proxy_ratio",
		func(ctx *Context) []string { //获取查询条件
			var queries []string
			return append(queries, getQueryFieldByDate(ctx, "ymd", "ymd_start", "ymd_end"))
		},
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			ymd := (*row)["ymd"]
			(*row)["ymd"] = ymd[:4] + "-" + ymd[4:6] + "-" + ymd[6:]
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, func(ctx *Context) (string, string, int) {
			return "", "ymd DESC", -1
		})
}

// 得到记录详情
func (self *ProxyStatistics) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			if createTime, createErr := strconv.ParseInt((*row)["created"], 10, 64); createErr == nil {
				(*row)["created"] = time.Unix(createTime, 0).Format("2006-01-02 15:04:05") //添加时间
			}
		})
}

// 添加记录
func (self *ProxyStatistics) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *ProxyStatistics) Delete(ctx *Context) error {
	return denyDelete()
}
