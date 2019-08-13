package models

// 模型
type AdminLogs struct{}

// 表名称
func (self *AdminLogs) GetTableName(ctx *Context) string {
	return "admin_logs"
}

// 得到所有记录-分页
func (self *AdminLogs) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, admin_id, type, node, content, created, admin_name",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"admin_id":   "=",
				"admin_name": "%",
			})
			return append(queries, getQueryFieldByTime(ctx, "created", "time_start", "time_end"))
		}, nil, nil)
}

// 得到记录详情
func (self *AdminLogs) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *AdminLogs) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *AdminLogs) Delete(ctx *Context) error {
	return denyDelete()
}
