package models

// 模型
type UserBanks struct{}

// 表名称
func (self *UserBanks) GetTableName(ctx *Context) string {
	return "user_banks"
}

// 得到所有记录-分页
func (self *UserBanks) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, name, logo",
		func(ctx *Context) []string { //获取查询条件
			return getQueryFields(ctx, &map[string]string{
				"name": "%",
			})
		}, nil, nil)
}

// 得到记录详情
func (self *UserBanks) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *UserBanks) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil, nil, nil, getSavedFunc("银行信息", "name"))
}

// 删除记录
func (self *UserBanks) Delete(ctx *Context) error {
	return denyDelete()
}
