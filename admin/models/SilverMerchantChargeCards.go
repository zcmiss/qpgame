package models

import (
	"strconv"
	"time"
)

// 模型
type SilverMerchantChargeCards struct{}

// 银商支付状态映射
var merchantChargeCardsTypes = map[string]string{
	"0": "停用",
	"1": "可用",
}

// 表名称
func (self *SilverMerchantChargeCards) GetTableName(ctx *Context) string {
	return "silver_merchant_charge_cards"
}

// 得到所有记录-分页
func (self *SilverMerchantChargeCards) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id,name,owner,card_number,bank_address,remark,logo,mfrom,mto,priority,state,created",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"merchant_id":    "=",
				"type":           "=",
				"order_id":       "%",
				"member_user_id": "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return append(queries)
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"created"}, row)
			(*row)["state_text"] = merchantChargeCardsTypes[(*row)["state"]]
		}, nil)
}

// 得到记录详情
func (self *SilverMerchantChargeCards) GetRecordDetail(ctx *Context) (map[string]string, error) {
	row, err := getRecordDetail(ctx, self, "", nil)
	processDatetime(&[]string{"created"}, &row)
	row["state_text"] = merchantChargeCardsTypes[row["state"]]
	return row, err
}

// 添加记录
func (self *SilverMerchantChargeCards) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool {
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10)
			return true
		}, nil, getSavedFunc("银商支付方式信息", "name"))
}

// 删除记录
func (self *SilverMerchantChargeCards) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("银商支付方式"))
}
