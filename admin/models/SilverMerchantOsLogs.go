package models

// 模型
type SilverMerchantOsLogs struct{}

// 表名称
func (self *SilverMerchantOsLogs) GetTableName(ctx *Context) string {
	return "silver_merchant_os_logs"
}

// 得到所有记录-分页
func (self *SilverMerchantOsLogs) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, merchant_id, content, ip, city, created",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"merchant_id": "=",
				"ip":          "%",
				"city":        "%",
				"status":      "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return append(queries)
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"created"}, row) //添加时间
			user := SilverMerchantUsers{}
			platform := (*ctx).Params().Get("platform")
			(*row)["merchant_name"] = user.GetMerchantName(platform, (*row)["merchant_id"])
		}, nil)
}

// 得到记录详情
func (self *SilverMerchantOsLogs) GetRecordDetail(ctx *Context) (map[string]string, error) {
	row, err := getRecordDetail(ctx, self, "", nil)
	processDatetime(&[]string{"created"}, &row) //添加时间
	if _, ok := row["merchant_name"]; ok {
		user := SilverMerchantUsers{}
		platform := (*ctx).Params().Get("platform")
		row["merchant_name"] = user.GetMerchantName(platform, row["merchant_id"])
	}
	return row, err
}

// 添加记录
func (self *SilverMerchantOsLogs) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *SilverMerchantOsLogs) Delete(ctx *Context) error {
	return denyDelete()
}
