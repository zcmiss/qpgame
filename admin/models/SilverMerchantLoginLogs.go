package models

// 模型
type SilverMerchantLoginLogs struct{}

// 表名称
func (self *SilverMerchantLoginLogs) GetTableName(ctx *Context) string {
	return "silver_merchant_login_logs"
}

// 得到所有记录-分页
func (self *SilverMerchantLoginLogs) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, merchant_id, login_time, ip, login_city",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"merchant_id": "=",
				"ip":          "%",
				"login_city":  "%",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "login_time", "login_time_start", "login_time_end"))
			return append(queries)
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"login_time"}, row) //添加时间
			user := SilverMerchantUsers{}
			platform := (*ctx).Params().Get("platform")
			(*row)["merchant_name"] = user.GetMerchantName(platform, (*row)["merchant_id"])
		}, nil)
}

// 得到记录详情
func (self *SilverMerchantLoginLogs) GetRecordDetail(ctx *Context) (map[string]string, error) {
	row, err := getRecordDetail(ctx, self, "", nil)
	processDatetime(&[]string{"login_time"}, &row) //添加时间
	if _, ok := row["merchant_name"]; ok {
		user := SilverMerchantUsers{}
		platform := (*ctx).Params().Get("platform")
		row["merchant_name"] = user.GetMerchantName(platform, row["merchant_id"])
	}
	return row, err
}

// 添加记录
func (self *SilverMerchantLoginLogs) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *SilverMerchantLoginLogs) Delete(ctx *Context) error {
	return denyDelete()
}
