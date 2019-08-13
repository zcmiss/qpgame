package common

// api网站配置
type ApiConfig struct {
	Code   string `json:"code"`
	Name   string `json:"name"`    //网站名称
	Url    string `json:"url"`     //网站地址
	ApiUrl string `json:"api_url"` //api接口地址
}

// 用户信息-登录时用到
type UserInfo struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	RealName        string `json:"real_name"`
	MaxManualAmount int    `json:"max_manual_amount"`
	RoleId          int    `json:"role_id"`
}

// 节点信息 //登录之后获取的节点信息，全部是需要授权的
type MenuNode struct {
	Id    int        `json:"id"`    //编号
	Title string     `json:"title"` //显示名称
	Level int        `json:"level"` //级别,1:主菜单;2:子菜单;3:标签页;4:功能
	Url   string     `json:"url"`   //api地址
	Nodes []MenuNode `json:"menus"` //菜单
}

// 登录之后返回的信息
type LoginInfo struct {
	Token string     `json:"token"`
	Info  UserInfo   `json:"info"`
	Menus []MenuNode `json:"menus"`
}

// 默认结果类型
type Result struct {
	Message string
	Data    interface{}
}

// 管理员结构信息
type AdminUser struct {
	Id       int    //用户编号
	Name     string //用户名称
	Password string //用户密码
	RoleId   int    //角色编号
	Status   int    //状态
}

type Error struct {
	What string //错误信息
}

func (self Error) Error() string {
	return self.What
}
