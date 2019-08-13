package ramcache

import "sync"

// 缓存user id索引
var UserIdCard sync.Map

// 缓存有效的tonken,缓存所有手机号码
var PhoneNumAndToken sync.Map

// 缓存有效的tonken,缓存所有用户名
var UserNameAndToken sync.Map

// 缓存有效的tonken,缓存所有唯一标识
var UniqueCodeAndToken sync.Map

// 缓存通过微信OpenId获取userId的字典
var WxOpenIdIndex sync.Map

//缓存所有游戏分类信息
var GameCache sync.Map

// 缓存游戏平台API接口的配置信息
var GamePlatformAPiConfigs sync.Map

//缓存用户游戏平台账号
var TablePlatformAccounts sync.Map

//缓存采集记录
var TableBetsKey sync.Map

//活动表
var TableActivities sync.Map

//活动分类表
var TableActivityClasses sync.Map

//缓存游戏分类表
var TableGameCategories sync.Map

//平台列表
var TablePlatforms sync.Map

//平台游戏列表
var TablePlatformGames sync.Map

//JDB平台游戏列表
var TablePlatformGamesByJDB sync.Map

//平台版本列表
var TableAppVersions sync.Map

//平台系统公告表
var TableSystemNotices sync.Map

//配置表
var TableConfigs sync.Map

//充值类型表
var TableChargeTypes sync.Map

//充值账号表
var TableChargeCards sync.Map

//第三方支付证书表
var TablePayCredential sync.Map

//vip等级表
var TableVipLevels sync.Map

//投注记录查询条件缓存
var BetsSearchType sync.Map

//棋牌代理等级表
var TableProxyChessLevels sync.Map

//银行卡列表
var TableUserBanks sync.Map

//真人代理等级表
var TableProxyRealLevels sync.Map

// 任务定时处理的任务，例如平台余额转入异常
var TableExceptionTasks sync.Map
