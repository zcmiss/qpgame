package models

// 模型
type PhoneCodes struct{}

// 表名称
func (self *PhoneCodes) GetTableName(ctx *Context) string {
	return "phone_codes"
}

// 得到所有记录-分页
func (self *PhoneCodes) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, phone, code, created",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"phone": "%",
				"code":  "%",
			})
			return append(queries, getQueryFieldByTime(ctx, "created", "time_start", "time_end"))
		}, nil, nil)
}

// 得到记录详情
func (self *PhoneCodes) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *PhoneCodes) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *PhoneCodes) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("手机验证码"))
}
