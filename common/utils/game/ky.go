package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
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
	"time"
)

type KY struct {
	platform string
	kyconfig ramcache.KYConfig
}

func GetKy(platform string) KY {
	return KY{platform: platform, kyconfig: ramcache.GetKYConfig(platform)}
}

//发送请求
func postKY(apiurl string, params map[string]string, platform string) []byte {
	timestamp := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	kyconfig := ramcache.GetKYConfig(platform)
	paramsnew := make(map[string]string)
	paramsnew["agent"] = kyconfig.AGENT
	paramsnew["timestamp"] = timestamp
	var str = ""
	for k, v := range params {
		str += k + "=" + v + "&"
	}
	str = str[0 : len(str)-1]
	a := utils.AesEncrypt(str, kyconfig.DESKEY)
	paramsnew["param"] = url.QueryEscape(a)
	paramsnew["key"] = utils.MD5(kyconfig.AGENT + timestamp + kyconfig.MD5KEY)
	result := []byte("")
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	httpBuildQuery := ""
	for k, v := range paramsnew {
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
	apiurl = apiurl + "?" + httpBuildQuery
	req, err := http.NewRequest("GET", apiurl, strings.NewReader(""))
	if err != nil {
		return result
	}
	//利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	//给一个key设定为响应的value.
	resp, err := client.Do(req)
	if err != nil {
		return result
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("{}")
	}
	fmt.Println(string(body))
	return body
}

//获取游戏列表
func (ky KY) getGameList(platform string) error {
	return nil
}

//创建游戏账号
func (ky KY) createPlayer(userid string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	params := make(map[string]string)
	params["member_code"] = utils.MD5By16(platform + userid + userNameConst)
	params["password"] = utils.MD5(platform + userid + userPwdConst)
	userId, _ := strconv.Atoi(userid)
	account := xorm.PlatformAccounts{PlatId: platId, UserId: userId, Username: params["member_code"], Password: params["password"], Created: utils.GetNowTime()}
	_, err := models.MyEngine[platform].Insert(&account)
	if err != nil {
		return account, false
	}
	return account, true
}

//获取游戏登录url
func (ky KY) getGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string) string {
	apiurl := ky.kyconfig.KYURL
	params := make(map[string]string)
	params["s"] = "0"
	params["account"] = accounts.Username
	params["money"] = "0"
	params["orderid"] = ky.kyconfig.AGENT + utils.GetFmtTime() + accounts.Username
	params["ip"] = ip
	params["lineCode"] = ky.kyconfig.LINE
	params["KindID"] = gamecode
	content := postKY(apiurl, params, ky.platform)
	if bytes.Contains(content, []byte("\"code\":0")) {
		data, t, _, err := jsonparser.Get(content, "d")
		if err == nil && t.String() == "object" {
			res := make(map[string]string)
			json.Unmarshal(data, &res)
			return res["url"]
		}
	}
	return ""
}

//玩家存取款 amount 单位元
func (ky KY) uchips(username string, exId string, amount string) bool {
	apiurl := ky.kyconfig.KYURL
	params := make(map[string]string)
	amount2, _ := decimal.NewFromString(amount)
	if amount2.GreaterThan(decimal.Zero) {
		params["s"] = "2"
		params["money"] = amount2.String()
	} else {
		params["s"] = "3"
		params["money"] = amount2.Mul(decimal.New(-1, 0)).String()
	}
	params["account"] = username

	params["orderid"] = ky.kyconfig.AGENT + utils.GetFmtTime() + username
	content := postKY(apiurl, params, ky.platform)
	if bytes.Contains(content, []byte("\"code\":0")) {
		return true
	}
	return false
}

//删除玩家会话
func (ky KY) delSession(username string) bool {
	apiurl := ky.kyconfig.KYURL
	params := make(map[string]string)
	params["s"] = "8"
	params["account"] = username
	content := postKY(apiurl, params, ky.platform)
	if bytes.Contains(content, []byte("\"code\":0")) {
		return true
	}
	return false
}

//查询玩家余额
func (ky KY) queryUchips(username string) (string, bool) {
	apiurl := ky.kyconfig.KYURL
	params := make(map[string]string)
	params["s"] = "7"
	params["account"] = username
	content := postKY(apiurl, params, ky.platform)
	if bytes.Contains(content, []byte("\"code\":0")) {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		balance := res["d"].(map[string]interface{})["totalMoney"].(float64)
		return decimal.NewFromFloat(balance).String(), true
	}
	return "0", false
}

