package models

import (
	"qpgame/admin/common"
	"strconv"
	"time"
)

// 模型
type ProxyCommissions struct{}

// 表名称
func (self *ProxyCommissions) GetTableName(ctx *Context) string {
	return "proxy_commissions"
}

// 得到所有记录-分页
func (self *ProxyCommissions) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, parent_id, bet_amount, total_amount, created, contributions, proxy_type, "+
			"proxy_level_rate, proxy_level_name, commission, total_commission, created_str, states AS state",
		func(ctx *Context) []string { //获取查询条件
			var queries []string
			return append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
		},
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			platform := (*ctx).Params().Get("platform")
			(*row)["user_name"] = common.GetUserName(platform, (*row)["user_id"])
			(*row)["parent_name"] = common.GetUserName(platform, (*row)["parent_id"])
			processOptionsFor("proxy_type", "proxy_type_name", &ProxyTypes, row)
			processOptionsFor("state", "state_name", &map[string]string{
				"0": "未领取",
				"1": "已领取",
			}, row)
			ymd := (*row)["created_str"]
			(*row)["created_str"] = ymd[:4] + "-" + ymd[4:6] + "-" + ymd[6:]
		}, nil)
}

// 得到记录详情
func (self *ProxyCommissions) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			if createTime, createErr := strconv.ParseInt((*row)["created"], 10, 64); createErr == nil {
				(*row)["created"] = time.Unix(createTime, 0).Format("2006-01-02 15:04:05") //添加时间
			}
		})
}

// 添加记录
func (self *ProxyCommissions) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, nil, getSavedFunc("代理佣金", "user_id"))
}

// 删除记录
func (self *ProxyCommissions) Delete(ctx *Context) error {
	return denyDelete()
}
