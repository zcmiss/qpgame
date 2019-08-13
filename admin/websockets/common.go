package websockets

import (
	"encoding/json"
	"strconv"
	"time"

	"qpgame/config"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/websocket"
)

var AdminWebSocket *websocket.Server

// 不同平台上每个用户显示的统计结果
var WSConnecition = map[string]map[*websocket.Connection]int{}

//启动后台的websocket处理
func Init() {
	AdminWebSocket = websocket.New(websocket.Config{
		ReadBufferSize:  40960,
		WriteBufferSize: 40960,
	})
	AdminWebSocket.OnConnection(handleConnection)
}

//处理websocket
func handleConnection(c websocket.Connection) {
	c.OnMessage(func(msg []byte) { // 处理客户端请求
		cmd := getUnpackData(msg)
		_, exists := config.PlatformCPs[cmd.Platform]
		if !exists {
			responseError(&c, "不存在的平台标识号:"+cmd.Platform)
		}
		c.Context().Params().Set("platform", cmd.Platform)                        //设置平台标识号
		c.Context().Values().Set("requestCurrentTime", time.Now().UnixNano()/1e3) //设置请求开始时间
		if _, ok := WSConnecition[cmd.Platform]; !ok {
			WSConnecition[cmd.Platform] = make(map[*websocket.Connection]int)
		}
		WSConnecition[cmd.Platform][&c] = 1
		switch cmd.Type {
		case "init":
			initConnection(&c, cmd.Id) //初始化连接
		case "statistic": //统计信息
			statistic(&c)
		case "heart":
			responseHeart(&c) //心跳-pc后台
		case "charge":
			charge(&c) //充值-手机端
		case "withdraw":
			withdraw(&c) //提现-手机端
		case "sound_charge": //{"type":"charge_sound", "user_id":"1", "turn":"off/on"}
			soundCharge(&c, cmd.Id) //关闭充值声音
		case "sound_withdraw": //{"type":"charge_sound", "user_id":"1", "turn":"off/on"}
			soundWithdraw(&c, cmd.Id) //关闭提现声音
		case "charge_phone":
			mobileCharge(&c)
		case "withdraw_phone":
			mobilewithdraw(&c)
		default:
			responseError(&c, "不存在的请求方法:"+cmd.Type)
		}
	})
	// 添加关闭的方法，从room删掉用户
	c.OnDisconnect(func() {
		platform := c.Context().Params().Get("platform")
		if c.IsJoined(platform) { //如果已进入过room则退出之
			c.Leave(platform)
		}
		if _, ok := WSConnecition[platform]; ok {
			if _, ok := WSConnecition[platform][&c]; ok {
				delete(WSConnecition[platform], &c)
			}
		}
	})
}

//心跳
func responseHeart(c *websocket.Connection) {
	platform := (*c).Context().Params().Get("platform")
	if !(*c).IsJoined(platform) {
		(*c).Join(platform)
	}
	data, _ := json.Marshal(map[string]string{
		"type": "hearted",
	})
	(*c).To((*c).ID()).EmitMessage(data)
}

// 返回错误的信息
func responseError(c *websocket.Connection, message string) {
	result := WsResult{
		Type:         "error",
		ClientMsg:    message,
		Data:         nil,
		TimeConsumed: getTimeConsumed(c),
	}
	data, _ := json.Marshal(result)
	(*c).To((*c).ID()).EmitMessage(data)
}

//返回带结果的数据
func ResponseData(c *websocket.Connection, typeStr string, data interface{}) {
	result := WsResult{
		Type:         typeStr,
		ClientMsg:    "",
		Data:         data,
		TimeConsumed: getTimeConsumed(c),
	}
	res, _ := json.Marshal(result)
	(*c).To((*c).ID()).EmitMessage(res)
}

//返回操作正确的数据
func responseMessage(c *websocket.Connection, message string, typeStr string) {
	(*c).Context().Params().Get("requestCurrentTime")
	result := WsResult{
		Type:         typeStr,
		ClientMsg:    message,
		Data:         nil,
		TimeConsumed: getTimeConsumed(c),
	}
	res, _ := json.Marshal(result)
	(*c).To((*c).ID()).EmitMessage(res)
}

//得到花费的时间
func getTimeConsumed(c *websocket.Connection) int64 {
	requestTime := (*c).Context().Values().GetInt64Default("requestCurrentTime", 0)
	currentTime := time.Now().UnixNano() / 1e3 //微秒
	return currentTime - requestTime
}

//获取统计信息: (充值数量, 提现数量)
func GetCountInfo(db *xorm.Engine) (int, int, int) {
	sql := "(SELECT 'charge' type,COUNT(*) total FROM charge_records WHERE is_tppay=0 AND state=0)" +
		" UNION ALL (SELECT 'withdraw' type,COUNT(*) total FROM withdraw_records WHERE status=0)" +
		" UNION ALL (SELECT 'silver_merchant' type,COUNT(*) total FROM silver_merchant_charge_records WHERE state=0)"
	rows, err := db.SQL(sql).QueryString()
	if err != nil {
		return 0, 0, 0
	}
	chargeTotal, withdrawTotal, silverMerchantTotal := 0, 0, 0
	for _, row := range rows {
		switch row["type"] {
		case "charge":
			count, err := strconv.Atoi(row["total"])
			if (err == nil) && (count > 0) {
				chargeTotal = count
			}
		case "withdraw":
			count, err := strconv.Atoi(row["total"])
			if (err == nil) && (count > 0) {
				withdrawTotal = count
			}
		case "silver_merchant":
			count, err := strconv.Atoi(row["total"])
			if (err == nil) && (count > 0) {
				silverMerchantTotal = count
			}
		}
	}
	return chargeTotal, withdrawTotal, silverMerchantTotal
}

//得到解包之后的数据: 命令类型,后台用户编号
func getUnpackData(msg []byte) WsCommand {
	data := WsCommand{}
	err := json.Unmarshal(msg, &data)
	if err == nil {
		return data
	}
	cmd := WsCommandNoId{}
	err = json.Unmarshal(msg, &cmd)
	if err != nil || cmd.Type == "" {
		return WsCommand{
			Platform: "NONAME",
			Type:     "",
			Id:       0,
		}
	}
	return WsCommand{
		Platform: cmd.Platform,
		Type:     cmd.Type,
		Id:       0,
	}
}
