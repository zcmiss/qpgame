package game

import (
	"encoding/xml"
	"fmt"
	"github.com/shopspring/decimal"
	"qpgame/common/utils"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"time"
)

var respCodes = map[string]string{
	"000000": "操作成功",
	"000001": "系统维护中",
	"000002": "操作权限被关闭",
	"000003": "IP不在白名单",
	"000004": "API密码错误",
	"000005": "系统繁忙",
	"000006": "查询时间范围越界",
	"000007": "输入参数错误",
	"000008": "请求过于频繁",
	"010001": "会员帐号不存在",
	"100001": "货币不可用,请联系技术",
	"100002": "无效帐号,帐号包含无效字符",
	"100003": "会员帐号已经被使用",
	"100004": "添加帐号失败",
	"100005": "货币错误",
	"200001": "登录操作失败",
	"200002": "帐号关闭",
	"300001": "输入存取款金额小于或者等于0",
	"300002": "存取款失败",
	"300003": "存款流水号已经存在系统中",
	"300004": "会员余额不足",
	"300005": "存取款Key校验不正确",
	"400001": "存取款流水号不存在",
	"500001": "修改限额错误",
	"500002": "佣金级别不存在",
	"600001": "限额不存在",
	"800001": "修改会员状态错误",
}

type ResponseType1 struct {
	XMLName xml.Name `xml:"response"`
	ErrCode string   `xml:"errcode"`
	ErrText string   `xml:"errtext"`
	Result  string   `xml:"result"`
}

type ResultType2 struct {
	Account      string `xml:"Account"`
	Balance      string `xml:"Balance"`
	UpdateTime   string `xml:"UpdateTime"`
	SerialNumber string `xml:"SerialNumber"`
	Amount       string `xml:"Amount"`
}

type ResponseType2 struct {
	XMLName xml.Name    `xml:"response"`
	ErrCode string      `xml:"errcode"`
	ErrText string      `xml:"errtext"`
	Result  ResultType2 `xml:"result"`
}

type UG struct {
	platform string
	ugConfig ramcache.UGConfig
}

//获取游戏列表
func (ug UG) getGameList(platform string) error {
	return nil
}

func GetUg(platform string) UG {
	return UG{platform: platform, ugConfig: ramcache.GetUGConfig(platform)}
}

func printRespErr(errCode string) {
	if res, exist := respCodes[errCode]; exist {
		fmt.Println(res)
	} else {
		fmt.Println("响应错误，错误未知")
	}
}

//创建游戏账号
func (ug UG) createPlayer(userId string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	cfg := ramcache.GetUGConfig(ug.platform)
	apiUrl := cfg.BaseUrl + "/ThirdAPI.asmx/Register"
	// 生成20位字符串作为会员账号
	uniStr := "ck" + strconv.Itoa(int(time.Now().Unix())) + utils.RandString(8, 1)
	engine := models.MyEngine[platform]
	var userBean xorm.Users
	iUserId, _ := strconv.Atoi(userId)
	has, _ := engine.Table("user").Where("id=?", iUserId).Cols("user_name").Get(&userBean)
	userName := uniStr
	if has == true {
		userName = userBean.UserName
	}
	params := map[string]string{
		"APIPassword":   cfg.Key,
		"MemberAccount": uniStr,
		"NickName":      userName,
		"Currency":      "RMB",
	}
	reqUrl := utils.BuildUrl(apiUrl, params)
	respData := utils.ReqGet(reqUrl, 10*time.Second)
	response := ResponseType1{}
	xmlErr := xml.Unmarshal(respData, &response)
	sPwd := utils.MD5(platform + userId + userPwdConst)
	account := xorm.PlatformAccounts{PlatId: platId, UserId: iUserId, Username: userName, Password: sPwd, Created: utils.GetNowTime()}
	if xmlErr != nil {
		fmt.Println(xmlErr)
	} else {
		if response.ErrCode == "000000" {
			_, err := engine.Insert(&account)
			if err != nil {
				return account, false
			}
			return account, true
		} else {
			printRespErr(response.ErrCode)
		}
	}
	return account, false
}

