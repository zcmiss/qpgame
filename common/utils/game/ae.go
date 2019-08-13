/**
对接EA平台
*/
package game

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"qpgame/common/utils"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/dgrijalva/jwt-go"
	"github.com/shopspring/decimal"
)

type AE struct {
	platform string
}

func GetAe(platform string) AE {
	return AE{platform: platform}
}

//正式环境
//const EAURL  = "https://api.fafafa3388.com"
//const EAKEY  = "mgUT6Pc7UxzfVaoRkXbKuW03BYrxLaQh"

//生成token
func (ae AE) getToken(params map[string]string) string {
	secret := []byte(ramcache.GetAEConfig(ae.platform).EAKEY) //secret key
	mapclaims := make(jwt.MapClaims)
	for k, v := range params {
		mapclaims[k] = v
	}
	mapclaims["site_id"] = ramcache.GetAEConfig(ae.platform).SITEID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapclaims)
	tokenString, _ := token.SignedString(secret)
	authHeader := fmt.Sprintf("Bearer %s", tokenString)
	return authHeader
}

//发送请求
func (ae AE) postEA(apiurl string, params map[string]string) []byte {
	result := []byte("")
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("POST", apiurl, strings.NewReader(""))
	if err != nil {
		return result
	}
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	//给一个key设定为响应的value.
	req.Header.Set("Accept", "application/json")         //必须设定该参数,POST参数才能正常提交
	req.Header.Set("Authorization", ae.getToken(params)) //必须设定该参数,POST参数才能正常提交

	if params["action"] == "get_bet_histories" {
		req.Header.Set("Accept-Encoding", "gzip") //必须设定该参数,POST参数才能正常提交
		resp, err := client.Do(req)
		defer resp.Body.Close()
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return result
		}
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Println(err.Error())
			return result
		}
		return body
	} else {
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err != nil {
			return result
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []byte("{}")
		}

		return body
	}

}

//获取游戏列表
func (ae AE) getGameList(platform string) error {
	apiurl := ramcache.GetAEConfig(ae.platform).EAURL + "/dms/api"
	params := make(map[string]string)
	params["action"] = "get_game_list"
	content := ae.postEA(apiurl, params)
	var dats []xorm.PlatformGames
	var dat = make([]map[string]interface{}, 0)
	data, _, _, _ := jsonparser.Get(content, "games")
	err := json.Unmarshal(data, &dat)
	for _, v := range dat {
		name := v["locale"].(map[string]interface{})["zhCN"].(map[string]interface{})["name"].(string)
		serviceId := int(v["id"].(float64))
		dats = append(dats, xorm.PlatformGames{ServiceCode: strconv.Itoa(serviceId), Name: name, PlatId: 2, Gamecode: strconv.Itoa(serviceId), Gt: "slot"})
	}
	_, err = models.MyEngine[platform].Delete(xorm.PlatformGames{PlatId: 2})
	if err == nil {
		go models.MyEngine[platform].Insert(dats)
	}
	return err
}

//创建游戏账号
func (ae AE) createPlayer(userid string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	apiurl := ramcache.GetAEConfig(ae.platform).EAURL + "/ams/api"
	params := make(map[string]string)
	params["action"] = "create_account"
	params["account_name"] = utils.MD5By16(platform + userid + userNameConst)
	params["currency"] = "CNY"

	content := ae.postEA(apiurl, params)
	userId, _ := strconv.Atoi(userid)
	account := xorm.PlatformAccounts{PlatId: platId, UserId: userId, Username: params["account_name"], Password: params["account_name"], Created: utils.GetNowTime()}
	if bytes.Contains(content, []byte("OK")) {
		_, err := models.MyEngine[platform].Insert(&account)
		if err != nil {
			return account, false
		}
		return account, true
	}
	return account, false
}

//获取游戏登录url
func (ae AE) getGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string) string {
	apiurl := ramcache.GetAEConfig(ae.platform).EAURL + "/ams/api"
	params := make(map[string]string)
	params["action"] = "register_token"
	params["account_name"] = accounts.Username
	params["game_id"] = gamecode
	params["lang"] = "zhCN"
	content := ae.postEA(apiurl, params)
	if bytes.Contains(content, []byte("OK")) {
		res := make(map[string]string)
		json.Unmarshal(content, &res)
		return res["game_url"]
	}
	return ""
}

//玩家存取款 amount 单位元
func (ae AE) uchips(username string, exId string, amount string) bool {
	apiurl := ramcache.GetAEConfig(ae.platform).EAURL + "/ams/api"
	params := make(map[string]string)
	params["account_name"] = username
	amount2, _ := decimal.NewFromString(amount)

	params["tx_id"] = exId
	if amount2.GreaterThan(decimal.Zero) {
		params["action"] = "deposit"
	} else {
		params["action"] = "withdraw"
		amount2 = amount2.Mul(decimal.New(-1, 0))
	}
	params["amount"] = amount2.String()
	content := ae.postEA(apiurl, params)
	if bytes.Contains(content, []byte("OK")) {
		return true
	}
	return false
}

