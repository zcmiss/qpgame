package models

import (
	"strconv"
	"time"
)

// 模型
type SilverMerchantBankCards struct{}

// 表名称
func (self *SilverMerchantBankCards) GetTableName(ctx *Context) string {
	return "silver_merchant_bank_cards"
}

// 得到所有记录-分页
func (self *SilverMerchantBankCards) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, merchant_id, card_number, address, bank_name, name, remark, status, created, updated",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"merchant_id": "=",
				"card_number": "%",
				"address":     "%",
				"bank_name":   "%",
				"name":        "%",
				"status":      "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return append(queries)
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"updated","created"}, row)
			user := SilverMerchantUsers{}
			platform := (*ctx).Params().Get("platform")
			(*row)["merchant_name"] = user.GetMerchantName(platform, (*row)["merchant_id"])
		}, nil)
}

// 得到记录详情
func (self *SilverMerchantBankCards) GetRecordDetail(ctx *Context) (map[string]string, error) {
	row, err := getRecordDetail(ctx, self, "", nil)
	processDatetime(&[]string{"updated","created"}, &row)
	if _, ok := row["merchant_name"]; ok {
		user := SilverMerchantUsers{}
		platform := (*ctx).Params().Get("platform")
		row["merchant_name"] = user.GetMerchantName(platform, row["merchant_id"])
	}
	return row, err
}

// 添加记录
func (self *SilverMerchantBankCards) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, nil, getSavedFunc("银商银行卡信息", "card_number"))
}

// 删除记录
func (self *SilverMerchantBankCards) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("银商银行卡"))
}
