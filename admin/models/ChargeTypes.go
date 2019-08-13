package models

// 模型
type ChargeTypes struct{}

// 表名称
func (self *ChargeTypes) GetTableName(ctx *Context) string {
	return "charge_types"
}

// 得到所有记录-分页
func (self *ChargeTypes) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, name, remark, state, charge_numbers, created, logo, updated, priority, logo_selected",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"name":  "%",
				"state": "=",
			})
			return append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
		},
		func(ctx *Context, row *map[string]string) {
			processOptions("state", &statusTypes, row)
		}, nil)
}

// 得到记录详情
func (self *ChargeTypes) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *ChargeTypes) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil, nil, nil, getSavedFunc("充值类型", "name"))
}

// 删除记录
func (self *ChargeTypes) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("支付方式"))
	//return denyDelete()
}
