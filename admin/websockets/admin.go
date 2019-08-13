package websockets

import (
	"strconv"

	"qpgame/models"

	"github.com/kataras/iris/websocket"
)

//此文件的函数处理 PC后台 => 服务端 的websocket操作请求

//后台初始化连接
func initConnection(c *websocket.Connection, id int) {
	platform := (*c).Context().Params().Get("platform")
	if !(*c).IsJoined(platform) { //如果还没有加入，则加入
		(*c).Join(platform)
	}
	chargeTotal, withdrawTotal, silverMerchantTotal := GetCountInfo(models.MyEngine[platform])
	//拿到声音开关的信息
	conn := models.MyEngine[platform]
	sql := "SELECT charge_alert, withdraw_alert FROM admins WHERE id = '" + strconv.Itoa(id) + "' LIMIT 1"
	rows, rowsErr := conn.SQL(sql).QueryString()
	if rowsErr != nil || len(rows) == 0 {
		responseError(c, "后台用户信息不存在")
		return
	}
	row := rows[0]
	chargeAlert, _ := strconv.Atoi(row["charge_alert"])
	withdrawAlert, _ := strconv.Atoi(row["withdraw_alert"])
	config := map[string]interface{}{
		"sound_charge":          "https://d2al1bcutpfcks.cloudfront.net/S3_file_2018_06_13_2018061321245589999.mp3", //入款提醒
		"sound_withdraw":        "https://d2al1bcutpfcks.cloudfront.net/S3_file_2018_06_13_2018061321250073653.mp3", //出款提醒
		"charge_count":          chargeTotal,
		"withdraw_count":        withdrawTotal,
		"silver_merchant_count": silverMerchantTotal,
		"charge_alert":          chargeAlert,
		"withdraw_alert":        withdrawAlert,
	}
	ResponseData(c, "init", config)
}

//后台pc前端连接来的初始化
func statistic(c *websocket.Connection) {
	platform := (*c).Context().Params().Get("platform")
	if (*c).IsJoined(platform) { //如果还没有加入房间，则加入房间
		(*c).Join(platform)
	}
	chargeTotal, withdrawTotal, silverMerchantTotal := GetCountInfo(models.MyEngine[platform])
	data := map[string]int{
		"charge_count":          chargeTotal,
		"withdraw_count":        withdrawTotal,
		"silver_merchant_count": silverMerchantTotal,
	}
	ResponseData(c, "statistic", data)
}

//广播提现信息
func charge(c *websocket.Connection) {
	platform := (*c).Context().Params().Get("platform")
	if (*c).IsJoined(platform) { //如果还没有加入房间，则加入房间
		(*c).Join(platform)
	}
	chargeTotal, withdrawTotal, silverMerchantTotal := GetCountInfo(models.MyEngine[platform])
	data := map[string]int{
		"charge_count":          chargeTotal,
		"withdraw_count":        withdrawTotal,
		"silver_merchant_count": silverMerchantTotal,
	}
	ResponseData(c, "charge", data)
}

//广播提现信息
func withdraw(c *websocket.Connection) {
	platform := (*c).Context().Params().Get("platform")
	if (*c).IsJoined(platform) { //如果还没有加入房间，则加入房间
		(*c).Join(platform)
	}
	chargeTotal, withdrawTotal, silverMerchantTotal := GetCountInfo(models.MyEngine[platform])
	data := map[string]int{
		"charge_count":          chargeTotal,
		"withdraw_count":        withdrawTotal,
		"silver_merchant_count": silverMerchantTotal,
	}
	ResponseData(c, "withdraw", data)
}

//充值声音处理
func soundCharge(c *websocket.Connection, id int) {
	platform := (*c).Context().Params().Get("platform")
	conn := models.MyEngine[platform]
	idStr := strconv.Itoa(id)
	sql := "SELECT charge_alert FROM admins WHERE id = '" + idStr + "' LIMIT 1"
	rows, rowsErr := conn.SQL(sql).QueryString()
	if rowsErr != nil || len(rows) == 0 {
		responseError(c, "后台用户不存在")
		return
	}
	row := rows[0]
	toStatus := 2
	if row["charge_alert"] == "0" {
		toStatus = 1
	} else {
		toStatus = 0
	}
	sql = "UPDATE admins SET charge_alert = '" + strconv.Itoa(toStatus) + "' WHERE id = '" + idStr + "' LIMIT 1"
	result, resErr := conn.Exec(sql)
	if resErr != nil {
		responseError(c, "变更充值提醒状态失败")
		return
	}
	affected, affErr := result.RowsAffected()
	if affErr != nil || affected <= 0 {
		responseError(c, "充值提醒状态未变更成功")
		return
	}
	ResponseData(c, "sound_charge", map[string]int{
		"status": toStatus,
	})
}

//提现声音处理
func soundWithdraw(c *websocket.Connection, id int) {
	platform := (*c).Context().Params().Get("platform")
	conn := models.MyEngine[platform]
	idStr := strconv.Itoa(id)
	sql := "SELECT withdraw_alert FROM admins WHERE id = '" + idStr + "' LIMIT 1"
	rows, rowsErr := conn.SQL(sql).QueryString()
	if rowsErr != nil || len(rows) == 0 {
		responseError(c, "后台用户不存在")
		return
	}
	row := rows[0]
	toStatus := 2
	if row["withdraw_alert"] == "0" {
		toStatus = 1
	} else {
		toStatus = 0
	}
	sql = "UPDATE admins SET withdraw_alert = '" + strconv.Itoa(toStatus) + "' WHERE id = '" + idStr + "' LIMIT 1"
	result, resErr := conn.Exec(sql)
	if resErr != nil {
		responseError(c, "变更充值提醒状态失败")
		return
	}
	affected, affErr := result.RowsAffected()
	if affErr != nil || affected <= 0 {
		responseError(c, "充值提醒状态未变更成功")
		return
	}
	ResponseData(c, "sound_withdraw", map[string]int{
		"status": toStatus,
	})
}
