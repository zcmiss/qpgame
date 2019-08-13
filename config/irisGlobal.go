package config

import (
	"os"
	"time"
)

type RedisSet struct {
	Addr         string //host:port
	Password     string //数据库密码
	DB           int    //选择的数据库
	MaxRetries   int    //重试最大次数
	PoolSize     int    //连接池大小
	MinIdleConns int    //最小空闲数
}

type MysqlSet struct {
	AliasName string //数据库别名
	DriveName string //驱动名称
	DriveDsn  string //数据库连接地址
	MaxIdle   int    //最大空闲数
	MaxConn   int    //最大连接数
}

type WxSet struct {
	AppId     string
	AppSecret string
}

var DevloperDebug bool = false //是否是开发阶段,这里部署的时候会自动替换

var IsApiInterfaceServer bool = true //是否是api接口服务

var IsAdminApiInterfaceServer bool = false //是否是后台管理服务器接口

var IsSilverMerchantServer bool = false //是否是银商服务器

var IsTimerServer bool = false //是否是投注记录采集定时器服务器
//3天
const JwtTokenExp = 3600 * 72 //ios,安卓客户端token过期时间

const AdminTokenExpire = time.Hour * 1 // 后台的token过期时间

const SilverMerchantTokenExpire = 3600

const TokenKey = "qpgame_20190325Dream"   //ios,安卓前端接口token加密密钥
const CodeKey = "agheury76j982ruj"        //注册验证码
const CodeIv = "kgjdsierjgkdfjss"         //注册验证码
var PlatformCPs = map[string]interface{}{ //合法游戏平台标识
	"CKYX": map[string]interface{}{
		//如果有需要在配置
		"mysql": MysqlSet{
			AliasName: "CKYXMYSQL",
			DriveName: "mysql",
			//DriveDsn:  "root:root@tcp(192.168.0.104:3306)/ckgame?charset=utf8&multiStatements=true",
			DriveDsn:  "root:root@tcp(127.0.0.1:3306)/ckgame?charset=utf8&multiStatements=true",
			MaxIdle:   10,
			MaxConn:   30,
		},
		"wx": WxSet{
			AppId:     "wx24aa8afd39c6163b",
			AppSecret: "b44e2f95146644468f0f38d89bdc961f",
		},
	},
	"QQYX": map[string]interface{}{
		//如果有需要在配置
		"mysql": MysqlSet{
			AliasName: "QQYXMYSQL",
			DriveName: "mysql",
			//DriveDsn:  "root:root@tcp(192.168.0.104:3306)/qqgame?charset=utf8&multiStatements=true",
			DriveDsn:  "root:root@tcp(127.0.0.1:3306)/qqgame?charset=utf8&multiStatements=true",
			MaxIdle:   10,
			MaxConn:   30,
		},
		"wx": WxSet{
			AppId:     "wx8a9dbde68d3a3bd8",
			AppSecret: "f29368f5a2aeb4ba716b29fc91f02bc0",
		},
	},
}

var MainDb MysqlSet //主数据库配置,所有游戏平台都用到的数据

// 模块初始化,进行环境处理
func init() {
	//主数据库单独处理
	MainDb = MysqlSet{
		AliasName: "maindb",
		DriveName: "mysql",
		DriveDsn:  "root:root@tcp(127.0.0.1:3306)/maindb?charset=utf8&multiStatements=true",
		MaxIdle:   5,
		MaxConn:   30,
	}
	//服务器分发检测
	checkApiOrAdminApi()

	//如果是发布模式就修改debug,并且移除cptest
	if os.Getenv("IRIS_MODE") == "release" {
		delete(PlatformCPs, "QPTEST")
		//每个运营平台独占一台服务器，移除其他不相干平台
		//不存在平台号就加载全部,在调试或者备用服务器上会用到
		curPlatform := os.Getenv("CURRENTPLATFORM")
		if PlatformCPs[curPlatform] != nil {
			deleteNotAssignPlatform([]string{curPlatform})
		}
		//开发模式可以随意调试指定平台
	} else {
		DevloperDebug = true
		//如果要调试指定平台就打开这个方法注释,否则加载所有平台
		//deleteNotAssignPlatform([]string{"QPTEST"})
	}
}

//检测是前端api服务器还是后台管理接口服务器
func checkApiOrAdminApi() {
	//如果是api接口服务器
	if os.Getenv("ISAPISERVER") == "yes" {
		IsApiInterfaceServer = true
		IsAdminApiInterfaceServer = false
		// 如果是后台管理接口服务器
	} else {
		IsApiInterfaceServer = false
		IsAdminApiInterfaceServer = true
	}
	//如果是投注记录采集定时器服务器
	if os.Getenv("ISTIMERSERVER") == "yes" {
		IsApiInterfaceServer = false
		IsAdminApiInterfaceServer = false
		IsTimerServer = true
	}
	//如果是银商系统服务器
	if os.Getenv("ISSIlVERMERCHAN") == "yes" {
		IsApiInterfaceServer = false
		IsAdminApiInterfaceServer = false
		IsTimerServer = false
		IsSilverMerchantServer = true
	}

}

//删除非指定平台
func deleteNotAssignPlatform(platforms []string) {
	for key, _ := range PlatformCPs {
		isDelete := false
		//如果没有这这个数组里面就移掉
		for _, v := range platforms {
			if key == v {
				isDelete = true
			}
		}
		if !isDelete {
			delete(PlatformCPs, key)
		}
	}
}
