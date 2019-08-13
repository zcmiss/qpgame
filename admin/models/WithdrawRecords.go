package models

import (
	"qpgame/admin/common"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

// 模型
type WithdrawRecords struct{}

// 表名称
func (self *WithdrawRecords) GetTableName(ctx *Context) string {
	return "withdraw_records"
}

// 得到所有记录-分页
func (self *WithdrawRecords) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, amount, real_name, order_id, updated, status, created, card_number, "+
			"bank_address, bank_name, withdraw_type, remark, refuse_reason, operator",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"order_id":    "%",
				"real_name":   "%",
				"status":      "=",
				"operator":    "%",
				"card_number": "%",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "updated", "updated_start", "updated_end"))
			return append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
		},
		func(ctx *Context, row *map[string]string) {
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, nil)
}

// 得到记录详情
func (self *WithdrawRecords) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *WithdrawRecords) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *WithdrawRecords) Delete(ctx *Context) error {
	return denyDelete()
}

// 通过审核
func (self *WithdrawRecords) Allow(ctx *Context) error {
	platform := (*ctx).Params().Get("platform")
	//获取提交过来的充值记录id
	idStr := (*ctx).URLParam("id")
	_, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		return Error{What: "错误的提现记录编号"}
	}

	//获取此充值记录
	conn := models.MyEngine[platform]
	transaction := conn.NewSession()
	defer transaction.Close()
	tableName := self.GetTableName(ctx)
	sql := "SELECT status FROM " + tableName + " WHERE id = '" + idStr + "' LIMIT 1 FOR UPDATE"
	rows, rowsErr := conn.SQL(sql).QueryString()
	if rowsErr != nil { //如果有错误发生
		transaction.Rollback()
		return Error{What: ""}
	}
	row := rows[0]

	if row["status"] != "0" {
		transaction.Rollback()
		return Error{What: "提现记录异常"}
	}

	withdrawAmount, _ := strconv.ParseFloat(row["amount"], 64) //充值金额-转化
	userId, _ := strconv.Atoi(row["user_id"])                  //用户编号-转化
	orderId := utils.CreationOrder("WH", strconv.Itoa(userId))
	info := map[string]interface{}{
		"user_id":     userId,              //用户编号
		"type_id":     config.FUNDWITHDRAW, //类型: 提现
		"amount":      withdrawAmount,      //row["amount"],       //金额
		"order_id":    orderId,             //订单编号
		"msg":         "后台通过审核",            //操作说明
		"finish_rate": 1.0,                 //打码倍率
		"transaction": transaction,         //事务对象
	}
	balance := fund.NewUserFundChange(platform)
	result := balance.BalanceUpdate(info, nil)
	if result["status"] == 0 { //如果修改失败
		transaction.Rollback()
		return Error{What: "审核用户提现操作失败"}
	}

	sql = "UPDATE " + tableName + " SET status = 1 WHERE id = '" + idStr + "'"
	res, resErr := transaction.Exec(sql)
	if resErr != nil {
		transaction.Rollback()
		return Error{What: "修改提现订单公告失败"}
	}

	affected, affErr := res.RowsAffected()
	if affErr != nil || affected <= 0 {
		transaction.Rollback()
		return Error{What: "修改提现记录失败"}
	}

	commitErr := transaction.Commit()
	if commitErr != nil {
		transaction.Rollback()
		return Error{What: "提现审核失败"}
	}
	return nil
}

