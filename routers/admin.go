package routers

import (
	"qpgame/admin/controllers"
	"qpgame/admin/websockets"
	"qpgame/common/utils"
	"qpgame/middlewares"

	"github.com/kataras/iris"
)

// 关于cors的跨域问题的返回
func responseCorsOptions(ctx *iris.Context) {
	utils.ResponseCors(ctx)
	(*ctx).StatusCode(200)
	(*ctx).WriteString("")
}

// 路由配置, 路由主要分三种情况:
// 1. 不需要平台识别号的, 同时不需要权限验证的
// 2. 需要平台识别号, 但是不需要权限验证的
// 3. 需要平台识别号, 并且需要权限验证的
func AdminApiRouters(app *iris.Application) {

	//关于websocket的处理
	app.Get("/ws", websockets.AdminWebSocket.Handler())

	// 1. 不需要平台识别号
	noPlatform := app.Party(
		"/admin/api/v1",
		middlewares.CorsMiddleware(), //跨域中间件
	).AllowMethods(iris.MethodOptions)
	{
		noPlatform.Get("/index/index", func(ctx iris.Context) { (&controllers.IndexController{}).Index(&ctx) })
		noPlatform.Get("/index/config", func(ctx iris.Context) { (&controllers.IndexController{}).Config(&ctx) })
	}

	api := app.Party(
		"/{platform:string}/admin/api",
		middlewares.CorsMiddleware(), //跨域中间件
	).AllowMethods(iris.MethodOptions)
	{
		// 2.1 需要平台识别号, 非登录路由
		apiNotAuthV1 := api.Party("/v1")
		{
			apiNotAuthV1.Post("/index/login_reset", func(ctx iris.Context) { (&controllers.IndexController{}).LoginReset(&ctx) })
			apiNotAuthV1.Post("/index/login", func(ctx iris.Context) { (&controllers.IndexController{}).Login(&ctx) })
			apiNotAuthV1.Get("/index/verify", func(ctx iris.Context) { (&controllers.IndexController{}).Verify(&ctx) })
		}

		// 2.2 需要平台识别号, 需要登录的路由
		apiAuthV1 := api.Party("/auth/v1", middlewares.AdminJwtHandler())
		{
			apiAuthV1.Get("/index/logout", func(ctx iris.Context) { (&controllers.IndexController{}).Logout(&ctx) })
			apiAuthV1.Get("/account_infos", func(ctx iris.Context) { (&controllers.AccountInfosController{}).Index(&ctx) })
			apiAuthV1.Get("/account_infos/view", func(ctx iris.Context) { (&controllers.AccountInfosController{}).View(&ctx) })
			apiAuthV1.Get("/account_statistics", func(ctx iris.Context) { (&controllers.AccountStatisticsController{}).Index(&ctx) })
			//apiAuthV1.Get("/account_statistics/view", func(ctx iris.Context) { (&controllers.AccountStatisticsController{}).View(&ctx) })
			//apiAuthV1.Get("/account_statistics/delete", func(ctx iris.Context) { (&controllers.AccountStatisticsController{}).Delete(&ctx) })
			apiAuthV1.Get("/accounts", func(ctx iris.Context) { (&controllers.AccountsController{}).Index(&ctx) })
			apiAuthV1.Get("/accounts/view", func(ctx iris.Context) { (&controllers.AccountsController{}).View(&ctx) })
			apiAuthV1.Get("/accounts/delete", func(ctx iris.Context) { (&controllers.AccountsController{}).Delete(&ctx) })

			apiAuthV1.Get("/activities", func(ctx iris.Context) { (&controllers.ActivitiesController{}).Index(&ctx) })
			apiAuthV1.Post("/activities/add", func(ctx iris.Context) { (&controllers.ActivitiesController{}).Save(&ctx) })
			apiAuthV1.Post("/activities/update", func(ctx iris.Context) { (&controllers.ActivitiesController{}).Save(&ctx) })
			apiAuthV1.Get("/activities/view", func(ctx iris.Context) { (&controllers.ActivitiesController{}).View(&ctx) })
			apiAuthV1.Get("/activities/delete", func(ctx iris.Context) { (&controllers.ActivitiesController{}).Delete(&ctx) })

			apiAuthV1.Get("/activity_classes", func(ctx iris.Context) { (&controllers.ActivityClassesController{}).Index(&ctx) })
			apiAuthV1.Post("/activity_classes/add", func(ctx iris.Context) { (&controllers.ActivityClassesController{}).Save(&ctx) })
			apiAuthV1.Post("/activity_classes/update", func(ctx iris.Context) { (&controllers.ActivityClassesController{}).Save(&ctx) })
			apiAuthV1.Get("/activity_classes/view", func(ctx iris.Context) { (&controllers.ActivityClassesController{}).View(&ctx) })
			apiAuthV1.Get("/activity_classes/delete", func(ctx iris.Context) { (&controllers.ActivityClassesController{}).Delete(&ctx) })

			apiAuthV1.Get("/activity_records", func(ctx iris.Context) { (&controllers.ActivityRecordsController{}).Index(&ctx) })
			apiAuthV1.Get("/activity_records/view", func(ctx iris.Context) { (&controllers.ActivityRecordsController{}).View(&ctx) })
			apiAuthV1.Get("/activity_records/delete", func(ctx iris.Context) { (&controllers.ActivityRecordsController{}).Delete(&ctx) })

			apiAuthV1.Get("/admin_ips", func(ctx iris.Context) { (&controllers.AdminIpsController{}).Index(&ctx) })
			apiAuthV1.Get("/admin_ips/delete", func(ctx iris.Context) { (&controllers.AdminIpsController{}).Delete(&ctx) })

			apiAuthV1.Get("/admin_login_logs", func(ctx iris.Context) { (&controllers.AdminLoginLogsController{}).Index(&ctx) })
			apiAuthV1.Get("/admin_login_logs/view", func(ctx iris.Context) { (&controllers.AdminLoginLogsController{}).View(&ctx) })
			apiAuthV1.Get("/admin_login_logs/delete", func(ctx iris.Context) { (&controllers.AdminLoginLogsController{}).Delete(&ctx) })

			apiAuthV1.Get("/admin_logs", func(ctx iris.Context) { (&controllers.AdminLogsController{}).Index(&ctx) })
			apiAuthV1.Get("/admin_logs/view", func(ctx iris.Context) { (&controllers.AdminLogsController{}).View(&ctx) })
			apiAuthV1.Get("/admin_logs/delete", func(ctx iris.Context) { (&controllers.AdminLogsController{}).Delete(&ctx) })

			apiAuthV1.Get("/admin_nodes", func(ctx iris.Context) { (&controllers.AdminNodesController{}).Index(&ctx) })
			apiAuthV1.Post("/admin_nodes/add", func(ctx iris.Context) { (&controllers.AdminNodesController{}).Save(&ctx) })
			apiAuthV1.Post("/admin_nodes/update", func(ctx iris.Context) { (&controllers.AdminNodesController{}).Save(&ctx) })
			apiAuthV1.Post("/admin_nodes/sort", func(ctx iris.Context) { (&controllers.AdminNodesController{}).Sort(&ctx) })
			apiAuthV1.Get("/admin_nodes/view", func(ctx iris.Context) { (&controllers.AdminNodesController{}).View(&ctx) })
			apiAuthV1.Get("/admin_nodes/delete", func(ctx iris.Context) { (&controllers.AdminNodesController{}).Delete(&ctx) })

			apiAuthV1.Get("/admin_roles", func(ctx iris.Context) { (&controllers.AdminRolesController{}).Index(&ctx) })
			apiAuthV1.Post("/admin_roles/add", func(ctx iris.Context) { (&controllers.AdminRolesController{}).Save(&ctx) })
			apiAuthV1.Post("/admin_roles/update", func(ctx iris.Context) { (&controllers.AdminRolesController{}).Save(&ctx) })
			apiAuthV1.Get("/admin_roles/view", func(ctx iris.Context) { (&controllers.AdminRolesController{}).View(&ctx) })
			apiAuthV1.Get("/admin_roles/delete", func(ctx iris.Context) { (&controllers.AdminRolesController{}).Delete(&ctx) })
			apiAuthV1.Post("/admin_roles/menus", func(ctx iris.Context) { (&controllers.AdminRolesController{}).SetMenus(&ctx) })

			apiAuthV1.Get("/admins", func(ctx iris.Context) { (&controllers.AdminsController{}).Index(&ctx) })
			apiAuthV1.Post("/admins/add", func(ctx iris.Context) { (&controllers.AdminsController{}).Save(&ctx) })
			apiAuthV1.Post("/admins/update", func(ctx iris.Context) { (&controllers.AdminsController{}).Save(&ctx) })
			apiAuthV1.Get("/admins/view", func(ctx iris.Context) { (&controllers.AdminsController{}).View(&ctx) })
			apiAuthV1.Post("/admins/update_password", func(ctx iris.Context) { (&controllers.AdminsController{}).UpdatePassword(&ctx) })

			apiAuthV1.Get("/app_versions", func(ctx iris.Context) { (&controllers.AppVersionsController{}).Index(&ctx) })
			apiAuthV1.Post("/app_versions/add", func(ctx iris.Context) { (&controllers.AppVersionsController{}).Save(&ctx) })
			apiAuthV1.Post("/app_versions/update", func(ctx iris.Context) { (&controllers.AppVersionsController{}).Save(&ctx) })
			apiAuthV1.Get("/app_versions/view", func(ctx iris.Context) { (&controllers.AppVersionsController{}).View(&ctx) })
			apiAuthV1.Get("/app_versions/delete", func(ctx iris.Context) { (&controllers.AppVersionsController{}).Delete(&ctx) })

			apiAuthV1.Get("/bets", func(ctx iris.Context) { (&controllers.BetsController{}).Index(&ctx) })
			apiAuthV1.Get("/bets/get_calculated_total", func(ctx iris.Context) { (&controllers.BetsController{}).GetCalculatedTotal(&ctx) })
			apiAuthV1.Get("/bets/view", func(ctx iris.Context) { (&controllers.BetsController{}).View(&ctx) })
			apiAuthV1.Get("/bets/delete", func(ctx iris.Context) { (&controllers.BetsController{}).Delete(&ctx) })

			apiAuthV1.Get("/blacklists", func(ctx iris.Context) { (&controllers.BlacklistsController{}).Index(&ctx) })
			apiAuthV1.Post("/blacklists/add", func(ctx iris.Context) { (&controllers.BlacklistsController{}).Save(&ctx) })
			apiAuthV1.Post("/blacklists/update", func(ctx iris.Context) { (&controllers.BlacklistsController{}).Save(&ctx) })
			apiAuthV1.Get("/blacklists/view", func(ctx iris.Context) { (&controllers.BlacklistsController{}).View(&ctx) })
			apiAuthV1.Get("/blacklists/delete", func(ctx iris.Context) { (&controllers.BlacklistsController{}).Delete(&ctx) })

			apiAuthV1.Get("/charge_cards", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).Index(&ctx) })
			apiAuthV1.Get("/charge_cards/get_user_groups", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).GetUserGroups(&ctx) })
			apiAuthV1.Get("/charge_cards/get_pay_credentials", func(ctx iris.Context) { (&controllers.PayCredentialsController{}).All(&ctx) })
			apiAuthV1.Get("/charge_cards/onlines", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).Onlines(&ctx) })
			apiAuthV1.Get("/charge_cards/delete_online", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).OnlineDelete(&ctx) })
			apiAuthV1.Get("/charge_cards/view_online", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).OnlineView(&ctx) })
			apiAuthV1.Post("/charge_cards/add_online", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).Save(&ctx) })
			apiAuthV1.Post("/charge_cards/add", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).Save(&ctx) })
			apiAuthV1.Post("/charge_cards/update", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).Save(&ctx) })
			apiAuthV1.Get("/charge_cards/view", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).View(&ctx) })
			apiAuthV1.Get("/charge_cards/delete", func(ctx iris.Context) { (&controllers.ChargeCardsController{}).Delete(&ctx) })

			apiAuthV1.Get("/charge_records", func(ctx iris.Context) { (&controllers.ChargeRecordsController{}).Index(&ctx) })
			apiAuthV1.Get("/charge_records/onlines", func(ctx iris.Context) { (&controllers.ChargeRecordsController{}).Onlines(&ctx) })
			apiAuthV1.Get("/charge_records/forced_deposit", func(ctx iris.Context) { (&controllers.ChargeRecordsController{}).ForcedDeposit(&ctx) })
			apiAuthV1.Get("/charge_records/view", func(ctx iris.Context) { (&controllers.ChargeRecordsController{}).View(&ctx) })
			apiAuthV1.Get("/charge_records/delete", func(ctx iris.Context) { (&controllers.ChargeRecordsController{}).Delete(&ctx) })
			apiAuthV1.Get("/charge_records/allow", func(ctx iris.Context) { (&controllers.ChargeRecordsController{}).Allow(&ctx) })
			apiAuthV1.Get("/charge_records/deny", func(ctx iris.Context) { (&controllers.ChargeRecordsController{}).Deny(&ctx) })

			apiAuthV1.Get("/charge_types", func(ctx iris.Context) { (&controllers.ChargeTypesController{}).Index(&ctx) })
			apiAuthV1.Post("/charge_types/add", func(ctx iris.Context) { (&controllers.ChargeTypesController{}).Save(&ctx) })
			apiAuthV1.Post("/charge_types/update", func(ctx iris.Context) { (&controllers.ChargeTypesController{}).Save(&ctx) })
			apiAuthV1.Get("/charge_types/view", func(ctx iris.Context) { (&controllers.ChargeTypesController{}).View(&ctx) })
			apiAuthV1.Get("/charge_types/delete", func(ctx iris.Context) { (&controllers.ChargeTypesController{}).Delete(&ctx) })

			apiAuthV1.Get("/configs", func(ctx iris.Context) { (&controllers.ConfigsController{}).Index(&ctx) })
			//apiAuthV1.Post("/configs/add", func(ctx iris.Context) { (&controllers.ConfigsController{}).Save(&ctx) })
			//apiAuthV1.Post("/configs/update", func(ctx iris.Context) { (&controllers.ConfigsController{}).Save(&ctx) })
			//apiAuthV1.Get("/configs/view", func(ctx iris.Context) { (&controllers.ConfigsController{}).View(&ctx) })
			//apiAuthV1.Get("/configs/delete", func(ctx iris.Context) { (&controllers.ConfigsController{}).Delete(&ctx) })
			apiAuthV1.Get("/configs/sets", func(ctx iris.Context) { (&controllers.ConfigsController{}).Sets(&ctx) })
			apiAuthV1.Post("/configs/set_sets", func(ctx iris.Context) { (&controllers.ConfigsController{}).SetSets(&ctx) })
			apiAuthV1.Get("/configs/service", func(ctx iris.Context) { (&controllers.ConfigsController{}).Service(&ctx) })
			apiAuthV1.Post("/configs/set_service", func(ctx iris.Context) { (&controllers.ConfigsController{}).SetService(&ctx) })
			apiAuthV1.Get("/configs/faq", func(ctx iris.Context) { (&controllers.ConfigsController{}).Faq(&ctx) })
			apiAuthV1.Post("/configs/set_faq", func(ctx iris.Context) { (&controllers.ConfigsController{}).SetFaq(&ctx) })
			apiAuthV1.Get("/configs/charges", func(ctx iris.Context) { (&controllers.ConfigsController{}).Charges(&ctx) })
			apiAuthV1.Post("/configs/set_charges", func(ctx iris.Context) { (&controllers.ConfigsController{}).SetCharges(&ctx) })
			apiAuthV1.Get("/configs/reg", func(ctx iris.Context) { (&controllers.ConfigsController{}).Reg(&ctx) })
			apiAuthV1.Post("/configs/set_reg", func(ctx iris.Context) { (&controllers.ConfigsController{}).SetReg(&ctx) })
			apiAuthV1.Get("/configs/order_alert", func(ctx iris.Context) { (&controllers.ConfigsController{}).OrderAlert(&ctx) })
			apiAuthV1.Get("/configs/fund_dama_rate", func(ctx iris.Context) { (&controllers.ConfigsController{}).FundDamaRate(&ctx) })
			apiAuthV1.Post("/configs/set_fund_dama_rate", func(ctx iris.Context) { (&controllers.ConfigsController{}).SetFundDamaRate(&ctx) })
			apiAuthV1.Get("/configs/com_bank_present_rate", func(ctx iris.Context) { (&controllers.ConfigsController{}).BankPresentRate(&ctx) })
			apiAuthV1.Post("/configs/set_com_bank_present_rate", func(ctx iris.Context) { (&controllers.ConfigsController{}).SetBankPresentRate(&ctx) })
			apiAuthV1.Get("/configs/silver_merchant", func(ctx iris.Context) { (&controllers.ConfigsController{}).SilverMerchant(&ctx) })
			apiAuthV1.Post("/configs/set_silver_merchant", func(ctx iris.Context) { (&controllers.ConfigsController{}).SetSilverMerchant(&ctx) })

			apiAuthV1.Get("/conversion_records", func(ctx iris.Context) { (&controllers.ConversionRecordsController{}).Index(&ctx) })
			apiAuthV1.Get("/conversion_records/view", func(ctx iris.Context) { (&controllers.ConversionRecordsController{}).View(&ctx) })
			apiAuthV1.Get("/conversion_records/delete", func(ctx iris.Context) { (&controllers.ConversionRecordsController{}).Delete(&ctx) })

			apiAuthV1.Get("/game_categories", func(ctx iris.Context) { (&controllers.GameCategoriesController{}).Index(&ctx) })
			apiAuthV1.Post("/game_categories/add", func(ctx iris.Context) { (&controllers.GameCategoriesController{}).Save(&ctx) })
			apiAuthV1.Post("/game_categories/update", func(ctx iris.Context) { (&controllers.GameCategoriesController{}).Save(&ctx) })
			apiAuthV1.Get("/game_categories/view", func(ctx iris.Context) { (&controllers.GameCategoriesController{}).View(&ctx) })
			apiAuthV1.Get("/game_categories/delete", func(ctx iris.Context) { (&controllers.GameCategoriesController{}).Delete(&ctx) })
			apiAuthV1.Get("/game_categories/relations", func(ctx iris.Context) { (&controllers.GameCategoriesController{}).Relations(&ctx) })

			apiAuthV1.Get("/notices", func(ctx iris.Context) { (&controllers.NoticesController{}).Index(&ctx) })
			apiAuthV1.Post("/notices/add", func(ctx iris.Context) { (&controllers.NoticesController{}).Save(&ctx) })
			apiAuthV1.Post("/notices/update", func(ctx iris.Context) { (&controllers.NoticesController{}).Save(&ctx) })
			apiAuthV1.Get("/notices/view", func(ctx iris.Context) { (&controllers.NoticesController{}).View(&ctx) })
			apiAuthV1.Get("/notices/delete", func(ctx iris.Context) { (&controllers.NoticesController{}).Delete(&ctx) })

			apiAuthV1.Get("/phone_codes", func(ctx iris.Context) { (&controllers.PhoneCodesController{}).Index(&ctx) })
			apiAuthV1.Get("/phone_codes/view", func(ctx iris.Context) { (&controllers.PhoneCodesController{}).View(&ctx) })
			apiAuthV1.Get("/phone_codes/delete", func(ctx iris.Context) { (&controllers.PhoneCodesController{}).Delete(&ctx) })

			apiAuthV1.Get("/platform_accounts", func(ctx iris.Context) { (&controllers.PlatformAccountsController{}).Index(&ctx) })
			apiAuthV1.Post("/platform_accounts/add", func(ctx iris.Context) { (&controllers.PlatformAccountsController{}).Save(&ctx) })
			apiAuthV1.Post("/platform_accounts/update", func(ctx iris.Context) { (&controllers.PlatformAccountsController{}).Save(&ctx) })
			apiAuthV1.Get("/platform_accounts/view", func(ctx iris.Context) { (&controllers.PlatformAccountsController{}).View(&ctx) })
			apiAuthV1.Get("/platform_accounts/delete", func(ctx iris.Context) { (&controllers.PlatformAccountsController{}).Delete(&ctx) })

			apiAuthV1.Get("/platform_games", func(ctx iris.Context) { (&controllers.PlatformGamesController{}).Index(&ctx) })
			apiAuthV1.Post("/platform_games/add", func(ctx iris.Context) { (&controllers.PlatformGamesController{}).Save(&ctx) })
			apiAuthV1.Post("/platform_games/update", func(ctx iris.Context) { (&controllers.PlatformGamesController{}).Save(&ctx) })
			apiAuthV1.Get("/platform_games/view", func(ctx iris.Context) { (&controllers.PlatformGamesController{}).View(&ctx) })
			apiAuthV1.Get("/platform_games/delete", func(ctx iris.Context) { (&controllers.PlatformGamesController{}).Delete(&ctx) })

			apiAuthV1.Get("/platforms", func(ctx iris.Context) { (&controllers.PlatformsController{}).Index(&ctx) })
			apiAuthV1.Post("/platforms/add", func(ctx iris.Context) { (&controllers.PlatformsController{}).Save(&ctx) })
			apiAuthV1.Post("/platforms/update", func(ctx iris.Context) { (&controllers.PlatformsController{}).Save(&ctx) })
			apiAuthV1.Get("/platforms/view", func(ctx iris.Context) { (&controllers.PlatformsController{}).View(&ctx) })
			apiAuthV1.Get("/platforms/delete", func(ctx iris.Context) { (&controllers.PlatformsController{}).Delete(&ctx) })

			apiAuthV1.Get("/proxy_chess_levels", func(ctx iris.Context) { (&controllers.ProxyChessLevelsController{}).Index(&ctx) })
			//apiAuthV1.Post("/proxy_chess_levels/add", func(ctx iris.Context) { (&controllers.ProxyChessLevelsController{}).Save(&ctx) })
			apiAuthV1.Post("/proxy_chess_levels/update", func(ctx iris.Context) { (&controllers.ProxyChessLevelsController{}).Save(&ctx) })
			apiAuthV1.Get("/proxy_chess_levels/view", func(ctx iris.Context) { (&controllers.ProxyChessLevelsController{}).View(&ctx) })
			//apiAuthV1.Get("/proxy_chess_levels/delete", func(ctx iris.Context) { (&controllers.ProxyChessLevelsController{}).Delete(&ctx) })

			apiAuthV1.Get("/proxy_real_levels", func(ctx iris.Context) { (&controllers.ProxyRealLevelsController{}).Index(&ctx) })
			//apiAuthV1.Post("/proxy_real_levels/add", func(ctx iris.Context) { (&controllers.ProxyRealLevelsController{}).Save(&ctx) })
			apiAuthV1.Post("/proxy_real_levels/update", func(ctx iris.Context) { (&controllers.ProxyRealLevelsController{}).Save(&ctx) })
			apiAuthV1.Get("/proxy_real_levels/view", func(ctx iris.Context) { (&controllers.ProxyRealLevelsController{}).View(&ctx) })
			//apiAuthV1.Get("/proxy_real_levels/delete", func(ctx iris.Context) { (&controllers.ProxyRealLevelsController{}).Delete(&ctx) })

			apiAuthV1.Get("/system_notices", func(ctx iris.Context) { (&controllers.SystemNoticesController{}).Index(&ctx) })
			apiAuthV1.Post("/system_notices/add", func(ctx iris.Context) { (&controllers.SystemNoticesController{}).Save(&ctx) })
			apiAuthV1.Post("/system_notices/update", func(ctx iris.Context) { (&controllers.SystemNoticesController{}).Save(&ctx) })
			apiAuthV1.Get("/system_notices/view", func(ctx iris.Context) { (&controllers.SystemNoticesController{}).View(&ctx) })

			apiAuthV1.Get("/system_notices/delete", func(ctx iris.Context) { (&controllers.SystemNoticesController{}).Delete(&ctx) })
			apiAuthV1.Get("/user_bank_cards", func(ctx iris.Context) { (&controllers.UserBankCardsController{}).Index(&ctx) })
			apiAuthV1.Post("/user_bank_cards/add", func(ctx iris.Context) { (&controllers.UserBankCardsController{}).Save(&ctx) })
			apiAuthV1.Post("/user_bank_cards/update", func(ctx iris.Context) { (&controllers.UserBankCardsController{}).Save(&ctx) })
			apiAuthV1.Get("/user_bank_cards/view", func(ctx iris.Context) { (&controllers.UserBankCardsController{}).View(&ctx) })
			apiAuthV1.Get("/user_bank_cards/delete", func(ctx iris.Context) { (&controllers.UserBankCardsController{}).Delete(&ctx) })

			apiAuthV1.Get("/user_banks", func(ctx iris.Context) { (&controllers.UserBanksController{}).Index(&ctx) })
			apiAuthV1.Post("/user_banks/add", func(ctx iris.Context) { (&controllers.UserBanksController{}).Save(&ctx) })
			apiAuthV1.Post("/user_banks/update", func(ctx iris.Context) { (&controllers.UserBanksController{}).Save(&ctx) })
			apiAuthV1.Get("/user_banks/view", func(ctx iris.Context) { (&controllers.UserBanksController{}).View(&ctx) })
			apiAuthV1.Get("/user_banks/delete", func(ctx iris.Context) { (&controllers.UserBanksController{}).Delete(&ctx) })

			apiAuthV1.Get("/user_login_logs", func(ctx iris.Context) { (&controllers.UserLoginLogsController{}).Index(&ctx) })
			apiAuthV1.Get("/user_login_logs/view", func(ctx iris.Context) { (&controllers.UserLoginLogsController{}).View(&ctx) })
			apiAuthV1.Get("/user_login_logs/delete", func(ctx iris.Context) { (&controllers.UserLoginLogsController{}).Delete(&ctx) })

			apiAuthV1.Get("/users", func(ctx iris.Context) { (&controllers.UsersController{}).Index(&ctx) })
			apiAuthV1.Post("/users/add", func(ctx iris.Context) { (&controllers.UsersController{}).Save(&ctx) })
			apiAuthV1.Post("/users/update", func(ctx iris.Context) { (&controllers.UsersController{}).Save(&ctx) })
			apiAuthV1.Get("/users/view", func(ctx iris.Context) { (&controllers.UsersController{}).View(&ctx) })
			apiAuthV1.Get("/users/delete", func(ctx iris.Context) { (&controllers.UsersController{}).Delete(&ctx) })
			apiAuthV1.Post("/users/update_password", func(ctx iris.Context) { (&controllers.UsersController{}).UpdatePassword(&ctx) })
			apiAuthV1.Post("/users/update_safe_password", func(ctx iris.Context) { (&controllers.UsersController{}).UpdateSafePassword(&ctx) })
			apiAuthV1.Get("/users/lock", func(ctx iris.Context) { (&controllers.UsersController{}).Lock(&ctx) })
			apiAuthV1.Get("/users/unlock", func(ctx iris.Context) { (&controllers.UsersController{}).Unlock(&ctx) })
			apiAuthV1.Get("/users/query", func(ctx iris.Context) { (&controllers.UsersController{}).QueryUser(&ctx) })
			apiAuthV1.Get("/users/invite", func(ctx iris.Context) { (&controllers.UsersController{}).Invite(&ctx) })

			apiAuthV1.Get("/vip_levels", func(ctx iris.Context) { (&controllers.VipLevelsController{}).Index(&ctx) })
			//apiAuthV1.Post("/vip_levels/add", func(ctx iris.Context) { (&controllers.VipLevelsController{}).Save(&ctx) })
			apiAuthV1.Post("/vip_levels/update", func(ctx iris.Context) { (&controllers.VipLevelsController{}).Save(&ctx) })
			apiAuthV1.Get("/vip_levels/view", func(ctx iris.Context) { (&controllers.VipLevelsController{}).View(&ctx) })
			//apiAuthV1.Get("/vip_levels/delete", func(ctx iris.Context) { (&controllers.VipLevelsController{}).Delete(&ctx) })

			apiAuthV1.Get("/user_groups", func(ctx iris.Context) { (&controllers.UserGroupsController{}).Index(&ctx) })
			apiAuthV1.Post("/user_groups/add", func(ctx iris.Context) { (&controllers.UserGroupsController{}).Save(&ctx) })
			apiAuthV1.Post("/user_groups/update", func(ctx iris.Context) { (&controllers.UserGroupsController{}).Save(&ctx) })
			apiAuthV1.Get("/user_groups/view", func(ctx iris.Context) { (&controllers.UserGroupsController{}).View(&ctx) })
			apiAuthV1.Get("/user_groups/delete", func(ctx iris.Context) { (&controllers.UserGroupsController{}).Delete(&ctx) })

			apiAuthV1.Get("/pay_credentials", func(ctx iris.Context) { (&controllers.PayCredentialsController{}).Index(&ctx) })
			apiAuthV1.Post("/pay_credentials/add", func(ctx iris.Context) { (&controllers.PayCredentialsController{}).Save(&ctx) })
			apiAuthV1.Post("/pay_credentials/update", func(ctx iris.Context) { (&controllers.PayCredentialsController{}).Save(&ctx) })
			apiAuthV1.Get("/pay_credentials/view", func(ctx iris.Context) { (&controllers.PayCredentialsController{}).View(&ctx) })
			apiAuthV1.Get("/pay_credentials/delete", func(ctx iris.Context) { (&controllers.PayCredentialsController{}).Delete(&ctx) })

			apiAuthV1.Get("/manual_charges", func(ctx iris.Context) { (&controllers.ManualChargesController{}).Index(&ctx) })
			apiAuthV1.Post("/manual_charges/add", func(ctx iris.Context) { (&controllers.ManualChargesController{}).Save(&ctx) })
			apiAuthV1.Post("/manual_charges/update", func(ctx iris.Context) { (&controllers.ManualChargesController{}).Save(&ctx) })
			apiAuthV1.Get("/manual_charges/view", func(ctx iris.Context) { (&controllers.ManualChargesController{}).View(&ctx) })
			apiAuthV1.Get("/manual_charges/delete", func(ctx iris.Context) { (&controllers.ManualChargesController{}).Delete(&ctx) })
			apiAuthV1.Get("/manual_charges/allow", func(ctx iris.Context) { (&controllers.ManualChargesController{}).Allow(&ctx) })
			apiAuthV1.Get("/manual_charges/deny", func(ctx iris.Context) { (&controllers.ManualChargesController{}).Deny(&ctx) })

			apiAuthV1.Get("/manual_withdraws", func(ctx iris.Context) { (&controllers.ManualWithdrawsController{}).Index(&ctx) })
			apiAuthV1.Post("/manual_withdraws/add", func(ctx iris.Context) { (&controllers.ManualWithdrawsController{}).Save(&ctx) })
			apiAuthV1.Post("/manual_withdraws/update", func(ctx iris.Context) { (&controllers.ManualWithdrawsController{}).Save(&ctx) })
			apiAuthV1.Get("/manual_withdraws/view", func(ctx iris.Context) { (&controllers.ManualWithdrawsController{}).View(&ctx) })
			apiAuthV1.Get("/manual_withdraws/delete", func(ctx iris.Context) { (&controllers.ManualWithdrawsController{}).Delete(&ctx) })
			apiAuthV1.Get("/manual_withdraws/allow", func(ctx iris.Context) { (&controllers.ManualWithdrawsController{}).Allow(&ctx) })
			apiAuthV1.Get("/manual_withdraws/deny", func(ctx iris.Context) { (&controllers.ManualWithdrawsController{}).Deny(&ctx) })

			apiAuthV1.Get("/wash_code_infos", func(ctx iris.Context) { (&controllers.WashCodeInfosController{}).Index(&ctx) })
			apiAuthV1.Get("/wash_code_infos/view", func(ctx iris.Context) { (&controllers.WashCodeInfosController{}).View(&ctx) })
			apiAuthV1.Get("/wash_code_infos/delete", func(ctx iris.Context) { (&controllers.WashCodeInfosController{}).Delete(&ctx) })

			apiAuthV1.Get("/wash_code_records", func(ctx iris.Context) { (&controllers.WashCodeRecordsController{}).Index(&ctx) })
			apiAuthV1.Get("/wash_code_records/view", func(ctx iris.Context) { (&controllers.WashCodeRecordsController{}).View(&ctx) })
			apiAuthV1.Get("/wash_code_records/delete", func(ctx iris.Context) { (&controllers.WashCodeRecordsController{}).Delete(&ctx) })

			apiAuthV1.Get("/withdraw_dama_records", func(ctx iris.Context) { (&controllers.WithdrawDamaRecordsController{}).Index(&ctx) })
			apiAuthV1.Get("/withdraw_dama_records/view", func(ctx iris.Context) { (&controllers.WithdrawDamaRecordsController{}).View(&ctx) })
			apiAuthV1.Get("/withdraw_dama_records/delete", func(ctx iris.Context) { (&controllers.WithdrawDamaRecordsController{}).Delete(&ctx) })

			apiAuthV1.Get("/proxy_statistics", func(ctx iris.Context) { (&controllers.ProxyStatisticsController{}).Index(&ctx) })
			apiAuthV1.Get("/proxy_statistics/view", func(ctx iris.Context) { (&controllers.ProxyStatisticsController{}).View(&ctx) })
			apiAuthV1.Get("/proxy_statistics/delete", func(ctx iris.Context) { (&controllers.ProxyStatisticsController{}).Delete(&ctx) })

			apiAuthV1.Get("/system_statistics", func(ctx iris.Context) { (&controllers.SystemStatisticsController{}).Index(&ctx) })
			apiAuthV1.Get("/system_statistics/view", func(ctx iris.Context) { (&controllers.SystemStatisticsController{}).View(&ctx) })
			apiAuthV1.Get("/system_statistics/delete", func(ctx iris.Context) { (&controllers.SystemStatisticsController{}).Delete(&ctx) })
			apiAuthV1.Get("/system_statistics/counts", func(ctx iris.Context) { (&controllers.SystemStatisticsController{}).Counts(&ctx) })
			apiAuthV1.Get("/system_statistics/first_chanrge", func(ctx iris.Context) { (&controllers.SystemStatisticsController{}).FirstCharge(&ctx) })
			apiAuthV1.Get("/system_statistics/first_withdraw", func(ctx iris.Context) { (&controllers.SystemStatisticsController{}).FirstWithdraw(&ctx) })
			apiAuthV1.Get("/system_statistics/backwater", func(ctx iris.Context) { (&controllers.SystemStatisticsController{}).BackWater(&ctx) })

			apiAuthV1.Get("/proxy_commissions", func(ctx iris.Context) { (&controllers.ProxyCommissionsController{}).Index(&ctx) })
			apiAuthV1.Get("/proxy_commissions/view", func(ctx iris.Context) { (&controllers.ProxyCommissionsController{}).View(&ctx) })
			apiAuthV1.Get("/proxy_commissions/delete", func(ctx iris.Context) { (&controllers.ProxyCommissionsController{}).Delete(&ctx) })

			apiAuthV1.Get("/withdraw_records", func(ctx iris.Context) { (&controllers.WithdrawRecordsController{}).Index(&ctx) })
			apiAuthV1.Get("/withdraw_records/view", func(ctx iris.Context) { (&controllers.WithdrawRecordsController{}).View(&ctx) })
			apiAuthV1.Get("/withdraw_records/delete", func(ctx iris.Context) { (&controllers.WithdrawRecordsController{}).Delete(&ctx) })
			apiAuthV1.Get("/withdraw_records/allow", func(ctx iris.Context) { (&controllers.WithdrawRecordsController{}).Allow(&ctx) })
			apiAuthV1.Get("/withdraw_records/deny", func(ctx iris.Context) { (&controllers.WithdrawRecordsController{}).Deny(&ctx) })

			apiAuthV1.Get("/redpacket_systems", func(ctx iris.Context) { (&controllers.RedpacketSystemsController{}).Index(&ctx) })
			apiAuthV1.Post("/redpacket_systems/add", func(ctx iris.Context) { (&controllers.RedpacketSystemsController{}).Save(&ctx) })
			apiAuthV1.Post("/redpacket_systems/update", func(ctx iris.Context) { (&controllers.RedpacketSystemsController{}).Save(&ctx) })
			apiAuthV1.Get("/redpacket_systems/view", func(ctx iris.Context) { (&controllers.RedpacketSystemsController{}).View(&ctx) })
			apiAuthV1.Get("/redpacket_systems/delete", func(ctx iris.Context) { (&controllers.RedpacketSystemsController{}).Delete(&ctx) })

			apiAuthV1.Get("/redpacket_receives", func(ctx iris.Context) { (&controllers.RedpacketReceivesController{}).Index(&ctx) })
			apiAuthV1.Get("/redpacket_receives/view", func(ctx iris.Context) { (&controllers.RedpacketReceivesController{}).View(&ctx) })
			apiAuthV1.Get("/redpacket_receives/delete", func(ctx iris.Context) { (&controllers.RedpacketReceivesController{}).Delete(&ctx) })

			apiAuthV1.Get("/withdraw_records", func(ctx iris.Context) { (&controllers.WithdrawRecordsController{}).Index(&ctx) })
			apiAuthV1.Get("/withdraw_records/view", func(ctx iris.Context) { (&controllers.WithdrawRecordsController{}).View(&ctx) })
			apiAuthV1.Get("/withdraw_records/delete", func(ctx iris.Context) { (&controllers.WithdrawRecordsController{}).Delete(&ctx) })

			// 银商接口
			apiAuthV1.Get("/silver_merchant_users", func(ctx iris.Context) { (&controllers.SilverMerchantUsersController{}).Index(&ctx) })
			apiAuthV1.Post("/silver_merchant_users/add", func(ctx iris.Context) { (&controllers.SilverMerchantUsersController{}).Save(&ctx) })
			apiAuthV1.Post("/silver_merchant_users/update", func(ctx iris.Context) { (&controllers.SilverMerchantUsersController{}).Save(&ctx) })
			apiAuthV1.Get("/silver_merchant_users/view", func(ctx iris.Context) { (&controllers.SilverMerchantUsersController{}).View(&ctx) })
			apiAuthV1.Get("/silver_merchant_users/delete", func(ctx iris.Context) { (&controllers.SilverMerchantUsersController{}).Delete(&ctx) })
			apiAuthV1.Get("/silver_merchant_charge_cards", func(ctx iris.Context) { (&controllers.SilverMerchantChargeCardsController{}).Index(&ctx) })
			apiAuthV1.Post("/silver_merchant_charge_cards/add", func(ctx iris.Context) { (&controllers.SilverMerchantChargeCardsController{}).Save(&ctx) })
			apiAuthV1.Post("/silver_merchant_charge_cards/update", func(ctx iris.Context) { (&controllers.SilverMerchantChargeCardsController{}).Save(&ctx) })
			apiAuthV1.Get("/silver_merchant_charge_cards/view", func(ctx iris.Context) { (&controllers.SilverMerchantChargeCardsController{}).View(&ctx) })
			apiAuthV1.Get("/silver_merchant_charge_cards/delete", func(ctx iris.Context) { (&controllers.SilverMerchantChargeCardsController{}).Delete(&ctx) })
			apiAuthV1.Get("/silver_merchant_bank_cards", func(ctx iris.Context) { (&controllers.SilverMerchantBankCardsController{}).Index(&ctx) })
			apiAuthV1.Get("/silver_merchant_capital_flows", func(ctx iris.Context) { (&controllers.SilverMerchantCapitalFlowsController{}).Index(&ctx) })
			apiAuthV1.Get("/silver_merchant_user_charge", func(ctx iris.Context) { (&controllers.SilverMerchantUserChargeController{}).Index(&ctx) })
			apiAuthV1.Get("/silver_merchant_reports", func(ctx iris.Context) { (&controllers.SilverMerchantReportsController{}).Index(&ctx) })
			apiAuthV1.Get("/silver_merchant_charge_records", func(ctx iris.Context) { (&controllers.SilverMerchantChargeRecordsController{}).Index(&ctx) })
			apiAuthV1.Get("/silver_merchant_charge_records/view", func(ctx iris.Context) { (&controllers.SilverMerchantChargeRecordsController{}).View(&ctx) })
			apiAuthV1.Post("/silver_merchant_charge_records/allow", func(ctx iris.Context) { (&controllers.SilverMerchantChargeRecordsController{}).Allow(&ctx) })
			apiAuthV1.Post("/silver_merchant_charge_records/deny", func(ctx iris.Context) { (&controllers.SilverMerchantChargeRecordsController{}).Deny(&ctx) })
			apiAuthV1.Get("/silver_merchant_login_logs", func(ctx iris.Context) { (&controllers.SilverMerchantLoginLogsController{}).Index(&ctx) })
			apiAuthV1.Get("/silver_merchant_os_logs", func(ctx iris.Context) { (&controllers.SilverMerchantOsLogsController{}).Index(&ctx) })
		}
	}

	//需要登录之后才能访问的路由,//jwt中间件
	//apiAuthV1 := apiPlatform.Party("/admin/api/auth/v1", middlewares.JwtAuthenticate().Serve, middlewares.JwtHandler())
	//{
	//	apiAuthV1.Get("/modifyPwd", func(ctx iris.Context) { frontEndControllers.NewApiConfigController(ctx).ModifyPwd() })
	//	apiAuthV1.Get("/logout", func(ctx iris.Context) { frontEndControllers.NewApiConfigController(ctx).Logout() })
	//}
}
