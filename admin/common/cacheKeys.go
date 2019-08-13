package common

//后台的权限菜单
// key: 平台名称 + 菜单编号, value: 菜单列表
var AdminMenus = map[string][]MenuNode{}

// 后台的权限菜单
// key: 平台名称 + 角色编号, value: 菜单列表
var AdminRoleMenus = map[string][]MenuNode{}

// 保存的token信息
// key: md5(Token), value: Token
var AdminTokens = map[string]string{}

// key: 平台标识 + 管理员编号, value: 管理员信息
var Admins = map[string]AdminUser{}

// key: 平台标识 + 角色编号, value: 角色信息
var AdminRoles = map[string]string{}

// key: 平台标识 + 域名, value: 配置信息
var AdminApiConfigs = map[string]ApiConfig{}

// key: 平台标识 + ip, value: 验证码
var AdminVerifyCodes = map[string][]string{}

// key: 平台标识 + ip, value: 验证码
var SliverMerchantVerifyCodes = map[string][]string{}

// key: 平台标识 + 游戏编号, value: 游戏名称
var PlatformGames = map[string]string{}

// key 平台标识 + 游戏编码, value: 游戏名称
var GameCodes = map[string]string{}

// key: 平台标识 + 游戏平台编号, value: 平台名称
var GamePlatforms = map[string]string{}

// key: 平台标识 + 分类编号, value: 分类名称
var GameCategories = map[string]string{}

// 活动分类信息
var ActivityClasses = map[string]string{}

// 活动信息
var Activities = map[string]string{}

//前台用户
// key: 平台标识, value: map[用户编号]用户名称
var FrontendUsers = map[string]map[string]string{}

// 充值类型
//key: 平台识别号 +id, value: 充值类型名称
var ChargeTypes = map[string]string{}

// 充值类型
//key: 平台识别号 +id, value: 充值卡银行 + "/" + 姓名
var ChargeCards = map[string]string{}

// 充值方式
//key: 平台识别号 +id, value: 充值卡银行 + "/" + 姓名
var ThirdPayments = map[string]string{}

// 红包数据
// key: 平台识别号, value: 红包编号-> 红包数据
var Redpackets = map[string]map[string]map[string]string{}
