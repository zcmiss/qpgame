/**
对接EA平台
*/
package game

import (
	"bytes"
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

type JDB struct {
	platform  string
	jdbconfig ramcache.JDBConfig
}

func GetJdb(platform string) JDB {
	return JDB{platform: platform, jdbconfig: ramcache.GetJDBConfig(platform)}
}

//正式环境
//const EAURL  = "https://api.fafafa3388.com"
//const EAKEY  = "mgUT6Pc7UxzfVaoRkXbKuW03BYrxLaQh"

//发送请求
func (jdb JDB) postJDB(params map[string]interface{}) []byte {
	apiurl := jdb.jdbconfig.URL
	result := []byte("")
	jsonParam, err := json.Marshal(params)
	fmt.Println(string(jsonParam) + "------------" + jdb.platform)
	if err != nil {
		return result
	}
	paramsMap := make(map[string]string)
	x := utils.PswEncrypt(string(jsonParam), jdb.jdbconfig.KEY, jdb.jdbconfig.IV)
	paramsMap["dc"] = jdb.jdbconfig.DC
	paramsMap["x"] = x
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	httpBuildQuery := ""
	for k, v := range paramsMap {
		//如果传进来的是已经拼接好的，就放入map,k的值就是拼接好的,value为空字符串
		if len(paramsMap) == 1 && v == "" {
			httpBuildQuery = k
		} else {
			httpBuildQuery += k + "=" + v + "&"
		}
	}
	if httpBuildQuery != "" {
		httpBuildQuery = strings.TrimRight(httpBuildQuery, "&")
	}
	fmt.Println(apiurl + "?" + httpBuildQuery)
	req, err := http.NewRequest("POST", apiurl+"?"+httpBuildQuery, nil)
	if err != nil {
		return result
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json") //必须设定该参数,POST参数才能正常提交
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return result
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("{}")
	}
	fmt.Println(string(body) + "------------" + jdb.platform)
	return body
}

//获取游戏列表
func (jdb JDB) getGameList(platform string) error {

	return nil
}

//创建游戏账号
func (jdb JDB) createPlayer(userid string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	params := make(map[string]interface{})
	params["action"] = 12
	params["ts"] = time.Now().UnixNano() / 1e6
	params["parent"] = jdb.jdbconfig.PARENT
	params["uid"] = utils.MD5By16(platform + userid + userNameConst)
	params["name"] = utils.MD5By16(platform + userid + userNameConst)
	content := jdb.postJDB(params)
	userId, _ := strconv.Atoi(userid)
	account := xorm.PlatformAccounts{PlatId: platId, UserId: userId, Username: params["uid"].(string), Password: params["uid"].(string), Created: utils.GetNowTime()}
	if bytes.Contains(content, []byte("\"status\":\"0000\"")) {
		_, err := models.MyEngine[platform].Insert(&account)
		if err != nil {
			return account, false
		}
		return account, true
	}
	return account, false
}

func getGType(platform, mType string) string {
	cache, _ := ramcache.TablePlatformGamesByJDB.Load(platform)
	for _, v := range cache.([]xorm.PlatformGames) {
		if v.Gamecode == mType {
			return v.Gt
		}
	}
	return "0"
}

//获取游戏登录url
func (jdb JDB) getGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string) string {
	params := make(map[string]interface{})
	params["action"] = 11
	params["ts"] = time.Now().UnixNano() / 1e6
	params["uid"] = accounts.Username
	params["gType"] = getGType(jdb.platform, gamecode)
	params["mType"] = gamecode
	params["windowMode"] = "1"
	params["isAPP"] = true
	params["lang"] = "ch"
	params["moreGame"] = 1
	params["mute"] = 0
	params["cardGameGroup"] = 0
	content := jdb.postJDB(params)
	if bytes.Contains(content, []byte("\"status\":\"0000\"")) {
		res := make(map[string]string)
		json.Unmarshal(content, &res)
		return res["path"]
	}
	return ""
}

func (jdb JDB) driveOutTheUser(username string) {
	params := make(map[string]interface{})
	params["parent"] = jdb.jdbconfig.PARENT
	params["uid"] = username
	params["action"] = 17
	params["ts"] = time.Now().UnixNano() / 1e6
	jdb.postJDB(params)
}

//玩家存取款 amount 单位元
func (jdb JDB) uchips(username string, exId string, amount string) bool {
	params := make(map[string]interface{})
	params["parent"] = jdb.jdbconfig.PARENT
	params["uid"] = username
	amount2, _ := decimal.NewFromString(amount)
	amountIn := amount2.LessThan(decimal.New(0, 0))
	if amountIn {
		jdb.driveOutTheUser(username)
	}
	params["ts"] = time.Now().UnixNano() / 1e6
	params["action"] = 19
	params["serialNo"] = exId
	params["allCashOutFlag"] = "0"
	params["amount"], _ = amount2.Float64()
	content := jdb.postJDB(params)
	if bytes.Contains(content, []byte("\"status\":\"0000\"")) || bytes.Contains(content, []byte("\"status\":\"6002\"")) {
		return true
	}
	//else if amountIn {
	//	cnt := 0
	//	for cnt < 3 {
	//		cnt++
	//		time.Sleep(1 * time.Second)
	//		content = jdb.postJDB(params)
	//		if bytes.Contains(content, []byte("\"status\":\"0000\"")) {
	//			return true
	//		}
	//	}
	//}
	return false
}

