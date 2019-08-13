package fund

import (
	"encoding/json"
	"qpgame/common/log"
	"qpgame/models"
	"qpgame/models/xorm"
	"time"

	goxorm "github.com/go-xorm/xorm"
	"github.com/shopspring/decimal"
)

type MerchantFundChange struct {
	platform string
	typeIds  map[string]interface{}
}

// 构造函数
func NewMerchantFundChange(platform string) *MerchantFundChange {
	newMFC := new(MerchantFundChange)
	newMFC.platform = platform
	return newMFC
}

// 错误日志记录
func (cthis *MerchantFundChange) errLog(info map[string]interface{}) {
	defer log.DeferRecover()
	bs, _ := json.Marshal(info)
	log.LogPrException(string(bs))
}

//info := map[string]interface{}{
//	"merchant_id"		:	merchant_id,	// 银商编号，字符串
//  "member_user_id"	:	iMemUserId,		// 会员ID	类型,int  充值用户ID，type_id 为1 时，值为0
//  "order_id"			:	sOrderId,		// 订单ID
//	"type_id"			:	1,				// 类型,数字: 1 额度充值，2 给棋牌用户充值
//	"amount"			:	265.51,			// 操作说明，float64数字
//	"msg"				:	"银商额度充值",	// 操作说明，字符串
//  "transaction"		:	nil,			// 外部事务，*goxorm.Session类型
//}
func (cthis *MerchantFundChange) BalanceUpdate(info map[string]interface{}, dbCallback BalanceUpdateCallback) map[string]interface{} {
	//对于返回结果的定义
	actionRes := map[string]interface{}{
		"status":    0,
		"msg":       "操作失败~!!!",
		"exception": make(map[string]interface{}),
	}
	//记录本函数产生错误时所需要排查错误的参数
	localErrParams := map[string]interface{}{
		"0_platform": cthis.platform,
		"1_info":     info,
	}
	// 操作类型
	Type := info["type_id"].(int)
	if (Type != 1) && (Type != 2) {
		localErrParams["2_msg"] = "操作类型错误"
		localErrParams["3_txErr"] = nil
		actionRes["exception"] = localErrParams
		cthis.errLog(localErrParams)
		return actionRes
	}
	// 判断是否有外部已开启的事务, 如果有, 则不再新建事务
	var session *goxorm.Session
	//获取外部事务
	trans, transExists := info["transaction"]
	if transExists && (trans != nil) {
		session = trans.(*goxorm.Session)
		delete(info, "transaction")
		localErrParams["1_info"] = info
	} else {
		//没有外部事务, 则新建事务
		session = models.MyEngine[cthis.platform].NewSession()
		defer session.Close()
		txErr := session.Begin()
		if txErr != nil {
			localErrParams["2_msg"] = "开启事物失败,检查mysql连接"
			localErrParams["3_txErr"] = txErr
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
	}
	// 获取银商信息
	musers := new(xorm.SilverMerchantUsers)
	_, err := session.Where("id=?", info["merchant_id"]).Get(musers)
	if err != nil {
		localErrParams["2_msg"] = "查询银商信息失败~!"
		localErrParams["3_txErr"] = err
		actionRes["exception"] = localErrParams
		cthis.errLog(localErrParams)
		return actionRes
	}
	// 保留小数点三位
	decimal.DivisionPrecision = 3
	InfoAmount := decimal.NewFromFloat(info["amount"].(float64)) // 资金金额
	var UsableAmountOld decimal.Decimal                          // 更新前的余额

	DonateRate, _ := decimal.NewFromString(musers.DonateRate)                 // 赠送比例
	UsableAmountOld, _ = decimal.NewFromString(musers.UsableAmount)           // 可用额度
	MerchantCashPledge, _ := decimal.NewFromString(musers.MerchantCashPledge) // 银商押金
	TotalChargeMoney, _ := decimal.NewFromString(musers.TotalChargeMoney)     // 累计充值金额
	TotalAuthAmount, _ := decimal.NewFromString(musers.TotalAuthAmount)       // 累计授权金额
	UsableAmount := UsableAmountOld                                           // 可用余额
	PresentedMoney := decimal.New(0, 0)                                       // 赠送金额
	//
	if Type == 1 {
		// 获取银商配置
		configs := new(xorm.Configs)
		config := make(map[string]interface{})
		_, err = session.Where("name=?", "silver_merchant").Get(configs)
		if err == nil {
			err = json.Unmarshal([]byte(configs.Value), &config)
		}
		if _, ok := config["cash_pledge"]; !ok {
			localErrParams["2_msg"] = "查询银商配置信息失败~!"
			localErrParams["3_txErr"] = err
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
		CashPledge := decimal.NewFromFloat(config["cash_pledge"].(float64)) // 银商配置的押金金额
		if MerchantCashPledge.Equal(decimal.New(0, 0)) && InfoAmount.LessThanOrEqual(CashPledge) {
			session.Rollback()
			actionRes["status"] = 1
			actionRes["msg"] = "充值金额不足抵扣押金，操作失败!"
			return actionRes
		}
		Amount := InfoAmount                             // 充值实际到账金额(赠送金额除外)
		if MerchantCashPledge.Equal(decimal.New(0, 0)) { // 扣除押金部分
			MerchantCashPledge = CashPledge
			Amount = Amount.Sub(MerchantCashPledge)
			UsableAmountOld = decimal.New(0, 0)
			UsableAmount = decimal.New(0, 0)
			TotalAuthAmount = decimal.New(0, 0)
			TotalChargeMoney = decimal.New(0, 0)
			// 额度充值押金扣除流水
			PledgeCapitalFlow := xorm.SilverMerchantCapitalFlows{}
			PledgeCapitalFlow.Amount = "-" + CashPledge.String()
			PledgeCapitalFlow.MerchantId = musers.Id
			PledgeCapitalFlow.Type = 4
			PledgeCapitalFlow.OrderId = info["order_id"].(string)
			PledgeCapitalFlow.Balance = "0"
			PledgeCapitalFlow.ChargedAmount = "0"
			PledgeCapitalFlow.ChargedAmountOld = "0"
			PledgeCapitalFlow.Created = int(time.Now().Unix())
			PledgeCapitalFlow.MemberUserId = 0
			//创建额度充值押金扣除流水记录
			_, err = session.Insert(PledgeCapitalFlow)
			if err != nil {
				session.Rollback()
				actionRes["msg"] = "审核失败4~!"
				localErrParams["2_msg"] = "插入银商额度充值押金扣除记录失败~!"
				localErrParams["3_txErr"] = err
				actionRes["exception"] = localErrParams
				cthis.errLog(localErrParams)
				return actionRes
			}
		}
		//
		PresentedMoney = Amount.Mul(DonateRate) // 赠送金额
		// 更新silver_merchant_users表
		musers.UsableAmount = UsableAmount.Add(Amount).Add(PresentedMoney).String()
		musers.MerchantCashPledge = MerchantCashPledge.String()
		musers.TotalAuthAmount = TotalAuthAmount.Add(Amount).Add(PresentedMoney).String()
		musers.TotalChargeMoney = TotalChargeMoney.Add(Amount).String()
		_, err = session.Cols("usable_amount", "merchant_cash_pledge", "total_charge_money", "total_auth_amount").ID(musers.Id).Update(musers)
		if err != nil {
			session.Rollback()
			actionRes["msg"] = "审核失败2~!"
			localErrParams["2_msg"] = "修改银商信息失败~!"
			localErrParams["3_txErr"] = err.Error()
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
		// 额度充值流水
		ChargeCapitalFlow := xorm.SilverMerchantCapitalFlows{}
		ChargeCapitalFlow.Amount = InfoAmount.String()
		ChargeCapitalFlow.MerchantId = musers.Id
		ChargeCapitalFlow.Type = 1
		ChargeCapitalFlow.OrderId = info["order_id"].(string)
		ChargeCapitalFlow.Balance = UsableAmount.Add(Amount).String()
		ChargeCapitalFlow.ChargedAmount = UsableAmount.Add(Amount).String()
		ChargeCapitalFlow.ChargedAmountOld = UsableAmount.String()
		ChargeCapitalFlow.Created = int(time.Now().Unix())
		ChargeCapitalFlow.MemberUserId = 0
		// 额度充值赠送流水记录
		PresentedCapitalFlow := xorm.SilverMerchantCapitalFlows{}
		PresentedCapitalFlow.Amount = PresentedMoney.String()
		PresentedCapitalFlow.MerchantId = musers.Id
		PresentedCapitalFlow.Type = 3
		PresentedCapitalFlow.OrderId = info["order_id"].(string)
		PresentedCapitalFlow.Balance = UsableAmount.Add(Amount).Add(PresentedMoney).String()
		PresentedCapitalFlow.ChargedAmount = UsableAmount.Add(Amount).Add(PresentedMoney).String()
		PresentedCapitalFlow.ChargedAmountOld = UsableAmount.Add(Amount).String()
		PresentedCapitalFlow.Created = int(time.Now().Unix())
		PresentedCapitalFlow.MemberUserId = 0
		//创建充值流水记录
		_, err = session.Insert(ChargeCapitalFlow)
		if err != nil {
			session.Rollback()
			actionRes["msg"] = "审核失败4~!"
			localErrParams["2_msg"] = "插入银商充值流水记录失败~!"
			localErrParams["3_txErr"] = err
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
		// 创建充值赠送流水记录
		_, err = session.Insert(PresentedCapitalFlow)
		if err != nil {
			session.Rollback()
			actionRes["msg"] = "审核失败4~!"
			localErrParams["2_msg"] = "插入银商充值赠送流水记录失败~!"
			localErrParams["3_txErr"] = err
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
	} else {
		// 会员充值扣款
		if UsableAmount.LessThan(InfoAmount) {
			session.Rollback()
			actionRes["msg"] = "审核失败2~!"
			localErrParams["2_msg"] = "可用充值额度不足，审核失败~!"
			localErrParams["3_txErr"] = nil
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
		UsableAmount = UsableAmount.Sub(InfoAmount)
		// 更新silver_merchant_users表
		musers.UsableAmount = UsableAmount.String()
		musers.TotalAuthAmount = TotalAuthAmount.String()
		_, err = session.Cols("usable_amount", "total_auth_amount").ID(musers.Id).Update(musers)
		if err != nil {
			session.Rollback()
			actionRes["msg"] = "审核失败2~!"
			localErrParams["2_msg"] = "修改银商信息失败~!"
			localErrParams["3_txErr"] = err.Error()
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
		// 会员充值流水
		ChargeCapitalFlow := xorm.SilverMerchantCapitalFlows{}
		ChargeCapitalFlow.Amount = "-" + InfoAmount.String()
		ChargeCapitalFlow.MerchantId = musers.Id
		ChargeCapitalFlow.Type = 2
		ChargeCapitalFlow.OrderId = info["order_id"].(string)
		ChargeCapitalFlow.Balance = UsableAmount.String()
		ChargeCapitalFlow.ChargedAmount = UsableAmount.String()
		ChargeCapitalFlow.ChargedAmountOld = UsableAmountOld.String()
		ChargeCapitalFlow.Created = int(time.Now().Unix())
		ChargeCapitalFlow.MemberUserId = info["member_user_id"].(int)
		ChargeCapitalFlow.Msg = info["remark"].(string)
		//创建会员充值流水记录
		_, err = session.Insert(ChargeCapitalFlow)
		if err != nil {
			session.Rollback()
			actionRes["msg"] = "审核失败4~!"
			localErrParams["2_msg"] = "插入银商会员充值流水记录失败~!"
			localErrParams["3_txErr"] = err
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
	}
	//处理自定义回调
	if dbCallback != nil {
		_, err := dbCallback(session, PresentedMoney.String())
		if err != nil {
			session.Rollback()
			actionRes["msg"] = "审核失败4~!"
			localErrParams["2_msg"] = "回调内部处理出现错误~!"
			localErrParams["3_txErr"] = err.Error()
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
	}
	//如果已存在外部事务, 则内部不再进行提交
	if !transExists {
		err = session.Commit()
		if err != nil {
			session.Rollback()
			actionRes["msg"] = "交易失败5~!"
			localErrParams["2_msg"] = "事务提交失败~!"
			localErrParams["3_txErr"] = err.Error()
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
	}
	actionRes["status"] = 1
	actionRes["msg"] = "审核成功~"
	return actionRes
}