//获取游戏登录url
func (ug UG) getGameUrl(accounts *xorm.PlatformAccounts, gamecode, ip string) string {
	cfg := ramcache.GetUGConfig(ug.platform)
	apiUrl := cfg.BaseUrl + "/ThirdAPI.asmx/Login"
	params := map[string]string{
		"APIPassword":   cfg.Key,
		"MemberAccount": accounts.Username,
		"GameID":        "1",
		"WebType":       "Smart",
		"LoginIP":       ip,
		"Language":      "CH",
		"PageStyle":     "SP9",
	}
	reqUrl := utils.BuildUrl(apiUrl, params)
	respData := utils.ReqGet(reqUrl, 10*time.Second)
	response := ResponseType1{}
	xmlErr := xml.Unmarshal(respData, &response)
	if xmlErr != nil {
		fmt.Println(xmlErr)
	} else {
		if response.ErrCode == "000000" {
			return response.Result
		} else {
			printRespErr(response.ErrCode)
		}
	}
	return ""
}

//玩家存取款 amount 单位元
func (ug UG) uchips(username string, exId string, amount string) bool {
	cfg := ramcache.GetUGConfig(ug.platform)
	apiUrl := cfg.BaseUrl + "/ThirdAPI.asmx/Transfer"
	dAmount, _ := decimal.NewFromString(amount)
	var transferType string
	if dAmount.GreaterThan(decimal.Zero) {
		// amount > 0，是从CK取款存入UG
		transferType = "0"
	} else {
		// amount < 0，是从UG取款存入CK
		transferType = "1"
		dAmount = dAmount.Mul(decimal.New(-1, 0))
	}
	fAmount, _ := dAmount.Float64()
	amountFormatted := fmt.Sprintf("%.4f", fAmount)
	sKey := cfg.Key + username + amountFormatted
	sKeyLower := strings.ToLower(sKey)
	sKey = utils.MD5(sKeyLower)
	sKey = sKey[len(sKey)-6:] // 取后6位
	params := map[string]string{
		"APIPassword":   cfg.Key,
		"MemberAccount": username,
		"SerialNumber":  exId,
		"Amount":        dAmount.String(),
		"TransferType":  transferType,
		"Key":           sKey,
	}
	reqUrl := utils.BuildUrl(apiUrl, params)
	respData := utils.ReqGet(reqUrl, 10*time.Second)
	response := ResponseType2{}
	xmlErr := xml.Unmarshal(respData, &response)
	//fmt.Println("@ErrCode", response.ErrCode)
	if xmlErr != nil {
		fmt.Println(xmlErr)
	} else {
		if response.ErrCode == "000000" {
			return true
		} else {
			printRespErr(response.ErrCode)
		}
	}
	return false
}

//查询玩家余额
func (ug UG) queryUchips(username string) (string, bool) {
	cfg := ramcache.GetUGConfig(ug.platform)
	apiUrl := cfg.BaseUrl + "/ThirdAPI.asmx/GetBalance"
	params := map[string]string{
		"APIPassword":   cfg.Key,
		"MemberAccount": username,
	}
	reqUrl := utils.BuildUrl(apiUrl, params)
	respData := utils.ReqGet(reqUrl, 10*time.Second)
	response := ResponseType2{}
	xmlErr := xml.Unmarshal(respData, &response)
	//fmt.Println("@ErrCode", response.ErrCode)
	//fmt.Println("@Account", response.Result.Account)
	//fmt.Println("@Balance", response.Result.Balance)
	//fmt.Println("@Amount", response.Result.Amount)
	//fmt.Println("@SerialNumber", response.Result.SerialNumber)
	//fmt.Println("@UpdateTime", response.Result.UpdateTime)
	if xmlErr != nil {
		fmt.Println(xmlErr)
	} else {
		if response.ErrCode == "000000" {
			return response.Result.Balance, true
		} else {
			printRespErr(response.ErrCode)
		}
	}
	return "", false
}