//查询玩家余额
func (jdb JDB) queryUchips(username string) (string, bool) {
	params := make(map[string]interface{})
	//params["action"] = 52
	params["ts"] = time.Now().UnixNano() / 1e6
	params["parent"] = jdb.jdbconfig.PARENT
	params["uid"] = username
	// update by aTian 提款时无需去判断使用者状态，如果使用者正在游戏当中无法提款
	// 已经和JDB客服核实过
	//content := jdb.postJDB(params)
	//if bytes.Contains(content, []byte("\"status\":\"0000\"")) {
	//	return "0", false
	//}
	params["action"] = 15
	content := jdb.postJDB(params)
	if bytes.Contains(content, []byte("\"status\":\"0000\"")) {
		data, _, _, _ := jsonparser.Get(content, "data")
		fmt.Println(string(data))
		res := make([]map[string]interface{}, 0)
		json.Unmarshal(data, &res)
		return strconv.FormatFloat(res[0]["balance"].(float64), 'f', 2, 64), true
	}
	return "0", false
}

//查询玩家余额
func (jdb JDB) GetBets() {
	bk, _ := ramcache.TableBetsKey.Load(jdb.platform)
	betsKey := bk.(map[string]string)
	lasttime, _ := strconv.Atoi(betsKey["9-"])
	//第一次启动，自动抓取之前一个小时的数据,最新数据为3分钟之前
	now := utils.GetNowTime() - 180
	if lasttime == 0 {
		lasttime = now - 3600*1
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
			jdb.getBetsByTime(times[i])
		}
	} else {
		jdb.getBetsByTime([]int{lasttime + 1, now})
	}
	_, err := models.MyEngine[jdb.platform].Exec("update bets_key set search_key=? where plat_id=? ", now, 9)
	if err == nil {
		betsKey["9-"] = strconv.Itoa(now)
		ramcache.TableBetsKey.Store(jdb.platform, betsKey)
	}

}

func (jdb JDB) getBetsByTime(times []int) {
	params := make(map[string]interface{})
	params["action"] = 29
	params["ts"] = time.Now().UnixNano() / 1e6
	params["parent"] = jdb.jdbconfig.PARENT
	startTime := time.Unix(int64(times[0]), int64(0))
	endtime := time.Unix(int64(times[1]), int64(0))
	params["starttime"] = startTime.Add(time.Hour * -12).Format("02-01-2006 15:04:00")
	params["endtime"] = endtime.Add(time.Hour * -12).Format("02-01-2006 15:04:00")
	if params["starttime"] == params["endtime"] {
		return
	}
	content := jdb.postJDB(params)
	sqlstr := "insert ignore into bets_0 (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values"
	sqlstrs := make([]string, 0)
	for i := 0; i < 10; i++ {
		sqlstrs = append(sqlstrs, "insert ignore into bets_"+strconv.Itoa(i)+" (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values")
	}
	if bytes.Contains(content, []byte("\"status\":\"0000\"")) {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		data := res["data"].([]interface{})
		if len(data) > 0 {
			for _, val := range data {
				v := val.(map[string]interface{})
				orderId := strconv.FormatFloat(v["seqNo"].(float64), 'f', 0, 64)
				accountName := v["playerId"].(string)
				//根据账号获取对应的user_id
				gameCode := strconv.FormatFloat(v["mtype"].(float64), 'f', 0, 64)
				var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
				ended, _ := time.ParseInLocation("02-01-2006 15:04:05", v["lastModifyTime"].(string), cstSh)
				amount := strconv.FormatFloat(v["bet"].(float64), 'f', 0, 64)[1:]
				reward := strconv.FormatFloat(v["win"].(float64), 'f', 0, 64)
				platAcc, _ := ramcache.TablePlatformAccounts.Load(jdb.platform)
				userId := platAcc.(map[string]int)[accountName]
				if userId == 0 {
					ramcache.UpdateTablePlatformAccounts(accountName, jdb.platform, models.MyEngine[jdb.platform])
					platAcc, _ := ramcache.TablePlatformAccounts.Load(jdb.platform)
					userId = platAcc.(map[string]int)[accountName]
				}
				sqlstrs[userId%10] += "('" + orderId + "','" + accountName + "','" + gameCode + "'," + strconv.Itoa(userId) + ",9," + strconv.Itoa(utils.GetNowTime()) + ",'" + amount + "','" + amount + "','" + reward + "','" + strconv.Itoa(int(ended.Add(time.Hour*12).Unix())) + "'),"
			}
			session := models.MyEngine[jdb.platform]
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
