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

type FG struct {
	platform string
}

func GetFG(platform string) FG {
	return FG{platform: platform}
}

//发送请求
func postFG(apiurl string, params map[string]string, platform string) []byte {
	fgcof := ramcache.GetFGConfig(platform)
	apiurl = fgcof.FGURL + apiurl
	result := []byte("")
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
		return result
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")         //必须设定该参数,POST参数才能正常提交
	req.Header.Set("merchantname", fgcof.FgMerchantName) //必须设定该参数,POST参数才能正常提交
	req.Header.Set("merchantcode", fgcof.FgMerchantCode) //必须设定该参数,POST参数才能正常提交
	resp, err := client.Do(req)
	if err != nil {
		return result
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("{}")
	}
	return body
}

//获取游戏列表
func (fg FG) getGameList(platform string) error {
	apiurl := "v3/games/game_type/h5/language/zh-cn/"
	content := postFG(apiurl, nil, platform)
	var dats []*xorm.PlatformGames
	data, _, _, _ := jsonparser.Get(content, "data")
	err := json.Unmarshal(data, &dats)
	for _, g := range dats {
		g.PlatId = 1
	}
	_, err = models.MyEngine[platform].Delete(xorm.PlatformGames{PlatId: 1})
	if err == nil {
		go models.MyEngine[platform].Insert(dats)
	}
	return err
}

//创建游戏账号
func (fg FG) createPlayer(userid string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	apiurl := "v3/players"
	params := make(map[string]string)
	params["member_code"] = utils.MD5By16(platform + userid + userNameConst)
	params["password"] = utils.MD5(platform + userid + userPwdConst)
	content := postFG(apiurl, params, platform)
	userId, _ := strconv.Atoi(userid)
	account := xorm.PlatformAccounts{PlatId: platId, UserId: userId, Username: params["member_code"], Password: params["password"], Created: utils.GetNowTime()}
	if bytes.Contains(content, []byte("success")) {
		_, err := models.MyEngine[platform].Insert(&account)
		if err != nil {
			return account, false
		}
		return account, true
	}
	return account, false
}

//获取游戏登录url
func (fg FG) getGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string) string {
	apiurl := "v3/launch_game/"
	params := make(map[string]string)
	params["member_code"] = accounts.Username
	params["game_code"] = gamecode
	params["game_type"] = "h5"
	params["language"] = "zh-cn"
	params["ip"] = ip
	content := postFG(apiurl, params, fg.platform)
	if bytes.Contains(content, []byte("success")) {
		data, t, _, err := jsonparser.Get(content, "data")
		if err == nil && t.String() == "object" {
			res := make(map[string]string)
			json.Unmarshal(data, &res)
			return res["game_url"] + "&token=" + res["token"]
		}
	}
	return ""
}

//玩家存取款 amount 单位分
func (fg FG) uchips(username string, exId string, amount string) bool {
	apiurl := "v3/player_uchips/member_code/" + username
	params := make(map[string]string)
	params["member_code"] = username
	amount2, _ := decimal.NewFromString(amount)
	params["amount"] = amount2.Mul(decimal.New(100, 0)).String()
	params["externaltransactionid"] = exId
	content := postFG(apiurl, params, fg.platform)
	if bytes.Contains(content, []byte("success")) {
		return true
	}
	return false
}

//删除玩家会话
func (fg FG) delFgSession(username string) bool {
	apiurl := "v3/player_sessions/member_code/" + username
	params := make(map[string]string)
	params["member_code"] = username
	content := postFG(apiurl, params, fg.platform)
	if bytes.Contains(content, []byte("success")) {
		return true
	}
	return false
}

//查询玩家余额
func (fg FG) queryUchips(username string) (string, bool) {
	apiurl := "v3/player_chips/member_code/" + username
	params := make(map[string]string)
	params["member_code"] = username
	fg.delFgSession(username)
	content := postFG(apiurl, params, fg.platform)
	if bytes.Contains(content, []byte("success")) {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		balance := res["data"].(map[string]interface{})["balance"].(float64)
		return decimal.New(int64(balance), -2).String(), true
	}
	return "0", false
}

func (fg FG) GetBets(gt string) {
	bk, _ := ramcache.TableBetsKey.Load(fg.platform)
	betsKey := bk.(map[string]string)
	apiurl := "v3/agent/log_by_page/gt/" + gt
	if betsKey["1-"+gt] != "" {
		apiurl += "/page_key/" + betsKey["1-"+gt]
	}
	params := make(map[string]string)
	content := postFG(apiurl, params, fg.platform)
	if bytes.Contains(content, []byte("success")) {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		pageKey := res["data"].(map[string]interface{})["page_key"].(string)
		if pageKey == "none" {
			return
		}
		bets := res["data"].(map[string]interface{})["data"].([]interface{})
		fmt.Println(gt + "-----------------" + pageKey)
		sqlstr := "insert ignore into bets_0 (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values"
		sqlstrs := make([]string, 0)
		for i := 0; i < 10; i++ {
			sqlstrs = append(sqlstrs, "insert ignore into bets_"+strconv.Itoa(i)+" (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,,reward,ented)values")
		}
		for _, val := range bets {
			v := val.(map[string]interface{})
			orderId := v["id"].(string)
			accountName := v["player_name"].(string)
			//根据账号获取对应的user_id
			gameCode := strconv.FormatFloat(v["game_id"].(float64), 'f', 0, 64)
			ended := int(v["time"].(float64))
			amount := strconv.FormatFloat(v["all_bets"].(float64), 'f', 2, 64)
			reward := strconv.FormatFloat(v["all_wins"].(float64), 'f', 2, 64)
			platAcc, _ := ramcache.TablePlatformAccounts.Load(fg.platform)
			userId := platAcc.(map[string]int)[accountName]
			if userId == 0 {
				ramcache.UpdateTablePlatformAccounts(accountName, fg.platform, models.MyEngine[fg.platform])
				platAcc, _ := ramcache.TablePlatformAccounts.Load(fg.platform)
				userId = platAcc.(map[string]int)[accountName]
			}
			sqlstrs[userId%10] += "('" + orderId + "','" + accountName + "','" + gameCode + "'," + strconv.Itoa(userId) + ",1," + strconv.Itoa(utils.GetNowTime()) + ",'" + amount + "','" + amount + "','" + reward + "','" + strconv.Itoa(ended) + "'),"
			//bet := xorm.Bets{PlatformId: 1, OrderId: orderId, Accountname: accountName, GameCode: gameCode, Ented: ended, Gt: gt, Amount: amount, Reward: reward, Created: utils.GetNowTime(), UserId: userId}
		}
		session := models.MyEngine[fg.platform].NewSession()
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
		_, err = session.Exec("update bets_key set search_key=? where plat_id=? and gt=?", pageKey, 1, gt)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
		err = session.Commit()
		if err == nil {
			betsKey["1-"+gt] = pageKey
			ramcache.TableBetsKey.Store(fg.platform, betsKey)
		}
	}
}
