/**
对接OG平台
*/
package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"net/url"
	"qpgame/common/utils"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"sync"
	"time"
)

var xTokenCache sync.Map

type tokenCache struct {
	token  string
	expire int64
}

type OG struct {
	platform string
	ogconfig ramcache.OGConfig
}

func GetOg(platform string) OG {
	return OG{platform: platform, ogconfig: ramcache.GetOGConfig(platform)}
}

//生成token
func (og OG) getToken() string {

	platform := og.platform
	tokenStruct, isToken := xTokenCache.Load(platform)
	if isToken {
		tokS := tokenStruct.(tokenCache)
		currentTime := time.Now().Unix()
		expire := tokS.expire
		//29分钟后重新获取
		if currentTime-expire < 29*60 {
			return tokS.token
		}
	}
	operator := og.ogconfig.XOPERATOR
	key := og.ogconfig.XKEY
	url := og.ogconfig.URL + "/token"
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	req.Header.Set("X-Operator", operator)
	req.Header.Set("X-Key", key) //必须设定该参数,POST参数才能正常提交
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	resultBody := make(map[string]interface{})
	errR := json.Unmarshal(body, &resultBody)
	if errR != nil {
		fmt.Println(errR.Error())
		return ""
	}
	status, existStatus := resultBody["status"]
	data, existData := resultBody["data"]
	if existStatus && existData && status == "success" {
		tk := data.(map[string]interface{})["token"].(string)
		xTokenCache.Store(platform, tokenCache{token: tk, expire: time.Now().Unix()})
		return tk
	}
	return ""
}

//发送请求
func (og OG) postOG(apiurl2 string, params map[string]string, method string, platform string) ([]byte, string) {
	apiurl := og.ogconfig.URL + apiurl2
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
	var req *http.Request
	if method == "POST" {
		jsonparam, _ := json.Marshal(params)
		reqa, err := http.NewRequest(method, apiurl, strings.NewReader(string(jsonparam)))
		req = reqa
		if err != nil {
			fmt.Println(err.Error())
			return result, ""
		}
	} else {
		apiurl = apiurl + "?" + httpBuildQuery
		reqa, err := http.NewRequest(method, apiurl, nil)
		req = reqa
		if err != nil {
			fmt.Println(err.Error())
			return result, ""
		}
	}

	req.Header.Set("X-Token", og.getToken())
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
	//fmt.Println(string(body))
	if bytes.Contains(body, []byte("success")) {
		return body, "OK"
	} else {
		return body, ""
	}
}

//发送请求
func (og OG) postOGBet(apiurl string, params map[string]string) ([]byte, string) {
	result := []byte("")
	data := make(url.Values)
	for k, v := range params {
		data[k] = []string{v}
	}
	fmt.Println(apiurl, data)
	req, err := http.PostForm(apiurl, data)
	if err != nil {
		fmt.Println(err.Error())
		return result, ""
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
		return result, ""
	}
	fmt.Println(string(body))
	if !bytes.Contains(body, []byte("Message")) {
		return body, "OK"
	} else {
		return body, ""
	}
}

//获取游戏列表
func (og OG) getGameList(platform string) error {
	return nil
}

//创建游戏账号
func (og OG) createPlayer(userid string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	apiurl := "/register"
	params := make(map[string]string)
	params["username"] = utils.MD5By16(platform + userid + userNameConst)
	params["country"] = "China"
	params["fullname"] = "MyUser"
	params["email"] = "myuser123@40test.com"
	params["language"] = "cn"
	params["birthdate"] = "1992-02-18"
	_, res := og.postOG(apiurl, params, "POST", og.platform)
	fmt.Println(string(res))
	userId, _ := strconv.Atoi(userid)
	account := xorm.PlatformAccounts{PlatId: platId, UserId: userId, Username: params["username"], Password: params["username"], Created: utils.GetNowTime()}
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
func (og OG) getGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string) string {
	apiurl := "/game-providers/1/games/oglive/key"
	params := make(map[string]string)
	params["username"] = accounts.Username
	content, res := og.postOG(apiurl, params, "GET", og.platform)
	fmt.Println(string(content))
	key := ""
	if res == "OK" {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		key = res["data"].(map[string]interface{})["key"].(string)
	}
	apiurl2 := "/game-providers/1/play"
	params2 := make(map[string]string)
	params2["key"] = key
	params2["type"] = "mobile"
	content, res = og.postOG(apiurl2, params2, "GET", og.platform)
	if res == "OK" {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		key = res["data"].(map[string]interface{})["url"].(string)
		return key
	}
	return ""
}

//玩家存取款 amount 单位元
func (og OG) uchips(username string, exId string, amount string) bool {
	apiurl := "/game-providers/1/balance"
	params := make(map[string]string)
	params["username"] = username
	params["playerId"] = username
	params["playerId"] = username
	params["transferId"] = exId
	amount2, _ := decimal.NewFromString(amount)
	if amount2.GreaterThan(decimal.Zero) {
		params["action"] = "IN"
	} else {
		amount2 = amount2.Mul(decimal.New(-1, 0))
		params["action"] = "OUT"
	}
	params["balance"] = amount2.String()
	_, res := og.postOG(apiurl, params, "POST", og.platform)
	if res == "OK" {
		return true
	}
	return false
}

