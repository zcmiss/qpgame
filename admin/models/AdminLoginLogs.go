package models

import "qpgame/common/utils"

// 模型
type AdminLoginLogs struct{}

// 表名称
func (self *AdminLoginLogs) GetTableName(ctx *Context) string {
	return "admin_login_logs"
}

// 得到所有记录-分页
func (self *AdminLoginLogs) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, admin_id, admin_name, login_time, ip",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"admin_id":   "=",
				"ip":         "%",
				"admin_name": "%",
			})
			return append(queries, getQueryFieldByTime(ctx, "login_time", "login_time_start", "login_time_end"))
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"login_time"}, row)
			(*row)["ip_info"] = utils.GetIpInfo((*row)["ip"])
		}, nil)
}

// 得到记录详情
func (self *AdminLoginLogs) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *AdminLoginLogs) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *AdminLoginLogs) Delete(ctx *Context) error {
	return denyDelete()
}
