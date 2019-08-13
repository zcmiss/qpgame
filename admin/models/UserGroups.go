package models

// 模型
type UserGroups struct{}

// 表名称
func (self *UserGroups) GetTableName(ctx *Context) string {
	return "user_groups"
}

// 得到所有记录-分页
func (self *UserGroups) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, group_name, remark, is_default",
		func(ctx *Context) []string { //获取查询条件
			return getQueryFields(ctx, &map[string]string{
				"group_name": "%",
				"is_default": "=",
			})
		},
		func(ctx *Context, row *map[string]string) {
			processOptions("is_default", &yesNo, row)
		}, nil)
}

// 得到记录详情
func (self *UserGroups) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *UserGroups) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil, nil, nil, getSavedFunc("用户分组", "group_name"))
}

// 删除记录
func (self *UserGroups) Delete(ctx *Context) error {
	return denyDelete()
}
