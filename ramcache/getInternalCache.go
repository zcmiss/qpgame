package ramcache

import "qpgame/models/mainxorm"

//获取系统配置信息
func GetConfigs(platform string, keyName string) interface{} {
	config, _ := TableConfigs.Load(platform)
	return config.(map[string]interface{})[keyName]
}

//获取主库platform表指定平台的数据
func GetMainDbPlatform(platform string) mainxorm.Platform {
	config, _ := MainTablePlatform.Load("platform")
	configs := config.([]mainxorm.Platform)
	for _, val := range configs {
		if val.Code == platform {
			return val
		}
	}

	return mainxorm.Platform{}
}
