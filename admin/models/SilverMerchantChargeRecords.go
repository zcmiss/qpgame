package models

import (
	"qpgame/admin/common"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/models"
	"regexp"
	"strconv"
	"time"

	"github.com/go-xorm/xorm"
)

// 模型
type SilverMerchantChargeRecords struct{}

// 充值状态映射
var merchantChargeRecordStates = map[string]string{
	"0": "待审核",
	"1": "成功",
	"2": "失败",
}

// 表名称
func (self *SilverMerchantChargeRecords) GetTableName(ctx *Context) string {
	return "silver_merchant_charge_records"
}

// 得到所有记录-分页
func (self *SilverMerchantChargeRecords) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, merchant_id, amount, order_id, card_number, bank_address, real_name, operator, remark, ip, state, created, bank_charge_time, updated, updated_last",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"merchant_id":  "=",
				"state":        "=",
				"order_id":     "%",
				"card_number":  "%",
				"bank_address": "%",
				"real_name":    "%",
				"operator":     "%",
				"ip":           "%",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return append(queries)
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"updated", "created", "bank_charge_time"}, row)
			user := SilverMerchantUsers{}
			platform := (*ctx).Params().Get("platform")
			(*row)["merchant_name"] = user.GetMerchantName(platform, (*row)["merchant_id"])
			(*row)["state_text"] = merchantChargeRecordStates[(*row)["state"]]
		}, nil)
}

// 得到记录详情
func (self *SilverMerchantChargeRecords) GetRecordDetail(ctx *Context) (map[string]string, error) {
	row, err := getRecordDetail(ctx, self, "", nil)
	processDatetime(&[]string{"updated", "created", "bank_charge_time"}, &row)
	if _, ok := row["merchant_name"]; ok {
		user := SilverMerchantUsers{}
		platform := (*ctx).Params().Get("platform")
		row["merchant_name"] = user.GetMerchantName(platform, row["merchant_id"])
		row["state_text"] = merchantChargeRecordStates[row["state"]]
	}
	return row, err
}

// 添加记录
func (self *SilverMerchantChargeRecords) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *SilverMerchantChargeRecords) Delete(ctx *Context) error {
	return denyDelete()
}

// 通过审核
func (self *SilverMerchantChargeRecords) Allow(ctx *Context) error {
	platform, id := (*ctx).Params().Get("platform"), (*ctx).URLParam("id")
	tableName, db := self.GetTableName(ctx), models.MyEngine[platform]
	// 获取此充值记录
	rows, err := db.SQL("SELECT merchant_id,amount,state FROM " + tableName + " WHERE id=" + id).QueryString()
	if (err != nil) || (len(rows) == 0) {
		return Error{What: "充值记录不存在"}
	}
	record := rows[0]
	if record["state"] != "0" {
		return Error{What: "充值记录已审核，请勿重复操作"}
	}
	amount, _ := strconv.ParseFloat(record["amount"], 64)
	if amount < 1 {
		return Error{What: "充值金额错误，操作失败"}
	}
	merchantId := record["merchant_id"]
	orderId := utils.TimestampToDateStr(int64(time.Now().Unix()), "060102150405") + utils.RandString(4, 4)
	info := map[string]interface{}{
		"merchant_id": merchantId, // 银商编号，字符串
		"order_id":    orderId,    // 订单号
		"type_id":     1,          // 类型,数字: 1 额度充值，2 给棋牌用户充值
		"amount":      amount,     // 操作说明，float64数字
		"msg":         "银商额度充值",   // 操作说明，字符串
	}
	// 回调
	var callback fund.BalanceUpdateCallback = func(db *xorm.Session, args ...interface{}) (interface{}, error) {
		// 操作者
		admin := common.GetAdmin(ctx)
		operator, operatorId := admin["name"], admin["id"]
		// 赠送金额
		presentedMoney := "0"
		if len(args) > 0 {
			presentedMoney = args[0].(string)
		}
		// 更新记录状态
		_, err = db.Exec("UPDATE " + tableName + " SET state=1,presented_money=" + presentedMoney + ",operator='" + operator + "' WHERE id=" + id)
		if err != nil {
			return nil, Error{What: "银商额度充值操作失败"}
		}
		// 添加操作记录
		sql := "INSERT INTO admin_logs (`admin_id`,`admin_name`,`type`,`node`,`content`,`created`)VALUES(" + operatorId + ", '" + operator + "', '银商额度充值','银商额度充值审核','银商额度充值审核通过(银商id:" + merchantId + ")', " + strconv.Itoa(int(time.Now().Unix())) + ")"
		_, err = db.Exec(sql)
		if err != nil {
			return nil, Error{What: "银商额度充值操作写入日志失败"}
		}
		return nil, nil
	}
	result := fund.NewMerchantFundChange(platform).BalanceUpdate(info, callback)
	if result["status"] == 0 { //如果修改失败
		return Error{What: "银商额度充值操作失败"}
	}
	return nil
}

// 拒绝银商充值
func (self *SilverMerchantChargeRecords) Deny(ctx *Context) error {
	platform := (*ctx).Params().Get("platform")
	//获取提交过来的充值记录id
	idStr := (*ctx).URLParam("id")
	//获取操作者
	operator := common.GetAdmin(ctx)
	match, _ := regexp.Match("^[1-9][0-9]*$", []byte(idStr))
	if !match {
		return Error{What: "错误的充值记录编号"}
	}
	//获取此充值记录
	conn := models.MyEngine[platform]
	tableName := self.GetTableName(ctx)
	sql := "SELECT state,merchant_id FROM " + tableName + " WHERE id=" + idStr
	rows, err := conn.SQL(sql).QueryString()
	if err != nil {
		return Error{What: "充值记录异常"}
	}
	row := rows[0]
	//
	transaction := conn.NewSession()
	defer transaction.Close()
	sql = "UPDATE " + tableName + " SET state=2,operator='" + operator["name"] + "' WHERE id=" + idStr
	result, err2 := transaction.Exec(sql)
	if err2 != nil {
		transaction.Rollback()
		return Error{What: "审核操作失败"}
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		transaction.Rollback()
		return Error{What: "拒绝充值操作失败"}
	}
	sql = "INSERT INTO admin_logs(admin_id,admin_name,type,node,content,created)VALUES(" + operator["id"] + ",'" + operator["name"] + "','拒绝银商充值','银商充值','拒绝银商充值(银商id:" + row["merchant_id"] + ")'," + strconv.FormatInt(time.Now().Unix(), 10) + ")"
	_, err = transaction.SQL(sql).QueryString()
	if err != nil {
		return Error{What: "日志写入失败"}
	}
	if transaction.Commit() != nil {
		transaction.Rollback()
		return Error{What: "审核操作处理失败"}
	}
	return nil
}
