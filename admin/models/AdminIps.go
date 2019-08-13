package models

// 模型
type AdminIps struct{}

// 表名称
func (self *AdminIps) GetTableName(ctx *Context) string {
	return "admin_ips"
}

// 得到所有记录-分页
func (self *AdminIps) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, ip",
		func(ctx *Context) []string { //获取查询条件
			return getQueryFields(ctx, &map[string]string{
				"ip": "%",
			})
		}, nil, nil)
}

// 得到记录详情
func (self *AdminIps) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *AdminIps) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *AdminIps) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("后台用户授权IP"))
}
