package websockets

import (
	"encoding/json"

	"qpgame/models"

	"github.com/kataras/iris/websocket"
)

//此文件的函数处理 手机 => 服务端 的websocket操作请求

//对于充值的处理
func mobileCharge(c *websocket.Connection) {
	responseMessage(c, "", "mobile_charge")
	broadcastStatistic(c)
}

//对于提现的处理
func mobilewithdraw(c *websocket.Connection) {
	responseMessage(c, "", "mobile_charge")
	broadcastStatistic(c)
}

//向同一平台的所有后台连接进行广播
func broadcastStatistic(c *websocket.Connection) {
	platform := (*c).Context().Params().Get("platform")
	chargeTotal, withdrawTotal, silverMerchantTotal := GetCountInfo(models.MyEngine[platform])
	info := map[string]int{
		"charge_count":          chargeTotal,
		"withdraw_count":        withdrawTotal,
		"silver_merchant_count": silverMerchantTotal,
	}
	result := WsResult{
		Type:         "statistic",
		ClientMsg:    "",
		Data:         info,
		TimeConsumed: getTimeConsumed(c),
	}
	data, _ := json.Marshal(result)
	(*c).To(platform).EmitMessage(data)
}
