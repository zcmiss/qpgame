package models

// 模型
type ConversionRecords struct{}

// 表名称
func (self *ConversionRecords) GetTableName(ctx *Context) string {
	return "conversion_records"
}

// 得到所有记录-分页
func (self *ConversionRecords) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, platform_id, type, app_order_id, order_id, amount, status, created, tp_remain",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"platform_id":  "=",
				"type":         "=",
				"app_order_id": "%",
				"order_id":     "%",
				"status":       "=",
			})
			return append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
		}, nil, nil)
}

// 得到记录详情
func (self *ConversionRecords) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *ConversionRecords) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *ConversionRecords) Delete(ctx *Context) error {
	return denyDelete()
}
