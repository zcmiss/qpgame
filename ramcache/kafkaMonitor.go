package ramcache

import (
	"github.com/Shopify/sarama"
	"github.com/buger/jsonparser"
	"os"
	"qpgame/config"
	"sync"
)

var (
	wg sync.WaitGroup
)

//启动kafka
func MonitoringKafka() {
	var brokerLIst string
	//生产环境配置
	brokerLIst = "aws.qpgame.cache:9092"
	//主题订阅,默认订阅所有平台
	//主要是应用在备份服务器上，备用服务器为所有平台共用
	topic := "awsGameCache"
	curPlatform := os.Getenv("CURRENTPLATFORM")
	//如果环境变量中有指定当前平台,那么只订阅对应平台的频道
	if config.PlatformCPs[curPlatform] != nil {
		topic = curPlatform + "_" + topic
	}
	consumer, err := sarama.NewConsumer([]string{brokerLIst}, nil)
	if err != nil {
		panic(err)
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				kafkaMessageDIspath(msg.Value)
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}

//消息处理
func kafkaMessageDIspath(value []byte) {
	byteAction, _, _, _ := jsonparser.Get(value, "action")
	switch string(byteAction) {
	//主库platform表
	case "maindb_platform":
		mainDbPlatform(value)
	//运营平台活动表
	case "table_activities":
		platformDbTableActivities(value)
	//运营平台活动分类表
	case "table_activity_classes":
		platformDbTableActivityClasses(value)
	//运营平台版本管理表
	case "table_app_versions":
		platformDbTableAppVersions(value)
	//运营平台投注记录拉取配置表
	case "table_bets_key":
		platformDbTableBetsKey(value)
	//运营平台充值账号表
	case "table_charge_cards":
		platformDbTableChargeCards(value)
	//运营平台支付分类表
	case "table_charge_types":
		platformDbTableChargeTypes(value)
	//运营平台配置表
	case "table_configs":
		platformDbTableConfigs(value)
	//运营平台游戏分类表
	case "table_game_categories":
		platformDbTableGameCategories(value)
	//运营平台站内通知表
	case "table_notices":
		platformDbTableNotices(value)
	//运营平台支付证书配置表
	case "table_pay_credentials":
		platformDbTablePayCredentials(value)
	//运营平台第三方游戏账号表
	case "table_platform_accounts":
		platformDbTablePlatformAccounts(value)
	//运营平台第三方游戏分类表
	case "table_platform_games":
		platformDbTablePlatformGames(value)
	//运营平台活第三方游戏平台表
	case "table_platforms":
		platformDbTablePlatforms(value)
	//运营平台棋牌代理等级表
	case "table_proxy_chess_levels":
		platformDbTableProxyChessLevels(value)
	//运营平台真人视讯等级表
	case "table_proxy_real_levels":
		platformDbTableProxyRealLevels(value)
	//运营平台系统通知等级表
	case "table_system_notices":
		platformDbTableSystemNotices(value)
	//运营平台银行卡列表
	case "table_user_banks":
		platformDbTableUserBanks(value)
	//运营平台用户表
	case "table_users":
		platformDbTableUsers(value)
	//运营平台用户组表
	case "table_user_groups":
		platformDbTableUserGroups(value)
	//运营平台vip等级表
	case "table_vip_levels":
		platformDbTableVipLevels(value)
	}
}
