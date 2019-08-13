package routers

import (
	"github.com/kataras/iris"
	"qpgame/app/frontEndControllers"
	"qpgame/app/frontEndControllers/pay"
	"qpgame/middlewares"
)

//路由配置
func AppConfigRouters(app *iris.Application) {
	//静态文件路径
	app.StaticWeb("/assets", "./assets")
	//注册模板文件路径
	app.RegisterView(iris.HTML("./views", ".html"))
	// @update by aTian 映射根路由至CK棋牌下载页
	app.Get("/", func(ctx iris.Context) {
		frontEndControllers.NewUserController(ctx).ToHomeDownload()
	})
	// 映射根路由至棋牌分享页
	app.Get("/download", func(ctx iris.Context) {
		frontEndControllers.NewUserController(ctx).ToDownloadAPP()
	})
	//前置中间件
	apiPlatform := app.Party(
		"/{platform:string}",
		middlewares.CorsMiddleware(),   //跨域中间件
		middlewares.BeforeMiddleware(), //前置中间件
	).AllowMethods(iris.MethodOptions)

	//单例
	{
		//生成二维码
		apiPlatform.Get("/qrcode/{text:string}", frontEndControllers.HtmlC.QrcodeText())
		//第三方wap支付post提交中转
		apiPlatform.Get("/tpwap", func(ctx iris.Context) { pay.NewPayController(ctx).TPWap() })
	}

	//非登录路由
	apiDown := apiPlatform.Party("/app")
	{
		//进入下载页
		apiDown.Get("/download", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).ToDownload() })
	}

	//非登录路由
	apiv1 := apiPlatform.Party("/api/v1")
	{
		//app版本更新检测
		apiv1.Get("/getAppVerStatus", func(ctx iris.Context) { frontEndControllers.NewApiConfigController(ctx).GetAppVerUpdate() })
		//启动配置信息获取
		apiv1.Get("/getConfig", func(ctx iris.Context) { frontEndControllers.NewApiConfigController(ctx).GetConfig() })
		//获取客服接口
		apiv1.Get("/getCustomerService", func(ctx iris.Context) { frontEndControllers.NewApiConfigController(ctx).GetCustomerService() })
		//vip详情
		apiv1.Get("/vipDetail", func(ctx iris.Context) { frontEndControllers.NewShowController(ctx).GetVipDetail() })
		//代理详情介绍
		apiv1.Get("/proxyDetail", func(ctx iris.Context) { frontEndControllers.NewShowController(ctx).GetProxyDetail() })
		//注册
		apiv1.Post("/register", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).Register() })
		apiv1.Post("/registerUserName", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).RegisterUserName() })
		//获取验证码
		apiv1.Get("/getVcode", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).GetVcode() })
		apiv1.Get("/getCode", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).GetCode() })
		//登录
		apiv1.Post("/login", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).Login() })
		//游客登录
		apiv1.Post("/visitorLogin", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).VisitorLogin() })
		//获取优惠活动
		apiv1.Get("/getActivity", func(ctx iris.Context) { frontEndControllers.NewShowController(ctx).GetActivity() })
		//获取优惠活动分类
		apiv1.Get("/getActivityClass", func(ctx iris.Context) { frontEndControllers.NewShowController(ctx).GetActivityClass() })
		//获取系统公告
		apiv1.Get("/getSystemNotice", func(ctx iris.Context) { frontEndControllers.NewShowController(ctx).GetSystemNotice() })
		//获取首页游戏公告
		apiv1.Get("/getHomeFirstSysNotice", func(ctx iris.Context) { frontEndControllers.NewShowController(ctx).GetHomeFirstSysNotice() })
		//获取游戏列表
		apiv1.Get("/getGameList", func(ctx iris.Context) { frontEndControllers.NewGameController(ctx).GetGameList() })
		//获取游戏分类列表前端缓存
		apiv1.Get("/getGameListNewCache", func(ctx iris.Context) { frontEndControllers.NewGameController(ctx).GetGameListNewCache() })
		//获取游戏子分类列表
		apiv1.Get("/getGameListSub", func(ctx iris.Context) { frontEndControllers.NewGameController(ctx).GetGameListSub() })

		apiv1.Get("/updateGameList", func(ctx iris.Context) { frontEndControllers.NewGameController(ctx).UpdateGameList() })

		//支付同步回调
		/*
			payNameCode为pay_credential表的plat_form字段
		*/
		apiv1.Get("/tpPayClientNotify/{payNameCode:int}/{credentialId:int}/{userId:int}", func(ctx iris.Context) {
			pay.NewPayController(ctx).TpPayClientNotify()
		})
		apiv1.Post("/tpPayClientNotify/{payNameCode:int}/{credentialId:int}/{userId:int}", func(ctx iris.Context) {
			pay.NewPayController(ctx).TpPayClientNotify()
		})
		//支付异步回调
		apiv1.Get("/tppayCallBack/{payNameCode:int}/{credentialId:int}/{userId:int}", func(ctx iris.Context) {
			pay.NewPayController(ctx).TppayCallBack()
		})
		apiv1.Post("/tppayCallBack/{payNameCode:int}/{credentialId:int}/{userId:int}", func(ctx iris.Context) {
			pay.NewPayController(ctx).TppayCallBack()
		})
		//银行卡列表
		apiv1.Get("/bankCardsList", func(ctx iris.Context) { pay.NewPayController(ctx).BankCardsList() })
		apiv1.Post("/wxLogin", func(ctx iris.Context) { frontEndControllers.NewWxController(ctx).WxLogin() })
	}

	//需要登录之后才能访问的路由,//jwt中间件
	apiAuthV1 := apiPlatform.Party("/api/auth/v1", middlewares.JwtAuthenticate().Serve, middlewares.JwtHandler())
	{
		//修改密码
		apiAuthV1.Post("/modifyPwd", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).ModifyPwd() })
		//退出登录
		apiAuthV1.Get("/logout", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).Logout() })
		//获取站内信
		apiAuthV1.Get("/getNotice", func(ctx iris.Context) { frontEndControllers.NewShowController(ctx).GetNotice() })
		//获取活动奖励
		apiAuthV1.Post("/getActivityAward", func(ctx iris.Context) { frontEndControllers.NewShowController(ctx).GetActivityAward() })
		//设置保险箱密码
		apiAuthV1.Post("/setSafeBoxPwd", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).SetSafeBoxPwd() })
		//进入保险箱
		apiAuthV1.Post("/intoSafeBox", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).IntoSafeBox() })
		//保险箱存取款
		apiAuthV1.Post("/safeBoxOperation", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).SafeBoxOperation() })
		//获取账号信息，刷新余额等操作
		apiAuthV1.Get("/getAccount", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).GetAccount() })
		//获取保险箱明细
		apiAuthV1.Get("/safeBoxInfo", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).SafeBoxInfo() })
		//获取用户个人信息
		apiAuthV1.Get("/getUser", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).GetUser() })
		//获取用户个人报表
		apiAuthV1.Get("/getUserReport", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).GetUserReport() })
		//修改个人信息
		apiAuthV1.Post("/updateUser", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).UpdateUser() })
		//绑定手机号
		apiAuthV1.Post("/bindPhone", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).BindPhone() })
		//启动游戏
		apiAuthV1.Get("/launchGame", func(ctx iris.Context) { frontEndControllers.NewGameController(ctx).LaunchGame() })
		//退出游戏
		apiAuthV1.Get("/exitGame", func(ctx iris.Context) { frontEndControllers.NewGameController(ctx).ExitGame() })
		//推广赚钱中的团队成员接口
		apiAuthV1.Get("/teamMembers", func(ctx iris.Context) { frontEndControllers.NewProxyController(ctx).TeamMembers() })
		//推广赚钱主要接口
		apiAuthV1.Get("/promotion", func(ctx iris.Context) { frontEndControllers.NewProxyController(ctx).Promotion() })
		//领取佣金
		apiAuthV1.Get("/receiveCommission", func(ctx iris.Context) { frontEndControllers.NewProxyController(ctx).ReceiveCommission() })
		//佣金记录详情
		apiAuthV1.Get("/proxyCommissionsInfo", func(ctx iris.Context) { frontEndControllers.NewProxyController(ctx).ProxyCommissionsInfo() })
		apiAuthV1.Get("/commissionRecord", func(ctx iris.Context) { frontEndControllers.NewProxyController(ctx).CommissionRecord() })
		//充值类型列表
		apiAuthV1.Get("/chargeTypes", func(ctx iris.Context) { frontEndControllers.NewFinanceController(ctx).ChargeTypesList() })
		//获取投注记录
		apiAuthV1.Get("/getBetsInfo", func(ctx iris.Context) { frontEndControllers.NewBetController(ctx).GetBetsInfo() })
		//获取洗码信息
		apiAuthV1.Get("/getWashCode", func(ctx iris.Context) { frontEndControllers.NewBetController(ctx).GetWashCode() })
		//手动洗码
		apiAuthV1.Get("/washCode", func(ctx iris.Context) { frontEndControllers.NewBetController(ctx).WashCode() })
		//获取洗码历史记录
		apiAuthV1.Get("/getWashCodeRecords", func(ctx iris.Context) { frontEndControllers.NewBetController(ctx).GetWashCodeRecords() })
		//获取洗码历史记录详情
		apiAuthV1.Get("/getWashCodeInfos", func(ctx iris.Context) { frontEndControllers.NewBetController(ctx).GetWashCodeInfos() })
		//获取资金明细记录
		apiAuthV1.Get("/getAccountInfo", func(ctx iris.Context) { frontEndControllers.NewBetController(ctx).GetAccountInfo() })

		//获取投注记录平台查询列表
		apiAuthV1.Get("/getBetsSearchType", func(ctx iris.Context) { frontEndControllers.NewBetController(ctx).GetBetsSearchType() })
		//请求支付生成订单接口
		apiAuthV1.Post("/payAction", func(ctx iris.Context) { pay.NewPayController(ctx).PayAction() })
		//公司入款支付(非第三方支付)
		apiAuthV1.Post("/payCompanyAction", func(ctx iris.Context) { pay.NewPayController(ctx).PayCompanyAction() })
		//公司入款二维码扫描
		apiAuthV1.Post("/payCompanyActionQrCan", func(ctx iris.Context) { pay.NewPayController(ctx).PayCompanyActionQrCan() })
		//获取充值记录
		apiAuthV1.Get("/chargeRecordList", func(ctx iris.Context) { pay.NewPayController(ctx).ChargeRecordList() })
		//绑定银行卡接口
		apiAuthV1.Post("/bindingBankCard", func(ctx iris.Context) { pay.NewPayController(ctx).BindingBankCard() })
		//个人银行卡列表获取接口
		apiAuthV1.Get("/userBankCards", func(ctx iris.Context) { pay.NewPayController(ctx).UserBankCards() })
		//提现记录接口
		apiAuthV1.Get("/withDrawRecordsList", func(ctx iris.Context) { pay.NewPayController(ctx).WithDrawRecordsList() })
		//提现接口
		apiAuthV1.Post("/withDrawAction", func(ctx iris.Context) { pay.NewPayController(ctx).WithDrawAction() })
		//提现页面流水详情
		apiAuthV1.Get("/runningWaterDetail", func(ctx iris.Context) { pay.NewPayController(ctx).RunningWaterDetail() })
		//提现页面资金明细
		apiAuthV1.Get("/withdrawFundParticulars", func(ctx iris.Context) { pay.NewPayController(ctx).WithdrawFundParticulars() })
		//红包查询接口
		apiAuthV1.Get("/redPacketQuery", func(ctx iris.Context) { frontEndControllers.NewFinanceController(ctx).RedPacketQuery() })
		//红包领取接口
		apiAuthV1.Post("/redPacketReceiveAction", func(ctx iris.Context) { frontEndControllers.NewFinanceController(ctx).RedPacketReceiveAction() })
		//会员每日签到
		apiAuthV1.Post("/checkIn", func(ctx iris.Context) { frontEndControllers.NewUserController(ctx).CheckIn() })
	}

}
