/**
对接MG平台
*/
package game

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"qpgame/common/utils"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"time"
)

type MG struct {
	platform string
	mgconf   ramcache.MGConfig
}

func GetMg(platform string) MG {
	return MG{platform: platform, mgconf: ramcache.GetMGConfig(platform)}
}

var Mgtoken string

//生成token
func GetMGToken(platform string) string {
	mgconf := ramcache.GetMGConfig(platform)
	apiurl := mgconf.TOKENURL + "/connect/token"
	params := make(map[string]string)
	params["grant_type"] = "client_credentials"
	params["client_id"] = mgconf.MGNAME
	params["client_secret"] = mgconf.MGKEY

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	httpBuildQuery := ""
	for k, v := range params {
		//如果传进来的是已经拼接好的，就放入map,k的值就是拼接好的,value为空字符串
		if len(params) == 1 && v == "" {
			httpBuildQuery = k
		} else {
			httpBuildQuery += k + "=" + v + "&"
		}
	}
	if httpBuildQuery != "" {
		httpBuildQuery = strings.TrimRight(httpBuildQuery, "&")
	}
	req, err := http.NewRequest("POST", apiurl, strings.NewReader(httpBuildQuery))
	if err != nil {
		return ""
	}
	//利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。
	//给一个key设定为响应的value.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json") //必须设定该参数,POST参数才能正常提交
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	accessToken, _, _, _ := jsonparser.Get(body, "access_token")
	authHeader := fmt.Sprintf("Bearer %s", string(accessToken))
	return authHeader
}

//发送请求
func postMG(apiurl2 string, params map[string]string, method string, platform string) ([]byte, string) {
	apiurl := ramcache.GetMGConfig(platform).MGURL + apiurl2
	result := []byte("")
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	httpBuildQuery := ""
	for k, v := range params {
		//如果传进来的是已经拼接好的，就放入map,k的值就是拼接好的,value为空字符串
		if len(params) == 1 && v == "" {
			httpBuildQuery = k
		} else {
			httpBuildQuery += k + "=" + v + "&"
		}
	}
	if httpBuildQuery != "" {
		httpBuildQuery = strings.TrimRight(httpBuildQuery, "&")
	}
	req, err := http.NewRequest(method, apiurl, strings.NewReader(httpBuildQuery))
	if err != nil {
		fmt.Println(err.Error())
		return result, ""
	}

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	//给一个key设定为响应的value.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json") //必须设定该参数,POST参数才能正常提交
	req.Header.Set("Authorization", Mgtoken)     //必须设定该参数,POST参数才能正常提交
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return result, ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return result, ""
	}
	if strings.Contains(resp.Status, "401") {
		Mgtoken = GetMGToken(platform)
		return postMG(apiurl2, params, method, platform)
	}
	if strings.Contains(resp.Status, "20") {
		return body, "OK"
	} else {
		return body, ""
	}

}

//获取游戏列表
func (mg MG) getGameList(platform string) error {
	return nil
}

//创建游戏账号
func (mg MG) createPlayer(userid string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	apiurl := "/api/v1/agents/" + mg.mgconf.MGNAME + "/players"
	params := make(map[string]string)
	params["playerId"] = utils.MD5By16(platform + userid + userNameConst)
	_, res := postMG(apiurl, params, "POST", mg.platform)
	userId, _ := strconv.Atoi(userid)
	account := xorm.PlatformAccounts{PlatId: platId, UserId: userId, Username: params["playerId"], Password: params["playerId"], Created: utils.GetNowTime()}
	if res == "OK" {
		_, err := models.MyEngine[platform].Insert(&account)
		if err != nil {
			return account, false
		}
		return account, true
	}
	return account, false
}

//获取游戏登录url
func (mg MG) getGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string) string {
	apiurl := "/api/v1/agents/" + mg.mgconf.MGNAME + "/players/" + accounts.Username + "/sessions"
	params := make(map[string]string)
	params["contentCode"] = gamecode
	params["platform"] = "Mobile"
	params["langCode"] = "zh-CN"
	params["launchType"] = "HTML5"
	content, res := postMG(apiurl, params, "POST", mg.platform)
	if res == "OK" {
		res := make(map[string]string)
		json.Unmarshal(content, &res)
		return res["gameURL"]
	}
	return ""
}

