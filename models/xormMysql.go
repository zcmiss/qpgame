package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"qpgame/config"
	"time"
)

//多平台多个数据库对应
var MyEngine = make(map[string]*xorm.Engine)

//主库数据库特殊处理
var MyEngineMainDb *xorm.Engine

//创建mysql数据库连接池
func CreateXormMysqlConnectionPool() {

	//注册数据库连接
	for k, v := range config.PlatformCPs {
		masterOrSlaveConfig := v.(map[string]interface{})["mysql"].(config.MysqlSet)
		registerDataBaseAction(k, masterOrSlaveConfig)
	}
	//注册主数据库,非平台数据库单独做处理
	registerDataBaseAction("maindb", config.MainDb)
	fmt.Println("[棋牌游戏] mysql连接池创建成功......")
}

//注册数据库连接
func registerDataBaseAction(masterOrSlave string, masterOrSlaveConfig config.MysqlSet) {
	db, err := xorm.NewEngine(masterOrSlaveConfig.DriveName, masterOrSlaveConfig.DriveDsn)

	if err != nil {
		fmt.Println("退出程序,检查配置是否正确")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(masterOrSlave + " 连接mysql连接池创建失败,马上检查mysql配置是否正确,程序已经终止运行....")
		os.Exit(1)
	}
	db.SetMaxIdleConns(masterOrSlaveConfig.MaxIdle)
	db.SetMaxOpenConns(masterOrSlaveConfig.MaxConn)
	//连接生命周期
	db.SetConnMaxLifetime(time.Duration(9 * time.Second))
	//主库单独存放
	if masterOrSlave == "maindb" {
		MyEngineMainDb = db
	} else {
		//保存各个平台的连接池对象
		MyEngine[masterOrSlave] = db
	}
	//如果是调试模式，就打印sql语句
	db.ShowSQL(config.DevloperDebug)
}
