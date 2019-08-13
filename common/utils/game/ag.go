package game

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"qpgame/common/log"
	"qpgame/common/utils"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"time"
)

type AG struct {
	platform string
	agconfig ramcache.AGConfig
}

func GetAg(platform string) AG {
	return AG{platform: platform, agconfig: ramcache.GetAGConfig(platform)}
}

//定义返回结果的结构体
type Result struct {
	Info string `xml:"info,attr"`
	Msg  string `xml:"msg,attr"`
}

//定义投注记录的结构体
type Rows struct {
	XMLName xml.Name `xml:"row"`
	row     []Row    `xml:"row"`
}
type Row struct {
	DataType string `xml:"dataType,attr"`
	//电子游戏和真人游戏属性
	BillNo         string `xml:"billNo,attr"`
	PlayerName     string `xml:"playerName,attr"`
	BetTime        string `xml:"betTime,attr"`
	GameType       string `xml:"gameType,attr"`
	BetAmount      string `xml:"betAmount,attr"`
	ValidBetAmount string `xml:"validBetAmount,attr"`
	NetAmount      string `xml:"netAmount,attr"`
	RecalcuTime    string `xml:"recalcuTime,attr"`
	//捕鱼属性
	TradeNo      string `xml:"tradeNo,attr"`
	SceneEndTime string `xml:"SceneEndTime,attr"`
	Cost         string `xml:"Cost,attr"`
	Earn         string `xml:"Earn,attr"`
	PlatformType string `xml:"platformType,attr"`
}

var password = "qazwsxecde"

