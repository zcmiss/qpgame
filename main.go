package main

import (
	"fmt"
	admin "qpgame/admin/common"
	"qpgame/admin/jobs"
	"qpgame/admin/websockets"
	"qpgame/common/timer"
	"qpgame/config"
	"qpgame/models"
	"qpgame/ramcache"
	"qpgame/routers"

	"github.com/kataras/iris"
)

func main() {

	app := iris.New()
	models.CreateXormMysqlConnectionPool() //建立mysql连接池 xorm

	//加载ios,安卓客户端接口服务配置
	if config.IsApiInterfaceServer {
		routers.AppConfigRouters(app) //加载客户端前台路由
		ramcache.MainDbLoadCache()    //主库缓存加载
		ramcache.LoadCache()          //加载全局缓存
		go timer.InitProxyTimerTask()

	}

	//后台管理初始化内容
	if config.IsAdminApiInterfaceServer {
		ramcache.LoadTableFields()   //加载平台-数据库-表-字段信息
		admin.LoadAll()              //加载后台的所有菜单/权限
		websockets.Init()            //启动websocket服务,此项一定要列于加载全部路由之前
		routers.AdminApiRouters(app) //加载后台管理系统的路由
		go jobs.InitAdminJobs()      //用于后台的定时任务
	}

	//如果是投注记录采集定时器服务器
	if config.IsTimerServer {
		ramcache.LoadCache() //加载全局缓存
		go timer.InitTimerTask()
	}

	//异步启动kafka,只针对前端接口服务或者银商服务器
	if (!config.DevloperDebug && config.IsApiInterfaceServer) || config.IsSilverMerchantServer {
		fmt.Println("加载kafka...")
		go ramcache.MonitoringKafka()
	}
	//如果是银商系统服务器
	if config.IsSilverMerchantServer {
		routers.SilverMerchantRouters(app) //加载银商系统接口路由
		ramcache.MainDbLoadCache()         //主库缓存加载
		ramcache.LoadCache()
	}

	//启动服务
	app.Run(iris.Addr(":39100"), iris.WithConfiguration(iris.YAML("./config/iris.yml")))
}
