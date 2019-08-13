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

type VG struct {
	platform string
	vgConfig ramcache.VGConfig
}

func GetVG(platform string) VG {
	return VG{platform: platform, vgConfig: ramcache.GetVGConfig(platform)}
}

//发送请求
func postVG(apiurl string, params map[string]string, platform string) []byte {
	vgcof := ramcache.GetVGConfig(platform)
	params["channel"] = vgcof.CHANNEL
	result := []byte("")
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	httpBuildQuery := ""
	values := ""
	for k, v := range params {
		//如果传进来的是已经拼接好的，就放入map,k的值就是拼接好的,value为空字符串
		if len(params) == 1 && v == "" {
			httpBuildQuery = k
		} else {
			httpBuildQuery += k + "=" + v + "&"
			values += v
		}
	}
	values += vgcof.PASSWORD
	httpBuildQuery += "verifyCode=" + strings.ToUpper(utils.MD5(values))
	apiurl = apiurl + "?" + httpBuildQuery
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println(err.Error())
		return result
	}
	resp, err := client.Do(req)
	if err != nil {
		return result
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("")
	}
	return body
}

//获取游戏列表
func (vg VG) getGameList(platform string) error {
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
func (vg VG) createPlayer(userid string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	params := make(map[string]string)
	params["action"] = "create"
	params["username"] = utils.MD5(platform + userid + userNameConst)
	content := postVG(vg.vgConfig.URL, params, platform)
	userId, _ := strconv.Atoi(userid)
	account := xorm.PlatformAccounts{PlatId: platId, UserId: userId, Username: params["username"], Password: utils.MD5(platform + userid + userPwdConst), Created: utils.GetNowTime()}
	if bytes.Contains(content, []byte("<errcode>0</errcode>")) {
		_, err := models.MyEngine[platform].Insert(&account)
		if err != nil {
			return account, false
		}
		return account, true
	}
	return account, false
}

//获取游戏登录url
func (vg VG) getGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string) string {
	params := make(map[string]string)
	params["action"] = "loginWithChannel"
	params["username"] = accounts.Username
	params["gametype"] = gamecode
	params["gameversion"] = "2"
	content := postVG(vg.vgConfig.URL, params, vg.platform)
	if bytes.Contains(content, []byte("<errcode>0</errcode>")) {
		res := utils.GetXmlText(content, "result")
		return res["result"]
	}
	return ""
}

//玩家存取款 amount 单位分
func (vg VG) uchips(username string, exId string, amount string) bool {
	params := make(map[string]string)
	params["username"] = username
	amount2, _ := decimal.NewFromString(amount)
	amountOut := amount2.GreaterThan(decimal.Zero)
	vg.delSession(username)
	if amountOut {
		params["action"] = "deposit"
	} else {
		params["action"] = "withdraw"
		amount2 = amount2.Mul(decimal.New(-1, 0))
	}
	params["amount"] = amount2.String()
	params["serial"] = exId
	content := postVG(vg.vgConfig.URL, params, vg.platform)
	if bytes.Contains(content, []byte("<errcode>0</errcode>")) {
		return true
	} else {
		vg.delSession(username)
		content = postVG(vg.vgConfig.URL, params, vg.platform)
		if bytes.Contains(content, []byte("<errcode>0</errcode>")) {
			return true
		}
	}
	return false
}

//删除玩家会话
func (vg VG) delSession(username string) bool {
	params := make(map[string]string)
	params["action"] = "resetuser"
	params["username"] = username
	content := postVG(vg.vgConfig.URL, params, vg.platform)
	if bytes.Contains(content, []byte("success")) {
		return true
	}
	return false
}

//查询玩家余额
func (vg VG) queryUchips(username string) (string, bool) {
	params := make(map[string]string)
	params["action"] = "balance"
	params["username"] = username
	content := postVG(vg.vgConfig.URL, params, vg.platform)
	if bytes.Contains(content, []byte("<errcode>0</errcode>")) {
		res := utils.GetXmlText(content, "result", "rate", "coins", "in_game")
		rate, _ := decimal.NewFromString(res["rate"])
		coins, _ := decimal.NewFromString(res["coins"])
		balance := coins.Div(rate).String()
		return balance, true
	}
	return "0", false
}

func (vg VG) GetBets() {
	bk, _ := ramcache.TableBetsKey.Load(vg.platform)
	betsKey := bk.(map[string]string)
	params := make(map[string]string)
	if betsKey["8-"] != "" {
		params["id"] = betsKey["8-"]
	} else {
		params["id"] = "0"
	}
	content := postVG(vg.vgConfig.BETURL, params, vg.platform)
	state, _ := jsonparser.GetInt(content, "state")
	message, _ := jsonparser.GetString(content, "message")
	messageNum, _ := strconv.Atoi(message)
	resCounts := 0
	if state == 0 {
		res := make(map[string]interface{})
		json.Unmarshal(content, &res)
		pageKey := "0"
		bets := res["value"].([]interface{})
		sqlstr := "insert ignore into bets_0 (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values"
		sqlstrs := make([]string, 0)
		for i := 0; i < 10; i++ {
			sqlstrs = append(sqlstrs, "insert ignore into bets_"+strconv.Itoa(i)+" (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values")
		}
		resCounts = len(bets)
		if resCounts == 0 {
			return
		}
		for _, val := range bets {
			v := val.(map[string]interface{})
			orderId := v["id"].(string)
			pageKey = orderId
			accountName := v["username"].(string)
			//根据账号获取对应的user_id
			gameCode := v["gametype"].(string)
			var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
			endtime, _ := time.ParseInLocation("2006/1/2 15:04:05", v["endtime"].(string), cstSh)
			ended := strconv.Itoa(int(endtime.Unix()))
			amount := v["validbetamount"].(string)
			amountAll := v["betamount"].(string)
			amountAllD, _ := decimal.NewFromString(amountAll)
			money, _ := decimal.NewFromString(v["money"].(string))
			reward := amountAllD.Add(money).String()
			platAcc, _ := ramcache.TablePlatformAccounts.Load(vg.platform)
			userId := platAcc.(map[string]int)[accountName]
			if userId == 0 {
				ramcache.UpdateTablePlatformAccounts(accountName, vg.platform, models.MyEngine[vg.platform])
				platAcc, _ := ramcache.TablePlatformAccounts.Load(vg.platform)
				userId = platAcc.(map[string]int)[accountName]
			}
			sqlstrs[userId%10] += "('" + orderId + "','" + accountName + "','" + gameCode + "'," + strconv.Itoa(userId) + ",8," + strconv.Itoa(utils.GetNowTime()) + ",'" + amount + "','" + amountAll + "','" + reward + "','" + ended + "'),"
		}
		session := models.MyEngine[vg.platform].NewSession()
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
		_, err = session.Exec("update bets_key set search_key=? where plat_id=?", pageKey, 8)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
		err = session.Commit()
		if err == nil {
			betsKey["8-"] = pageKey
			ramcache.TableBetsKey.Store(vg.platform, betsKey)
		}
	}
	if resCounts == messageNum {
		time.Sleep(time.Second * 10)
		vg.GetBets()
	}
}