//发送请求
func (ag AG) postAG(apiurl string, params map[string]string, platform string) []byte {
	paramsnew := make(map[string]string)
	var str = ""
	for k, v := range params {
		str += k + "=" + v + "/\\\\\\\\/"
	}
	str = str[0 : len(str)-len("/\\\\\\\\/")]
	a := utils.DesEncrypt(str, ag.agconfig.DESKEY)
	paramsnew["params"] = a
	paramsnew["key"] = utils.MD5(a + ag.agconfig.MD5KEY)
	result := []byte("")
	client := &http.Client{
		Timeout: time.Second * 35,
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
	if strings.Contains(apiurl, "forwardGame.do") {
		return []byte(apiurl)
	}
	req, err := http.NewRequest("POST", apiurl, strings.NewReader(""))
	if err != nil {
		return result
	}
	if strings.Contains(apiurl, "doBusiness.do") {
		req.Header.Set("User-Agent", "WEB_LIB_GI_"+ag.agconfig.CAGENT) //必须设定该参数,才能正常提交
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
	return body
}

//获取游戏列表
func (ag AG) getGameList(platform string) error {
	return nil
}

//创建游戏账号
func (ag AG) createPlayer(userid string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	apiurl := ag.agconfig.GIURL
	params := make(map[string]string)
	params["cagent"] = ag.agconfig.CAGENT
	params["loginname"] = ag.agconfig.CAGENTQ + userid
	//BBIN平台的密码须为6~12码英文或数字且符
	//**BBIN 平台的会员帐号(请输入 4-20 个字符, 仅可输入英文字母以及数字的组合)
	//**MG 平台的会员帐号(必须加上 cagent 前缀，如 cagent=AAA_NMGE,会员账号必須加上前綴
	//AAA，賬號就是 AAAxxxx)
	//***沙巴平台的会员帐号(仅可输入英文字母以及数字的组合)
	//***IPM 平台的会员帐号(不可超过 13 位元长度)
	params["password"] = password
	params["method"] = "lg"
	params["actype"] = "1"
	params["oddtype"] = "A"
	params["cur"] = "CNY"
	userId, _ := strconv.Atoi(userid)
	content := ag.postAG(apiurl, params, ag.platform)
	//如果
	var account xorm.PlatformAccounts
	if bytes.Contains(content, []byte("msg=\"\"")) {
		account := xorm.PlatformAccounts{PlatId: platId, UserId: userId, Username: params["loginname"], Password: params["password"], Created: utils.GetNowTime()}
		_, err := models.MyEngine[platform].Insert(&account)
		if err != nil {
			return account, false
		}
		return account, true
	}
	return account, false
}

//获取游戏登录url
func (ag AG) getGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string) string {
	apiurl := ag.agconfig.GCIURL
	params := make(map[string]string)
	params["cagent"] = ag.agconfig.CAGENT
	params["loginname"] = accounts.Username
	params["password"] = accounts.Password
	params["sid"] = ag.agconfig.CAGENT + strconv.Itoa(time.Now().Nanosecond())
	params["actype"] = "1"
	params["lang"] = "1"
	params["gameType"] = gamecode
	params["oddtype"] = "A"
	params["cur"] = "CNY"
	params["mh5"] = "y"
	content := ag.postAG(apiurl, params, ag.platform)
	return string(content)
}

//玩家存取款 amount 单位元
func (ag AG) uchips(username string, exId string, amount string) bool {
	apiurl := ag.agconfig.GIURL
	params := make(map[string]string)
	amount2, _ := decimal.NewFromString(amount)
	if amount2.GreaterThan(decimal.Zero) {
		params["type"] = "IN"
		params["credit"] = amount2.String()
	} else {
		params["type"] = "OUT"
		params["credit"] = amount2.Mul(decimal.New(-1, 0)).String()
	}
	params["cagent"] = ag.agconfig.CAGENT
	params["loginname"] = username
	params["method"] = "tc"
	//billno = (cagent+序列), 序列是唯一的 13~16 位数, 例如:
	//cagent = ‘XXXXX’ 及 序列 = 1234567890987, 那么 billno = XXXXX1234567890987,
	exIdlen := len(exId)
	if len(exId) < 13 {
		for i := 0; i < 13-exIdlen; i++ {
			exId = "0" + exId
		}
	}
	params["billno"] = ag.agconfig.CAGENT + exId
	params["actype"] = "1"
	params["password"] = password
	params["cur"] = "CNY"
	content := ag.postAG(apiurl, params, ag.platform)
	//如果
	if bytes.Contains(content, []byte("msg=\"\"")) {
		params["method"] = "tcc"
		params["flag"] = "1"
		content := ag.postAG(apiurl, params, ag.platform)
		if bytes.Contains(content, []byte("msg=\"\"")) {
			return true
		}
	}
	return false
}

//查询玩家余额
func (ag AG) queryUchips(username string) (string, bool) {
	apiurl := ag.agconfig.GIURL
	params := make(map[string]string)
	params["cagent"] = ag.agconfig.CAGENT
	params["loginname"] = username
	params["method"] = "gb"
	params["actype"] = "1"
	params["password"] = password
	params["cur"] = "CNY"
	content := ag.postAG(apiurl, params, ag.platform)
	if bytes.Contains(content, []byte("msg=\"\"")) {
		result := new(Result)
		xml.Unmarshal(content, &result)
		return result.Info, true
	}
	return "0", false
}

func (ag AG) GetBets() {
	gts := []string{"AGIN", "HUNTER", "XIN", "YOPLAY"}
	date := time.Now().Add(time.Hour * 12 * -1).Format("20060102")
	ch := make(chan int, 4)
	bk, _ := ramcache.TableBetsKey.Load(ag.platform)
	for _, v := range gts {
		betsKey := bk.(map[string]string)
		path := "/" + v + "/" + date + "/"
		go func(filePath string, betsKey string, plat string) {
			ch <- ag.readXml(filePath, betsKey, plat)
		}(path, betsKey["5-"+v], v)
	}
	//for i := range ch {
	//fmt.Println(i)
	//}

}

func (ag AG) readXml(path string, betskey string, plat string) int {
	sqlstrs := make([]string, 0)
	sqls := ""
	sqlstr := "insert ignore into bets_0 (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values"
	for i := 0; i < 10; i++ {
		sqlstrs = append(sqlstrs, "insert ignore into bets_"+strconv.Itoa(i)+" (order_id,accountname,game_code,user_id,platform_id,created,amount,amount_all,reward,ented)values")
	}
	conn, err := ftp.DialTimeout(ag.agconfig.FTPURL, 5*time.Second)
	if err != nil {
		log.Log.Fatal(err)
	}
	err = conn.Login(ag.agconfig.FTPNAME, ag.agconfig.FTPPWD)
	if err != nil {
		log.Log.Fatal(err)
	}
	list, err := conn.List(path)
	defer conn.Logout()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	maxname := ""
	for i, f := range list {
		if strings.Replace(f.Name, ".xml", "", 1) > maxname {
			maxname = strings.Replace(f.Name, ".xml", "", 1)
		}
		//每次都要读取最后12个文件
		//tj1 := len(list)-i <= 12
		if strings.Replace(f.Name, ".xml", "", 1) > betskey || len(list)-i <= 12 {
			resp, err := conn.Retr(path + f.Name)
			if err != nil {
				fmt.Println(err.Error())
			}
			//逐行读取，否则会出现数据缺失
			sc := bufio.NewScanner(resp)
			platAcc, _ := ramcache.TablePlatformAccounts.Load(ag.platform)
			for sc.Scan() {
				var row Row
				xml.Unmarshal(sc.Bytes(), &row)
				//电子和视讯的投注记录
				if row.DataType == "BR" || row.DataType == "EBR" {
					orderId := row.BillNo
					accountName := row.PlayerName
					//根据账号获取对应的user_id
					gameCode := row.GameType
					ended, _ := time.Parse("2006-01-02 15:04:05", row.RecalcuTime)
					ended = ended.Add(time.Hour * 12)
					amount := row.ValidBetAmount
					amountAll := row.BetAmount
					reward := row.NetAmount
					betAmountD, _ := decimal.NewFromString(row.BetAmount)
					rewardD, _ := decimal.NewFromString(reward)
					newReward := betAmountD.Add(rewardD).String()
					userId := platAcc.(map[string]int)[accountName]
					if userId == 0 {
						ramcache.UpdateTablePlatformAccounts(accountName, ag.platform, models.MyEngine[ag.platform])
						platAcc, _ := ramcache.TablePlatformAccounts.Load(ag.platform)
						userId = platAcc.(map[string]int)[accountName]
					}
					sqlstrs[userId%10] += "('" + orderId + "','" + accountName + "','" + gameCode + "'," + strconv.Itoa(userId) + ",5," + strconv.Itoa(utils.GetNowTime()) + ",'" + amount + "','" + amountAll + "','" + newReward + "','" + strconv.Itoa(int(ended.Unix())) + "'),"

				}
				//捕鱼的投注记录
				if row.DataType == "HSR" {
					orderId := row.TradeNo
					accountName := row.PlayerName
					//根据账号获取对应的user_id
					gameCode := row.PlatformType
					ended, _ := time.Parse("2006-01-02 15:04:05", row.SceneEndTime)
					ended = ended.Add(time.Hour * 12)
					amount := row.Cost
					reward := row.Earn
					userId := platAcc.(map[string]int)[accountName]
					if userId == 0 {
						ramcache.UpdateTablePlatformAccounts(accountName, ag.platform, models.MyEngine[ag.platform])
						platAcc, _ := ramcache.TablePlatformAccounts.Load(ag.platform)
						userId = platAcc.(map[string]int)[accountName]
					}
					sqlstrs[userId%10] += "('" + orderId + "','" + accountName + "','" + gameCode + "'," + strconv.Itoa(userId) + ",5," + strconv.Itoa(utils.GetNowTime()) + ",'" + amount + "','" + amount + "','" + reward + "','" + strconv.Itoa(int(ended.Unix())) + "'),"
				}
			}
			resp.Close()
		} else {
			conn.Logout()
		}
	}
	for i := 0; i < 10; i++ {
		if len(sqlstrs[i]) != len(sqlstr) {
			sqls += sqlstrs[i][0:len(sqlstrs[i])-1] + ";"
		}
	}
	if sqls == "" {
		return 0
	}
	go models.MyEngine[ag.platform].Exec(sqls)
	_, err = models.MyEngine[ag.platform].Exec("update bets_key set search_key=? where plat_id=? and gt=?", maxname, 5, plat)
	if err == nil {
		bk, _ := ramcache.TableBetsKey.Load(ag.platform)
		bmap := bk.(map[string]string)
		bmap["5-"+plat] = maxname
		ramcache.TableBetsKey.Store(ag.platform, bmap)
	}
	return 1
}