//玩家存取款 amount 单位元
func (mg MG) uchips(username string, exId string, amount string) bool {
	apiurl := "/api/v1/agents/" + mg.mgconf.MGNAME + "/WalletTransactions"
	params := make(map[string]string)
	params["playerId"] = username
	amount2, _ := decimal.NewFromString(amount)
	params["externalTransactionId"] = exId
	if amount2.GreaterThan(decimal.Zero) {
		params["type"] = "Deposit"
	} else {
		amount2 = amount2.Mul(decimal.New(-1, 0))
		params["type"] = "Withdraw"
	}
	params["amount"] = amount2.String()
	_, res := postMG(apiurl, params, "POST", mg.platform)
	if res == "OK" {
		return true
	}
	return false
}

//查询玩家余额
func (mg MG) queryUchips(username string) (string, bool) {
	apiurl := "/api/v1/agents/" + mg.mgconf.MGNAME + "/players/" + username
	params := make(map[string]string)
	params["properties"] = "balance"
	content, res := postMG(apiurl, params, "GET", mg.platform)
	if res == "OK" {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		balance := res["balance"].(map[string]interface{})["total"].(float64)
		return strconv.FormatFloat(balance, 'f', 4, 32), true
	}
	return "0", false
}

func (mg MG) GetBets() {
	bk, _ := ramcache.TableBetsKey.Load(mg.platform)
	betsKey := bk.(map[string]string)
	apiurl := "/api/v1/agents/" + mg.mgconf.MGNAME + "/bets"
	params := make(map[string]string)
	if betsKey["3-"] != "" {
		params["startingAfter"] = betsKey["3-"]
	}
	params["limit"] = "2000"
	content, res := postMG(apiurl, params, "GET", mg.platform)
	fmt.Println(string(content))
	if res == "OK" {
		res := make([]interface{}, 0)
		json.Unmarshal(content, &res)
		if len(res) == 0 {
			return
		}
		lastId := ""
		sqlstr := "insert ignore into bets_0 (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values"
		sqlstrs := make([]string, 0)
		for i := 0; i < 10; i++ {
			sqlstrs = append(sqlstrs, "insert ignore into bets_"+strconv.Itoa(i)+" (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values")
		}
		for i, val := range res {
			v := val.(map[string]interface{})
			orderId := v["betUID"].(string)
			accountName := v["playerId"].(string)
			//根据账号获取对应的user_id
			gameCode := v["gameCode"].(string)
			gameEndTimeUTC, _ := time.Parse("2006-01-02T15:04:05", v["gameEndTimeUTC"].(string))
			ended := gameEndTimeUTC.Unix()
			amount := strconv.FormatFloat(v["betAmount"].(float64), 'f', 3, 64)
			reward := strconv.FormatFloat(v["payoutAmount"].(float64), 'f', 3, 64)
			platAcc, _ := ramcache.TablePlatformAccounts.Load(mg.platform)
			userId := platAcc.(map[string]int)[accountName]
			if userId == 0 {
				ramcache.UpdateTablePlatformAccounts(accountName, mg.platform, models.MyEngine[mg.platform])
				platAcc, _ := ramcache.TablePlatformAccounts.Load(mg.platform)
				userId = platAcc.(map[string]int)[accountName]
			}
			sqlstrs[userId%10] += "('" + orderId + "','" + accountName + "','" + gameCode + "'," + strconv.Itoa(userId) + ",3," + strconv.Itoa(utils.GetNowTime()) + ",'" + amount + "','" + amount + "','" + reward + "','" + strconv.Itoa(int(ended)) + "'),"
			if i == len(res)-1 {
				lastId = orderId
			}
			//bet := xorm.Bets{PlatformId: 1, OrderId: orderId, Accountname: accountName, GameCode: gameCode, Ented: ended, Gt: gt, Amount: amount, Reward: reward, Created: utils.GetNowTime(), UserId: userId}
		}
		session := models.MyEngine[mg.platform].NewSession()
		defer session.Close()
		err := session.Begin()
		sqls := ""
		for i := 0; i < 10; i++ {
			if len(sqlstrs[i]) != len(sqlstr) {
				sqls += sqlstrs[i][0:len(sqlstrs[i])-1] + ";"
			}
		}
		_, err = session.Exec(sqls)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
		_, err = session.Exec("update bets_key set search_key=? where plat_id=? ", lastId, 3)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
		err = session.Commit()
		if err == nil {
			betsKey["3-"] = lastId
			ramcache.BetsSearchType.Store(mg.platform, betsKey)
		}
	}
}
