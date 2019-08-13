package pool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"qpgame/config"
	"time"
)

//mysql对象容器，需要用到redis直接到这里获取
var MysqlClientMap = make(map[string]*sql.DB)

//创建mysql数据库连接池
func CreateMysqlConnectionPool() {
	//注册数据库连接
	for key, value := range config.PlatformCPs {
		masterOrSlave := value.(map[string]interface{})["mysql"].(config.MysqlSet)
		registerDataBaseAction(key, masterOrSlave)
	}
	fmt.Println("[棋牌游戏] mysql连接池创建成功......")
}

//注册数据库连接
func registerDataBaseAction(masterOrSlave string, masterOrSlaveConfig config.MysqlSet) {
	db, err := sql.Open(masterOrSlaveConfig.DriveName, masterOrSlaveConfig.DriveDsn)
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
	//保存各个平台的连接池对象
	MysqlClientMap[masterOrSlave] = db

}
