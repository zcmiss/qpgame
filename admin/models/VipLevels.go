package models

// 模型
type VipLevels struct{}

// 表名称
func (self *VipLevels) GetTableName(ctx *Context) string {
	return "vip_levels"
}

// 得到所有记录-分页
func (self *VipLevels) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, level, name, valid_bet_min, valid_bet_max, upgrade_amount, weekly_amount, month_amount, "+
			"has_deposit_speed, has_own_service, wash_code, upgrade_amount_total",
		func(ctx *Context) []string { //获取查询条件
			return getQueryFields(ctx, &map[string]string{
				"has_own_service":   "=",
				"has_deposit_speed": "=",
				"name":              "%",
			})
		},
		func(ctx *Context, row *map[string]string) {
			processOptions("has_deposit_speed", &yesNo, row)
			processOptions("has_own_service", &yesNo, row)
		}, nil)
}

// 得到记录详情
func (self *VipLevels) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *VipLevels) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil, nil, nil, getSavedFunc("VIP等级", "name"))
}

// 删除记录
func (self *VipLevels) Delete(ctx *Context) error {
	return denyDelete()
}
