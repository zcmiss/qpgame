package utils

import (
	"encoding/json"
	"qpgame/config"
	"time"

	"github.com/kataras/iris"
)

/*
 internalMsg 内部错误给程序员看的,开发和定位问题的时候特别有用
 clientMsg 展示给客户看的
*/

//失败响应处理
func ResFaiJSON(c *iris.Context, internalMsg string, clientMsg string, code int16) {
	currentTime := time.Now().UnixNano() / 1e3 //微秒
	diffTime := (*c).Values().GetInt64Default("requestCurrentTime", currentTime)
	timeConsumed := currentTime - diffTime
	result := iris.Map{"code": code, "clientMsg": clientMsg, "internalMsg": internalMsg, "timeConsumed": timeConsumed}
	(*c).JSON(result)
}

// 输入默认的cors
func ResponseCors(ctx *iris.Context) {
	(*ctx).Header("Access-Control-Allow-Credentials", "1")
	(*ctx).Header("Access-Control-Allow-Headers", "Authorization,Origin,Content-Type,Accept,X-Requested-With,PLATFORM")
	(*ctx).Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	(*ctx).Header("Access-Control-Allow-Origin", "*")
	(*ctx).Header("Content-Type", "application/json")
}

//失败响应处理
func ResFaiJSON2(c *iris.Context, error string, clientMsg string) {
	ResFaiJSON(c, error, clientMsg, config.NOTGETDATA)
}

/*
 internalMsg 内部错误给程序员看的,开发和定位问题的时候特别有用
 clientMsg 展示给客户看的
*/
//成功响应处理
func ResSuccJSON(c *iris.Context, internalMsg string, clientMsg string, code int16, data interface{}) {
	currentTime := time.Now().UnixNano() / 1e3 //微秒
	diffTime := (*c).Values().GetInt64Default("requestCurrentTime", currentTime)
	timeConsumed := currentTime - diffTime
	result := iris.Map{"code": code, "clientMsg": clientMsg, "internalMsg": internalMsg, "data": data, "timeConsumed": timeConsumed}
	res, _ := json.Marshal(result)

	(*c).Header("Content-Type", "application/json")
	if (len(res) > 5*1024) && (*c).ClientSupportsGzip() {
		(*c).WriteGzip(res)
	} else {
		(*c).Write(res)
	}

}
