package models

import (
	"strconv"

	"qpgame/config"
	"qpgame/models"
)

// 模型
type PayCredentials struct{}

// 表名称
func (self *PayCredentials) GetTableName(ctx *Context) string {
	return "pay_credentials"
}

// 得到所有记录-分页
func (self *PayCredentials) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, plat_form, pay_name, merchant_number, private_key, corporate, id_umber, card_number, phone_number, status, "+
			"public_key, private_key_file, credential_key, callback_key, charge_amount_conf",
		func(ctx *Context) []string { //获取查询条件
			return getQueryFields(ctx, &map[string]string{
				"pay_name":        "%",
				"merchant_number": "%",
				"status":          "=",
				"card_number":     "%",
				"phone_number":    "%",
			})
		}, nil, nil)
}

// 得到所有记录,包含通道数据
func (self *PayCredentials) GetAll(ctx *Context) ([]map[string]interface{}, error) {
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	tablename := self.GetTableName(ctx)
	records := make([]map[string]interface{},0)
	rows, err := conn.SQL("select id,plat_form,pay_name from "+tablename+" where status=1").QueryString()
	if err != nil {
		return records, Error{What: "执行查询失败"}
	}
	for _, v := range rows{
		plat_form,_ := strconv.Atoi(v["plat_form"])
		items := make(map[int][]map[string]string,0)
		if _,ok := config.PayCredentialsPlayTypes[plat_form];ok{
			items = config.PayCredentialsPlayTypes[plat_form]
		}
		records = append(records, map[string]interface{}{
			"id": v["id"],
			"plat_form": v["plat_form"],
			"pay_name": v["pay_name"],
			"items": items,
		})
	}
	return records, nil
}

// 得到记录详情
func (self *PayCredentials) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *PayCredentials) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil, nil, nil, getSavedFunc("支付方式", "pay_name"))
}

// 删除记录
/*func (self *PayCredentials) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("支付方式"))
}
*/
func (self *PayCredentials) Delete(ctx *Context) error {
	return denyDelete()
}