func (ky KY) GetBets() {
	bk, _ := ramcache.TableBetsKey.Load(ky.platform)
	betsKey := bk.(map[string]string)
	lasttime, _ := strconv.Atoi(betsKey["4-"])
	//第一次启动，自动抓取之前一个小时的数据
	now := utils.GetNowTime()
	if lasttime == 0 {
		lasttime = now - 3600*4
	}
	//往前多拉取两分钟的数据，避免第三方平台延时
	lasttime = lasttime - 120
	forCount := 1
	if now-lasttime > 3600 {
		forCount = (now-lasttime)/3600 + 1
		times := make(map[int][]int)
		fromtime := lasttime
		totime := lasttime + 3600
		for i := 0; i < forCount-1; i++ {
			fromtime = lasttime + 3600*i
			totime = fromtime + 3600 - 1
			time := []int{fromtime, totime}
			times[i] = time
		}
		times[forCount-1] = []int{totime, now}
		for i := 0; i < forCount; i++ {
			ky.getBetsByTime(times[i])
			time.Sleep(time.Second * 10)
		}
	} else {
		ky.getBetsByTime([]int{lasttime + 1, now})
	}
}

func (ky KY) getBetsByTime(times []int) {
	apiurl := ky.kyconfig.KYBETURL
	params := make(map[string]string)
	params["s"] = "6"
	params["startTime"] = strconv.Itoa(times[0]) + "000"
	params["endTime"] = strconv.Itoa(times[1]) + "999"
	content := postKY(apiurl, params, ky.platform)
	sqlstr := "insert ignore into bets_0 (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values"
	sqlstrs := make([]string, 0)
	for i := 0; i < 10; i++ {
		sqlstrs = append(sqlstrs, "insert ignore into bets_"+strconv.Itoa(i)+" (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values")
	}
	if bytes.Contains(content, []byte("\"code\":16")) {
		models.MyEngine[ky.platform].Exec("update bets_key set search_key=? where plat_id=? ", times[1], 4)
		bk, _ := ramcache.TableBetsKey.Load(ky.platform)
		betsKey := bk.(map[string]string)
		betsKey["4-"] = strconv.Itoa(times[1])
		ramcache.TableBetsKey.Store(ky.platform, betsKey)
		return
	}
	if bytes.Contains(content, []byte("\"code\":0")) {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		count := int(res["d"].(map[string]interface{})["count"].(float64))
		platAcc, _ := ramcache.TablePlatformAccounts.Load(ky.platform)
		if count > 0 {
			for i := 0; i < count; i++ {
				list := res["d"].(map[string]interface{})["list"].(map[string]interface{})
				orderId := list["GameID"].([]interface{})[i].(string)
				accountName := strings.Replace(list["Accounts"].([]interface{})[i].(string), ky.kyconfig.AGENT+"_", "", 1)
				//根据账号获取对应的user_id
				gameCode := strconv.FormatFloat(list["KindID"].([]interface{})[i].(float64), 'f', 0, 64)
				ended, _ := time.Parse("2006-01-02 15:04:05", list["GameEndTime"].([]interface{})[i].(string))
				amount := list["CellScore"].([]interface{})[i].(string)
				amountAll := list["AllBet"].([]interface{})[i].(string)
				reward := list["Profit"].([]interface{})[i].(string)
				userId := platAcc.(map[string]int)[accountName]
				amountD, _ := decimal.NewFromString(amount)
				rewardD, _ := decimal.NewFromString(reward)
				newReward := amountD.Add(rewardD).String()
				if userId == 0 {
					ramcache.UpdateTablePlatformAccounts(accountName, ky.platform, models.MyEngine[ky.platform])
					platAcc, _ := ramcache.TablePlatformAccounts.Load(ky.platform)
					userId = platAcc.(map[string]int)[accountName]
				}
				if userId == 0 {
					continue
				}
				sqlstrs[userId%10] += "('" + orderId + "','" + accountName + "','" + gameCode + "'," + strconv.Itoa(userId) + ",4," + strconv.Itoa(utils.GetNowTime()) + ",'" + amount + "','" + amountAll + "','" + newReward + "','" + strconv.Itoa(int(ended.Add(time.Hour*-8).Unix())) + "'),"
			}
			session := models.MyEngine[ky.platform]
			sqls := ""
			for i := 0; i < 10; i++ {
				if len(sqlstrs[i]) != len(sqlstr) {
					sqls += sqlstrs[i][0:len(sqlstrs[i])-1] + ";"
				}
			}
			if sqls == "" {
				return
			}
			_, err := session.Exec(sqls)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			_, err = models.MyEngine[ky.platform].Exec("update bets_key set search_key=? where plat_id=? ", times[1], 4)
			if err == nil {
				bk, _ := ramcache.TableBetsKey.Load(ky.platform)
				betsKey := bk.(map[string]string)
				betsKey["4-"] = strconv.Itoa(times[1])
				ramcache.TableBetsKey.Store(ky.platform, betsKey)
			}
		}
	}
}