func (ug UG) GetBets() {
	type Bet3 struct {
		BetID        string `xml:"BetID"`
		GameID       string `xml:"GameID"`
		SubGameID    string `xml:"SubGameID"`
		Account      string `xml:"Account"`
		BetAmount    string `xml:"BetAmount"`
		BetOdds      string `xml:"BetOdds"`
		AllWin       string `xml:"AllWin"`
		DeductAmount string `xml:"DeductAmount"`
		BackAmount   string `xml:"BackAmount"`
		Win          string `xml:"Win"`
		Turnover     string `xml:"Turnover"`
		Oddstyle     string `xml:"Oddstyle"`
		BetDate      string `xml:"BetDate"`
		Status       string `xml:"Status"`
		Result       string `xml:"Result"`
		ReportDate   string `xml:"ReportDate"`
		BetIP        string `xml:"BetIP"`
		UpdateTime   string `xml:"UpdateTime"`
		BetInfo      string `xml:"BetInfo"`
		BetResult    string `xml:"BetResult"`
		BetType      string `xml:"BetType"`
		BetPos       string `xml:"BetPos"`
		AgentID      string `xml:"AgentID"`
		SortNo       string `xml:"SortNo"`
	}
	type ResultType3 struct {
		Bet []Bet3 `xml:"bet"`
	}
	type ResponseType3 struct {
		XMLName xml.Name    `xml:"response"`
		ErrCode string      `xml:"errcode"`
		ErrText string      `xml:"errtext"`
		Result  ResultType3 `xml:"result"`
	}
	betKeyCache, _ := ramcache.TableBetsKey.Load(ug.platform)
	betsKey := betKeyCache.(map[string]string)
	params := make(map[string]string)
	cfg := ramcache.GetUGConfig(ug.platform)
	params["APIPassword"] = cfg.Key
	if betsKey["11-"] != "" {
		params["SortNo"] = betsKey["11-"]
	} else {
		params["SortNo"] = "0"
	}
	params["Rows"] = "2000"
	apiUrl := cfg.BaseUrl + "/ThirdAPI.asmx/GetBetSheetBySort"
	reqUrl := utils.BuildUrl(apiUrl, params)
	respData := utils.ReqGet(reqUrl, 10*time.Second)
	response := ResponseType3{}
	xmlErr := xml.Unmarshal(respData, &response)

	if xmlErr != nil {
		fmt.Println(xmlErr)
	} else {
		if response.ErrCode == "000000" {
			platAcc, _ := ramcache.TablePlatformAccounts.Load(ug.platform)
			lastIdx := ""
			sqlArr := make([]string, 10)
			valGroup := make([][]string, 10)
			for i := 0; i < 10; i++ {
				sqlArr[i] = fmt.Sprintf("INSERT IGNORE INTO bets_%d (order_id,user_id,platform_id,created,game_code,amount,amount_all,reward,ented,accountname,started) VALUES ", i)
			}

			betSportSql := "INSERT IGNORE INTO bets_sport (order_id,user_id,platform_id,created,game_code,amount,amount_all,reward,ented,accountname,started) VALUES "
			sportValArr := make([]string, 0)
			delBetSportSql := "DELETE FROM bets_sport WHERE order_id IN "
			delBetSportOrderIdArr := make([]string, 0)
			for _, bet := range response.Result.Bet {
				// 需要入库的字段
				sAccountName := bet.Account // accountname √
				iUserId := platAcc.(map[string]int)[sAccountName]
				// iUserId = 0时，获取不到用户记录
				if iUserId > 0 {
					iSqlArrIdx := iUserId % 10
					sOrderId := bet.BetID                                     // order_id √
					sUserId := strconv.Itoa(iUserId)                          // user_id √
					sGameCode := bet.SubGameID                                // game_code √
					sPlatformId := "11"                                       // platform_id √
					sCreated := strconv.Itoa(utils.GetNowTime())              // created √
					fTurnover, _ := strconv.ParseFloat(bet.Turnover, 64)      // 有效投注金额
					fBetAmount, _ := strconv.ParseFloat(bet.BetAmount, 64)    // 有效投注金额
					fWin, _ := strconv.ParseFloat(bet.Win, 64)                // 输赢金额，负数为输，正数为赢
					sAmount := strconv.FormatFloat(fTurnover, 'f', 3, 64)     // amount √
					sAmountAll := strconv.FormatFloat(fBetAmount, 'f', 3, 64) // amount √
					sReward := strconv.FormatFloat(fWin+fTurnover, 'f', 3, 64)
					var timeLoc, _ = time.LoadLocation("Asia/Shanghai")
					endTimeUTC, _ := time.ParseInLocation("2006-01-02 15:04:05", bet.UpdateTime, timeLoc)
					startTimeUTC, _ := time.ParseInLocation("2006-01-02 15:04:05", bet.BetDate, timeLoc)
					sEnded := strconv.FormatInt(endTimeUTC.Unix(), 10)     // ented √
					sStarted := strconv.FormatInt(startTimeUTC.Unix(), 10) // started √
					// order_id,user_id,platform_id,created,game_code,amount,reward,ented,accountname,started
					valStr := "('" + sOrderId + "', '" + sUserId + "', '" + sPlatformId + "', '" +
						sCreated + "', '" + sGameCode + "', '" + sAmount + "', '" + sAmountAll + "', '" + sReward + "', '" +
						sEnded + "', '" + sAccountName + "', '" + sStarted + "')"
					if bet.Status == "1" {
						sportValArr = append(sportValArr, valStr)
					} else if bet.Status == "2" {
						delBetSportOrderIdArr = append(delBetSportOrderIdArr, "'"+sOrderId+"'")
						valGroup[iSqlArrIdx] = append(valGroup[iSqlArrIdx], valStr)
					}
				}
				lastIdx = bet.SortNo
			}
			session := models.MyEngine[ug.platform].NewSession()
			defer session.Close()
			err := session.Begin()
			var execSqlArr []string

			if len(delBetSportOrderIdArr) > 0 {
				sDelBetSportOrderIds := "(" + strings.Join(delBetSportOrderIdArr, ",") + ")"
				delSql := delBetSportSql + sDelBetSportOrderIds
				_, err = session.Exec(delSql)
				if err != nil {
					fmt.Println(err.Error())
					session.Rollback()
					return
				}
			}

			if len(sportValArr) > 0 {
				sportVals := strings.Join(sportValArr, ",")
				_, err = session.Exec(betSportSql + sportVals)
				if err != nil {
					fmt.Println(err.Error())
					session.Rollback()
					return
				}
			}
			for i, valArr := range valGroup {
				if len(valArr) > 0 {
					execSqlArr = append(execSqlArr, sqlArr[i]+strings.Join(valArr, ",")+";")
				}
			}
			bulkInsertSql := strings.Join(execSqlArr, "")
			if bulkInsertSql != "" {
				_, err = session.Exec(bulkInsertSql)
				if err != nil {
					fmt.Println(err.Error())
					session.Rollback()
					return
				}
				_, err = session.Exec("UPDATE bets_key SET search_key=? WHERE plat_id=?", lastIdx, 11)
				if err != nil {
					fmt.Println(err.Error())
					session.Rollback()
					return
				}
				err = session.Commit()
				if err == nil {
					betsKey["11-"] = lastIdx
					ramcache.BetsSearchType.Store(ug.platform, betsKey)
				}
			} else {
				fmt.Println("*******UG平台没有获取到新的投注记录*******")
			}
		} else {
			printRespErr(response.ErrCode)
		}
	}
}
