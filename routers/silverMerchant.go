package routers

import (
	"github.com/kataras/iris"
	"qpgame/app/silverMerchant"
	"qpgame/middlewares"
)

//路由配置
func SilverMerchantRouters(app *iris.Application) {
	//启动配置信息获取
	noPlatform := app.Party(
		"/silverMerchant",
		middlewares.CorsMiddleware(), //跨域中间件
	).AllowMethods(iris.MethodOptions)
	{
		noPlatform.Get("/getConfig", func(ctx iris.Context) { silverMerchant.NewApiConfigController(ctx).GetConfig() })
	}

	//前置中间件
	apiPlatform := app.Party(
		"/{platform:string}/silverMerchant/",
		middlewares.CorsMiddleware(),   //跨域中间件
		middlewares.BeforeMiddleware(), //前置中间件
	).AllowMethods(iris.MethodOptions)
	//非登录路由
	apiv1 := apiPlatform.Party("/api/v1")
	{
		apiv1.Get("/index/index", func(ctx iris.Context) {})
		//登陆
		apiv1.Post("/login", func(ctx iris.Context) { silverMerchant.NewSilverUserController(ctx).Login() })
		//获取验证码
		apiv1.Get("/verify", func(ctx iris.Context) { silverMerchant.NewSilverUserController(ctx).GetVerify() })
	}

	//需要登录之后才能访问的路由,//jwt中间件
	apiAuthV1 := apiPlatform.Party("/api/auth/v1", middlewares.JwtAuthenticate().Serve, middlewares.SilverMerchantHandler())
	{
		//会员充值
		apiAuthV1.Post("/pay", func(ctx iris.Context) { silverMerchant.NewMemberPayController(ctx).PayPlayer() })
		//会员搜索
		apiAuthV1.Post("/searchMember", func(ctx iris.Context) { silverMerchant.NewMemberPayController(ctx).CheckUserName() })
		//登出
		apiAuthV1.Post("/logout", func(ctx iris.Context) { silverMerchant.NewSilverUserController(ctx).LogOut() })
		//修改密码
		apiAuthV1.Post("/modifyPwd", func(ctx iris.Context) { silverMerchant.NewSilverUserController(ctx).ModifyPassWord() })
		//登录日志
		apiAuthV1.Post("/loginLog", func(ctx iris.Context) { silverMerchant.NewSilverLogController(ctx).LoginLog() })
		//操作日志
		apiAuthV1.Get("/operationLog", func(ctx iris.Context) { silverMerchant.NewSilverLogController(ctx).OperationLog() })
		// 银商会员充值记录查询接口O
		apiAuthV1.Get("/payRecords", func(ctx iris.Context) { silverMerchant.NewMemberPayController(ctx).PayForMemberRecords() })
		//银商首页
		apiAuthV1.Get("/welcome", func(ctx iris.Context) { silverMerchant.NewSilverUserController(ctx).Welcome() })
		//银商报表统计接口
		apiAuthV1.Get("/report", func(ctx iris.Context) { silverMerchant.NewSilverUserController(ctx).Report() })

		//银商充值账号获取接口
		apiAuthV1.Get("/getChargeCard", func(ctx iris.Context) { silverMerchant.NewSilverCardController(ctx).GetChargeCard() })
		//银商获取银行卡信息接口
		apiAuthV1.Get("/getSilverCardInfo", func(ctx iris.Context) { silverMerchant.NewSilverCardController(ctx).GetSilverCardInfo() })
		//银商绑定银行卡接口
		apiAuthV1.Post("/bindCard", func(ctx iris.Context) { silverMerchant.NewSilverCardController(ctx).BindCard() })
		//银商充值额度接口
		apiAuthV1.Post("/chargeSilver", func(ctx iris.Context) { silverMerchant.NewSilverCardController(ctx).ChargeSilver() })
		//银商充值记录查询接口
		apiAuthV1.Get("/getChargeSilverList", func(ctx iris.Context) { silverMerchant.NewSilverCardController(ctx).GetChargeSilverList() })
	}
}