//查询玩家余额
func (og OG) queryUchips(username string) (string, bool) {
	apiurl := "/game-providers/1/balance"
	params := make(map[string]string)
	params["username"] = username
	content, res := og.postOG(apiurl, params, "GET", og.platform)
	if res == "OK" {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		balance := res["data"].(map[string]interface{})["balance"].(string)
		return balance, true
	}
	return "0", false
}

//获取投注记录
func (og OG) GetBets() {
	bk, _ := ramcache.TableBetsKey.Load(og.platform)
	betsKey := bk.(map[string]string)
	lasttime, _ := strconv.Atoi(betsKey["10-"])
	//第一次启动，自动抓取之前一个小时的数据
	//每次抓取10分钟之前的数据
	now := utils.GetNowTime() - 600
	if lasttime == 0 {
		lasttime = now - 3600*4
	}
	forCount := 1
	if now-lasttime > 600 {
		forCount = (now-lasttime)/600 + 1
		times := make(map[int][]int)
		fromtime := lasttime
		totime := lasttime + 600
		for i := 0; i < forCount-1; i++ {
			fromtime = lasttime + 600*i
			totime = fromtime + 600
			time := []int{fromtime, totime}
			times[i] = time
		}
		times[forCount-1] = []int{totime, now}
		for i := 0; i < forCount; i++ {
			if times[i][0]+1 == times[i][1] {
				continue
			}
			og.getBetsByTime(times[i])
			//测试环境有10秒的访问限制
			//time.Sleep(time.Second*10)
		}
	} else {
		og.getBetsByTime([]int{lasttime + 1, now})
	}
	_, err := models.MyEngine[og.platform].Exec("update bets_key set search_key=? where plat_id=? ", now, 10)
	if err == nil {
		betsKey["10-"] = strconv.Itoa(now)
		ramcache.TableBetsKey.Store(og.platform, betsKey)
	}

}

func (og OG) getBetsByTime(times []int) {
	apiurl := og.ogconfig.BETURL
	params := make(map[string]string)
	params["Operator"] = og.ogconfig.XOPERATOR
	params["Provider"] = "og"
	params["Key"] = og.ogconfig.XKEY
	params["SDate"] = time.Unix(int64(times[0]), 0).Format("2006-01-02 15:04:05")
	params["EDate"] = time.Unix(int64(times[1]), 0).Format("2006-01-02 15:04:05")
	content, res := og.postOGBet(apiurl, params)
	sqlstr := "insert ignore into bets_0 (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values"
	sqlstrs := make([]string, 0)
	for i := 0; i < 10; i++ {
		sqlstrs = append(sqlstrs, "insert ignore into bets_"+strconv.Itoa(i)+" (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values")
	}
	if strings.Contains(res, "OK") {
		res := make([]interface{}, 0)
		json.Unmarshal(content, &res)
		if len(res) > 0 {
			for _, val := range res {
				v := val.(map[string]interface{})
				orderId := v["bettingcode"].(string)
				accountName := strings.Replace(v["membername"].(string), og.ogconfig.PREFIX, "", -1)
				//根据账号获取对应的user_id
				gameCode := v["gameid"].(string)
				var timeLoc, _ = time.LoadLocation("Asia/Shanghai")
				bettingdate, _ := time.ParseInLocation("2006-01-02 15:04:05", v["bettingdate"].(string), timeLoc)
				ended := strconv.Itoa(int(bettingdate.Unix()))
				amountD, _ := decimal.NewFromString(v["validbet"].(string))
				amount := amountD.String()
				amountAllD, _ := decimal.NewFromString(v["bettingamount"].(string))
				amountAll := amountAllD.String()
				rewardD, _ := decimal.NewFromString(v["winloseamount"].(string))
				reward := amountD.Add(rewardD).String()
				platAcc, _ := ramcache.TablePlatformAccounts.Load(og.platform)
				userId := platAcc.(map[string]int)[accountName]
				if userId == 0 {
					ramcache.UpdateTablePlatformAccounts(accountName, og.platform, models.MyEngine[og.platform])
					platAcc, _ := ramcache.TablePlatformAccounts.Load(og.platform)
					userId = platAcc.(map[string]int)[accountName]
				}
				sqlstrs[userId%10] += "('" + orderId + "','" + accountName + "','" + gameCode + "'," + strconv.Itoa(userId) + ",10," + strconv.Itoa(utils.GetNowTime()) + ",'" + amount + "','" + amountAll + "','" + reward + "','" + ended + "'),"
			}
			session := models.MyEngine[og.platform]
			sqls := ""
			for i := 0; i < 10; i++ {
				if len(sqlstrs[i]) != len(sqlstr) {
					sqls += sqlstrs[i][0:len(sqlstrs[i])-1] + ";"
				}
			}
			_, err := session.Exec(sqls)
			if err != nil {
				fmt.Println("dddddddd" + err.Error())
			}
		}
	}
}
