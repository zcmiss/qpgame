package middlewares

import (
	"qpgame/common/utils"
	"qpgame/config"
	"time"

	"github.com/kataras/iris"
)

// 前置中间件 - 所有请求开始
func BeforeMiddleware() (handler iris.Handler) {
	platformcps := config.PlatformCPs
	handler = func(ctx iris.Context) {
		platform := ctx.Params().Get("platform")
		//请求时间开始计时
		ctx.Values().Set("requestCurrentTime", time.Now().UnixNano()/1e3)
		if _, existPlatform := platformcps[platform]; !existPlatform {
			utils.ResFaiJSON(&ctx, "没有这个平台号", "无法识别的平台号", config.NOTGETDATA)
			return
		}
		ctx.Next()
	}
	return handler
}
