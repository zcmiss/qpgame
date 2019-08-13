package models

import (
	"qpgame/admin/common"
	"qpgame/common/utils"
)

// 模型
type UserLoginLogs struct{}

// 表名称
func (self *UserLoginLogs) GetTableName(ctx *Context) string {
	return "user_login_logs"
}

// 得到所有记录-分页
func (self *UserLoginLogs) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, login_time, ip, addr, logout_time, login_from",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"user_id": "=",
				"ip":      "%",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "login_time", "time_start", "time_end"))
			mobileType := (*ctx).URLParam("login_from")
			mobile, has := appTypes[mobileType]
			if !has {
				return queries
			}
			return append(queries, "login_from = '"+mobile+"'")
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"logout_time", "login_time"}, row)
			(*row)["ip_info"] = utils.GetIpInfo((*row)["ip"])
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, nil)
}

// 得到记录详情
func (self *UserLoginLogs) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *UserLoginLogs) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *UserLoginLogs) Delete(ctx *Context) error {
	return denyDelete()
}