// 拒绝公司入款
func (self *WithdrawRecords) Deny(ctx *Context) error {
	platform := (*ctx).Params().Get("platform")
	//获取提交过来的充值记录id
	idStr := (*ctx).URLParam("id")
	_, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		return Error{What: "错误的提现记录编号"}
	}

	refuseReason := (*ctx).URLParam("refuse_reason")
	if refuseReason == "" {
		return Error{What: "审核说明不能为空"}
	}

	//获取此充值记录
	conn := models.MyEngine[platform]
	tableName := self.GetTableName(ctx)
	sql := "SELECT * FROM " + tableName + " WHERE id = '" + idStr + "' LIMIT 1 FOR UPDATE"
	rows, rowsErr := conn.SQL(sql).QueryString()
	if rowsErr != nil || len(rows) == 0 { //如果有错误发生
		return Error{What: "此提现申请记录不存在"}
	}
	row := rows[0]

	if row["status"] != "0" {
		return Error{What: "提现记录状态异常"}
	}
	// 操作者
	admin := common.GetAdmin(ctx)
	operator, updated := admin["name"], strconv.FormatInt(time.Now().Unix(), 10)
	// 开启事务
	transaction := conn.NewSession()
	defer transaction.Close()
	//修改提现记录的状态为拒绝
	sql = "UPDATE " + tableName + " SET status = '2', refuse_reason = '" + refuseReason + "', operator='" + operator + "', updated='" + updated + "'  WHERE id = '" + idStr + "' LIMIT 1"
	res, resErr := transaction.Exec(sql)
	if resErr != nil {
		transaction.Rollback()
		return Error{What: "订单操作失败"}
	}
	affected, affErr := res.RowsAffected()
	if affErr != nil || affected <= 0 {
		transaction.Rollback()
		return Error{What: "拒绝提现操作失败"}
	}
	userId, _ := strconv.Atoi(row["user_id"])
	amount, _ := decimal.NewFromString(row["amount"])

	//更新account表
	accounts := new(xorm.Accounts)
	_, err := transaction.Where("user_id = ?", row["user_id"]).ForUpdate().Get(accounts)
	if err != nil {
		transaction.Rollback()
		return Error{What: err.Error()}
	}
	//保留小数点三位
	decimal.DivisionPrecision = 3
	BalanceLucky, _ := decimal.NewFromString(accounts.BalanceLucky)     // 总中奖金额
	ChargedAmount, _ := decimal.NewFromString(accounts.ChargedAmount)   // 充值总金额
	ConsumedAmount, _ := decimal.NewFromString(accounts.ConsumedAmount) // 消费总金额
	WithdrawAmount, _ := decimal.NewFromString(accounts.WithdrawAmount) // 提现总金额
	TotalBetAmount, _ := decimal.NewFromString(accounts.TotalBetAmount) // 累计打码量
	BalanceSafe, _ := decimal.NewFromString(accounts.BalanceSafe)       // 保险箱余额
	BalanceWallet, _ := decimal.NewFromString(accounts.BalanceWallet)   // 钱包余额
	WashCodeAmount, _ := decimal.NewFromString(accounts.WashCodeAmount) // 洗码总金额
	ProxyAmount, _ := decimal.NewFromString(accounts.ProxyAmount)       // 代理佣金总金额
	infoAmout, _ := decimal.NewFromString(row["amount"])                //资金金额

	BalanceWallet = BalanceWallet.Add(infoAmout)

	accounts.BalanceLucky = BalanceLucky.String()
	accounts.ChargedAmount = ChargedAmount.String()
	accounts.ConsumedAmount = ConsumedAmount.String()
	accounts.WithdrawAmount = WithdrawAmount.Sub(infoAmout).String()
	accounts.TotalBetAmount = TotalBetAmount.String()
	accounts.BalanceSafe = BalanceSafe.String()
	accounts.BalanceWallet = BalanceWallet.String()
	accounts.WashCodeAmount = WashCodeAmount.String()
	accounts.ProxyAmount = ProxyAmount.String()
	_, err = transaction.Cols("charged_amount",
		"consumed_amount",
		"withdraw_amount",
		"total_bet_amount",
		"wash_code_amount",
		"proxy_amount",
		"balance_lucky",
		"balance_safe",
		"balance_wallet").ID(accounts.Id).Update(accounts)
	if err != nil {
		transaction.Rollback()
		return Error{What: err.Error()}
	}
	// 插入流水
	accountInfo := new(xorm.AccountInfos)
	accountInfo.Type = 10
	accountInfo.UserId = userId
	accountInfo.Balance = BalanceWallet.String()
	accountInfo.Amount = amount.String()
	accountInfo.Created = int(time.Now().Unix())
	accountInfo.Msg = "提现退款"
	accountInfo.OrderId = row["order_id"]
	accountInfo.ChargedAmount = BalanceWallet.String()
	accountInfo.ChargedAmountOld = BalanceWallet.Sub(amount).String()

	_, err = transaction.Insert(accountInfo)
	if err != nil {
		transaction.Rollback()
		return Error{What: err.Error()}
	}

	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return Error{What: "操作处理失败"}
	}

	return nil
}
