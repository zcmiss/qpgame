package fund

import (
	"encoding/json"
	"fmt"
	"qpgame/common/log"
	"qpgame/common/utils"
	"qpgame/common/utils/game"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"time"

	goxorm "github.com/go-xorm/xorm"
	"github.com/kataras/iris/core/errors"
	"github.com/shopspring/decimal"
)

//资金流水自定义回调
type BalanceUpdateCallback func(session *goxorm.Session, args ...interface{}) (interface{}, error)

type UserFundChange struct {
	platform string
	typeIds  map[string]interface{} //自己变动类型
}

//构造函数
func NewUserFundChange(platform string) *UserFundChange {
	newUFC := new(UserFundChange)
	newUFC.platform = platform
	newUFC.typeIds = map[string]interface{}{
		"charge":   []int{1}, //充值
		"withdraw": []int{2}, //提现
	}
	return newUFC
}

/*
	用户自己变动处理
	info := map[string]interface{}{
		"user_id":  userId, //用户ID
		"type_id":  config.GetFundChangeTypeId("充值"), //交易类型
		"amount":   result["amount"],  //交易金额
		"order_id": result["orderId"], //订单号
		"msg":      "用户充值", //交易信息
	}
*/
func (cthis *UserFundChange) BalanceUpdate(info map[string]interface{}, dbCallback BalanceUpdateCallback) map[string]interface{} {
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

	/*** 判断是否有外部已开启的事务, 如果有, 则不再新建事务 ***/
	var session *goxorm.Session
	trans, transExists := info["transaction"] //获取外部事务
	if transExists && trans != nil {          //如果有外部已经打开的事务
		session = trans.(*goxorm.Session) //转换interface{}为*goxorm.Session
		delete(info, "transaction")       //将保存的session删除掉
		localErrParams["1_info"] = info
	} else { //没有外部事务, 则新建事务
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

	sqlUpdate := "update withdraw_dama_records set state=2 where user_id = ? and state=0"
	accounts := new(xorm.Accounts)
	_, accountError := session.Where("user_id = ?", info["user_id"]).ForUpdate().Get(accounts)
	if accountError != nil {
		session.Rollback()
		localErrParams["2_msg"] = "查询账户信息失败~!"
		localErrParams["3_txErr"] = accountError
		actionRes["exception"] = localErrParams
		cthis.errLog(localErrParams)
		return actionRes
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
	infoTypeId := info["type_id"].(int)                                 //资金变动类型
	infoAmout := decimal.NewFromFloat(info["amount"].(float64))         //资金金额

	//如果是进入游戏进行额度转换并且余额小于1的话就转换失败
	if infoTypeId == config.FUNDPLATCHANGEOUT && !BalanceWallet.GreaterThan(decimal.New(1, 0)) {
		session.Rollback()
		actionRes["status"] = 1
		actionRes["msg"] = "余额不足请进行充值!"
		return actionRes
	}

	//如果是退出游戏进行额度转换并且余额小于0的话就不做处理了
	if infoTypeId == config.FUNDPLATCHANGEIN && infoAmout.Equal(decimal.New(0, 0)) {
		session.Rollback()

		actionRes["status"] = 1
		actionRes["msg"] = "金额已经小于0"
		actionRes["accounts"] = accounts
		return actionRes
	}
	//账户明细表数据
	accountInfo := new(xorm.AccountInfos)
	accountInfo.Type = info["type_id"].(int)
	accountInfo.UserId = info["user_id"].(int)
	accountInfo.Amount = infoAmout.String()
	accountInfo.Created = int(time.Now().Unix())
	accountInfo.Msg = info["msg"].(string)
	accountInfo.OrderId = info["order_id"].(string)
	accountInfo.ChargedAmount = ChargedAmount.String()
	accountInfo.ChargedAmountOld = ChargedAmount.String()
	//充值
	if infoTypeId == config.FUNDCHARGE {
		cnf, _ := ramcache.TableConfigs.Load(cthis.platform)
		//系统配置打码失效清空阙值
		clearDmTxLimit := cnf.(map[string]interface{})["clear_dm_tx_limit"].(float64)
		//如果充值之前金额小于1元，所有打码量失效
		if BalanceWallet.Add(BalanceSafe).LessThanOrEqual(decimal.NewFromFloat(clearDmTxLimit)) {
			_, err := session.Exec(sqlUpdate, info["user_id"])
			if err != nil {
				session.Rollback()
				actionRes["msg"] = "交易失败~!"
				localErrParams["2_msg"] = "更新账户流水失败~!"
				localErrParams["3_txErr"] = err.Error()
				actionRes["exception"] = localErrParams
				cthis.errLog(localErrParams)
				return actionRes
			}
		}
		BalanceWallet = BalanceWallet.Add(infoAmout)

		ChargedAmount = ChargedAmount.Add(infoAmout)
		accountInfo.ChargedAmount = ChargedAmount.String()
		accountInfo.ChargedAmountOld = ChargedAmount.Sub(infoAmout).String()
	}
	//提现
	if infoTypeId == config.FUNDWITHDRAW {
		BalanceWallet = BalanceWallet.Add(infoAmout)
		//负负得正
		WithdrawAmount = WithdrawAmount.Sub(infoAmout)
	}
	//洗码
	if infoTypeId == config.FUNDXIMA {
		BalanceWallet = BalanceWallet.Add(infoAmout)
		WashCodeAmount = WashCodeAmount.Add(infoAmout)
	}
	//佣金领取
	if infoTypeId == config.FUNDBROKERAGE {
		BalanceWallet = BalanceWallet.Add(infoAmout)
		ProxyAmount = ProxyAmount.Add(infoAmout)
	}
	//保险箱存取款
	if infoTypeId == config.FUNDSAFEBOX {
		BalanceWallet = BalanceWallet.Add(infoAmout)
		BalanceSafe = BalanceSafe.Sub(infoAmout)
		actionRes["accounts"] = accounts
	}

	if infoTypeId == config.FUNDREDPACKET ||
		infoTypeId == config.FUNDBINDPHONE ||
		infoTypeId == config.FUNDGENERALIZEAWAED ||
		infoTypeId == config.FUNDVIPWEEK ||
		infoTypeId == config.FUNDVIPMONTH ||
		infoTypeId == config.FUNDVIPUPLEVEL ||
		infoTypeId == config.FUNDACTIVITYAWARD ||
		infoTypeId == config.FUNDSIGNIN ||
		infoTypeId == config.FUNDPRESENTER {
		BalanceWallet = BalanceWallet.Add(infoAmout)
	}

	//进入游戏额度转换
	if infoTypeId == config.FUNDPLATCHANGEOUT {
		abs := BalanceWallet.Truncate(0)
		infoAmout = abs.Mul(decimal.New(-1, 0))
		accountInfo.Amount = infoAmout.String()
		//转入的时候转整数
		BalanceWallet = BalanceWallet.Sub(abs)
	}

	//退出游戏额度转换
	if infoTypeId == config.FUNDPLATCHANGEIN {
		//转入的时候转整数
		BalanceWallet = BalanceWallet.Add(infoAmout)
	}

	//更新account表
	accounts.BalanceLucky = BalanceLucky.String()
	accounts.ChargedAmount = ChargedAmount.String()
	accounts.ConsumedAmount = ConsumedAmount.String()
	accounts.WithdrawAmount = WithdrawAmount.String()
	accounts.TotalBetAmount = TotalBetAmount.String()
	accounts.BalanceSafe = BalanceSafe.String()
	accounts.BalanceWallet = BalanceWallet.String()
	accounts.WashCodeAmount = WashCodeAmount.String()
	accounts.ProxyAmount = ProxyAmount.String()
	_, errUpdateAffec := session.Cols("charged_amount",
		"consumed_amount",
		"withdraw_amount",
		"total_bet_amount",
		"wash_code_amount",
		"proxy_amount",
		"balance_lucky",
		"balance_safe",
		"balance_wallet").ID(accounts.Id).Update(accounts)
	if errUpdateAffec != nil {
		session.Rollback()
		actionRes["msg"] = "交易失败2~!"
		localErrParams["2_msg"] = "修改账户信息失败~!"
		localErrParams["3_txErr"] = errUpdateAffec.Error()
		actionRes["exception"] = localErrParams
		cthis.errLog(localErrParams)
		return actionRes
	}
	accountInfo.Balance = accounts.BalanceWallet
	//创建资金流水记录
	_, errInsertId := session.Insert(accountInfo)
	if errInsertId != nil {
		session.Rollback()
		actionRes["msg"] = "交易失败3~!"
		localErrParams["2_msg"] = "插入余额记录失败~!"
		localErrParams["3_txErr"] = errInsertId
		actionRes["exception"] = localErrParams
		cthis.errLog(localErrParams)
		return actionRes
	}
	//处理自定义回调
	if dbCallback != nil {
		var errBack error
		//进入游戏额度转换
		if infoTypeId == config.FUNDPLATCHANGEOUT {
			_, errBack = dbCallback(session, accountInfo.Id, infoAmout.Mul(decimal.New(-1, 0)).String())
		}
		//退出游戏额度转换
		if infoTypeId == config.FUNDPLATCHANGEIN {
			actionRes["accounts"] = accounts
			_, errBack = dbCallback(session, accountInfo.Id, infoAmout)
		}
		//红包领取,提现回调处理
		if infoTypeId == config.FUNDCHARGE || infoTypeId == config.FUNDWITHDRAW || infoTypeId == config.FUNDREDPACKET || infoTypeId == config.FUNDPRESENTER {
			_, errBack = dbCallback(session)
		}
		if errBack != nil {
			session.Rollback()
			actionRes["msg"] = "交易失败4~!"
			localErrParams["2_msg"] = "回调内部处理出现错误~!"
			localErrParams["3_txErr"] = errBack.Error()
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
	}

	//需要创建提现打码量数据记录
	if utils.InArrayInt(infoTypeId, config.NeedWithDrawDamaReords) {
		_, errWdR := cthis.WithdrawDaMaRecord(session, info)
		if errWdR != nil {
			session.Rollback()
			actionRes["msg"] = errWdR.Error()
			localErrParams["2_msg"] = "事物提交失败~!"
			localErrParams["3_txErr"] = errWdR.Error()
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
	}
	if !transExists { //如果已存在外部事务, 则内部不再进行提交
		err3 := session.Commit()
		//事物提交判断
		if err3 != nil {
			session.Rollback()
			actionRes["msg"] = "交易失败5~!"
			localErrParams["2_msg"] = "事物提交失败~!"
			localErrParams["3_txErr"] = err3.Error()
			actionRes["exception"] = localErrParams
			cthis.errLog(localErrParams)
			return actionRes
		}
	}

	//洗码
	if infoTypeId == config.FUNDXIMA {
		actionRes["accounts"] = accounts
	}
	actionRes["status"] = 1
	actionRes["msg"] = "交易成功~"
	return actionRes
}

//进入游戏之前进行余额转换
func (cthis *UserFundChange) BeforeLaunchGameFundChange(userId int, platid int, plataccounts *xorm.PlatformAccounts) map[string]interface{} {
	info := map[string]interface{}{
		"user_id":  userId,
		"type_id":  config.FUNDPLATCHANGEOUT,
		"order_id": utils.CreationOrder("CH", strconv.Itoa(userId)), //生成订单
		"amount":   0.0,
		"msg":      config.GetFundChangeInfoByTypeId(config.FUNDPLATCHANGEOUT),
	}
	//自定义业务处理
	dbCallback := func(session *goxorm.Session, args ...interface{}) (interface{}, error) {
		user := xorm.Users{Id: userId, LastPlatformId: platid}
		accountInfoId := args[0].(int)
		balanceWallet := args[1].(string)

		_, err := session.Cols("last_platform_id").Where("id = ?", userId).Update(user)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		if !game.Uchips(plataccounts.Username, strconv.Itoa(accountInfoId), balanceWallet, platid, cthis.platform) {
			return nil, errors.New("错误")
		}
		return nil, nil
	}
	res := cthis.BalanceUpdate(info, dbCallback)
	return res
}

//退出游戏之后进行余额转换
func (cthis *UserFundChange) AfterExitGameFundChange(userId int, platid int, plataccounts *xorm.PlatformAccounts, amount float64) map[string]interface{} {
	orderId := utils.CreationOrder("CH", strconv.Itoa(userId))
	var taskFlag = 0 // 任务标记，0为不需要处理，1为需要处理。如果平台已经将余额转出，在这之后出现事务回滚，用户余额将丢失，需要起一个任务处理
	info := map[string]interface{}{
		"user_id":  userId,
		"type_id":  config.FUNDPLATCHANGEIN,
		"order_id": orderId, //生成订单
		"amount":   amount,
		"msg":      config.GetFundChangeInfoByTypeId(config.FUNDPLATCHANGEIN),
	}
	//自定义业务处理
	dbCallback := func(session *goxorm.Session, args ...interface{}) (interface{}, error) {
		user := xorm.Users{Id: userId, LastPlatformId: 0}
		accountsId := args[0].(int)
		infoAmout := args[1].(decimal.Decimal)
		_, err := session.Cols("last_platform_id").Where("id = ?", userId).Update(user)
		if err != nil {
			return nil, err
		}
		if game.Uchips(plataccounts.Username, strconv.Itoa(accountsId), infoAmout.Mul(decimal.New(-1, 0)).String(), platid, cthis.platform) {
			// 平台已将余额转出，在这之后如果有导致事务回滚的情况，需要使用定时器处理一下。
			taskFlag = 1
		} else {
			return nil, errors.New("错误")
		}
		return nil, nil
	}
	res := cthis.BalanceUpdate(info, dbCallback)
	statusVal, statusOk := res["status"]
	// 处理taskFlag = 1 的情况
	if statusOk && statusVal.(int) != 1 {
		idx := strconv.Itoa(userId) + "_" + strconv.Itoa(platid) + "_" + strconv.Itoa(taskFlag)
		platform := cthis.platform
		tet, _ := ramcache.TableExceptionTasks.Load(platform)
		tetEntity := tet.(map[string]xorm.ExceptionTasks)
		if _, isExist := tetEntity[idx]; isExist == false {
			// 平台已经转出余额，但CK没有正确处理的情况，通过任务来处理
			info["plat_id"] = platid
			info["plat_account"] = plataccounts.Username
			taskContent, _ := json.Marshal(info)
			et := xorm.ExceptionTasks{
				UserId:      userId,
				PlatId:      platid,
				TaskContent: string(taskContent),
				Flag:        taskFlag,
				Created:     utils.GetNowTime(),
			}
			engine := models.MyEngine[platform]
			engine.Insert(&et)
			tetEntity[idx] = et
		}
	}
	return res
}

//进行提现打码量日志插入
func (cthis *UserFundChange) WithdrawDaMaRecord(session *goxorm.Session, info map[string]interface{}) (interface{}, error) {
	// 从缓存里读取打码量比例
	cnf, _ := ramcache.TableConfigs.Load(cthis.platform)
	// 资金打码量比例配置
	fundDamaRateMap := cnf.(map[string]interface{})["fund_dama_rate"].(map[string]interface{})
	iTypId := info["type_id"].(int)
	sTypId := strconv.Itoa(iTypId)
	fundDamaRate, fundDamaRateOk := fundDamaRateMap[sTypId].(map[string]interface{})
	if !fundDamaRateOk {
		fundDamaRate, fundDamaRateOk = fundDamaRateMap["1000"].(map[string]interface{})
		if !fundDamaRateOk {
			return nil, errors.New("从缓存里读取默认打码量比例失败")
		}
	}
	var fDamaRate float64
	fDamaRate = fundDamaRate["dama_rate"].(float64)
	var wiDamaR = new(xorm.WithdrawDamaRecords)
	wiDamaR.UserId = info["user_id"].(int)
	amount := decimal.NewFromFloat(info["amount"].(float64))
	wiDamaR.Amount = amount.String()
	wiDamaR.FundType = iTypId
	finishRate := decimal.NewFromFloat(fDamaRate)
	wiDamaR.FinishRate = finishRate.String()
	wiDamaR.Updated = int(time.Now().Unix())
	wiDamaR.Created = int(time.Now().Unix())
	wiDamaR.FinishedNeeded = amount.Mul(finishRate).String()
	wiDamaR.State = 0
	wiDamaR.FinishedProgress = "0"
	_, err := session.Insert(wiDamaR)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//错误日志记录
func (cthis *UserFundChange) errLog(info map[string]interface{}) {
	defer log.DeferRecover()
	//转成json串
	jsonStr, _ := json.Marshal(info)
	//异常日志
	log.LogPrException(string(jsonStr))
}