//查询玩家余额
func (ae AE) queryUchips(username string) (string, bool) {
	apiurl := ramcache.GetAEConfig(ae.platform).EAURL + "/ams/api"
	params := make(map[string]string)
	params["action"] = "get_balance"
	params["account_name"] = username
	content := ae.postEA(apiurl, params)
	if bytes.Contains(content, []byte("OK")) {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		return res["balance"].(string), true
	}
	return "0", false
}

//查询玩家余额
func (ae AE) GetBets() {
	bk, _ := ramcache.TableBetsKey.Load(ae.platform)
	betsKey := bk.(map[string]string)
	lasttime, _ := strconv.Atoi(betsKey["2-"])
	//第一次启动，自动抓取之前一个小时的数据
	now := utils.GetNowTime()
	if lasttime == 0 {
		lasttime = now - 3600*4
	}
	//往前多拉取两分钟的数据，避免第三方平台延时
	lasttime = lasttime - 120
	forCount := 1
	if now-lasttime > 900 {
		forCount = (now-lasttime)/900 + 1
		times := make(map[int][]int)
		fromtime := lasttime
		totime := lasttime + 900
		for i := 0; i < forCount-1; i++ {
			fromtime = lasttime + 900*i
			totime = fromtime + 900 - 1
			time := []int{fromtime, totime}
			times[i] = time
		}
		times[forCount-1] = []int{totime, now}
		for i := 0; i < forCount; i++ {
			go ae.getBetsByTime(times[i])
		}
	} else {
		ae.getBetsByTime([]int{lasttime + 1, now})
	}
	_, err := models.MyEngine[ae.platform].Exec("update bets_key set search_key=? where plat_id=? ", now, 2)
	if err == nil {
		betsKey["2-"] = strconv.Itoa(now)
		ramcache.TableBetsKey.Store(ae.platform, betsKey)
	}

}

func (ae AE) getBetsByTime(times []int) {
	apiurl := ramcache.GetAEConfig(ae.platform).EAURL + "/dms/api"
	params := make(map[string]string)
	params["action"] = "get_bet_histories"
	params["from_time"] = time.Unix(int64(times[0]), 0).Format(time.RFC3339)
	params["to_time"] = time.Unix(int64(times[1]), 0).Format(time.RFC3339)
	content := ae.postEA(apiurl, params)
	sqlstr := "insert ignore into bets_0 (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values"
	sqlstrs := make([]string, 0)
	for i := 0; i < 10; i++ {
		sqlstrs = append(sqlstrs, "insert ignore into bets_"+strconv.Itoa(i)+" (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values")
	}
	if bytes.Contains(content, []byte("OK")) {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		bet_histories := res["bet_histories"].([]interface{})
		if len(bet_histories) > 0 {
			for _, val := range bet_histories {
				v := val.(map[string]interface{})
				orderId := strconv.FormatFloat(v["round_id"].(float64), 'f', 0, 64)
				accountName := v["account_name"].(string)
				//根据账号获取对应的user_id
				gameCode := strconv.FormatFloat(v["game_id"].(float64), 'f', 0, 64)
				var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
				ended, _ := time.ParseInLocation(time.RFC3339, v["completed_at"].(string), cstSh)
				amount := v["bet_amt"].(string)
				reward := v["payout_amt"].(string)
				platAcc, _ := ramcache.TablePlatformAccounts.Load(ae.platform)
				userId := platAcc.(map[string]int)[accountName]
				if userId == 0 {
					ramcache.UpdateTablePlatformAccounts(accountName, ae.platform, models.MyEngine[ae.platform])
					platAcc, _ := ramcache.TablePlatformAccounts.Load(ae.platform)
					userId = platAcc.(map[string]int)[accountName]
				}
				sqlstrs[userId%10] += "('" + orderId + "','" + accountName + "','" + gameCode + "'," + strconv.Itoa(userId) + ",2," + strconv.Itoa(utils.GetNowTime()) + ",'" + amount + "','" + amount + "','" + reward + "','" + strconv.Itoa(int(ended.Unix())) + "'),"
			}
			session := models.MyEngine[ae.platform]
			sqls := ""
			for i := 0; i < 10; i++ {
				if len(sqlstrs[i]) != len(sqlstr) {
					sqls += sqlstrs[i][0:len(sqlstrs[i])-1] + ";"
				}
			}
			_, err := session.Exec(sqls)
			if err != nil {
				//fmt.Println("dddddddd" + err.Error())
			}
		}
	}
}
