package xorm

import (
	"github.com/kataras/iris/core/errors"
	"qpgame/models"
	"time"
)

type Users struct {
	Id              int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Phone           string `xorm:"not null default '' comment('手机号') CHAR(11)"`
	Password        string `xorm:"-> not null comment('密码') CHAR(40)"`
	UserName        string `xorm:"not null comment('用户名') unique VARCHAR(20)"`
	Name            string `xorm:"not null comment('用户姓名') unique CHAR(30)"`
	Email           string `xorm:"not null default '' comment('邮箱') VARCHAR(50)"`
	Created         int    `xorm:"not null comment('创建时间') INT(10)"`
	Birthday        string `xorm:"not null comment('生日') VARCHAR(10)"`
	MobileType      int    `xorm:"not null default 1 comment('1andorid,2ios') TINYINT(2)"`
	Sex             int    `xorm:"not null default 1 comment('性别:1男,2女') TINYINT(2)"`
	Path            string `xorm:"comment('代理层级id路径例子:(,1,2,4,5,7,)') TEXT"`
	VipLevel        int    `xorm:"not null default 0 comment('vip等级(1-10)') INT(10)"`
	Qq              string `xorm:"not null default '' comment('用户QQ号') VARCHAR(15)"`
	Wechat          string `xorm:"not null default '' comment('用户微信号') VARCHAR(30)"`
	Status          int    `xorm:"not null default 1 comment('用户状态,1正常，0锁定') TINYINT(1)"`
	ProxyStatus     int    `xorm:"not null default 1 comment('代理状态,1正常，0锁定') TINYINT(1)"`
	UserType        int    `xorm:"not null default 0 comment('0正常用户1游客用户2虚拟用户') TINYINT(3)"`
	Token           string `xorm:"not null default '' comment('用户登录token,要保持到程序内存中') VARCHAR(200)"`
	RegIp           string `xorm:"not null default '' comment('注册IP') VARCHAR(15)"`
	UniqueCode      string `xorm:"not null default '' comment('手机唯一标识') VARCHAR(50)"`
	TokenCreated    int    `xorm:"not null comment('token创建时间,根据这个来双层判断是否已过期') INT(10)"`
	SafePassword    string `xorm:"not null default '' comment('保险箱密码') VARCHAR(40)"`
	UserGroupId     string `xorm:"not null default '1' comment('用户组id') INT(10)"`
	ParentId        int    `xorm:"not null default 0 comment('上级代理用户Id') INT(11)"`
	LastLoginTime   int    `xorm:"not null default 0 comment('上次登录时间') INT(10)"`
	LastPlatformId  int    `xorm:"comment('上次登录的游戏平台') INT(11)"`
	GroupSize       int    `xorm:"-"` //  <-只从数据库读取   ->只写入到数据库   -不进行字段映射
	WxOpenId        string `xorm:"not null default '' comment('微信openID') VARCHAR(32)"`
	IsModifiedUname int    `xorm:"not null default 0 comment('是否修改过用户名,0未修改，1已修改') TINYINT(1)"`
}

func DefaultUser() Users {
	return Users{Email: "", Created: int(time.Now().Unix()), MobileType: 1, Sex: 1, VipLevel: 1, Status: 1, UserType: 0, ParentId: 1}
}

func GetUser(platform string, idOrPhone interface{}) (Users, error) {
	var users []Users
	err := models.MyEngine[platform].SQL("select u.*,(SELECT COUNT(*) FROM users t WHERE t.parent_id=u.id) group_size from users u where (u.id=? or u.phone=? or u.user_name = ?)", idOrPhone, idOrPhone, idOrPhone).Find(&users)
	//如果没有数据
	if len(users) == 0 && err == nil {
		return users[0], errors.New("没有数据")
	}
	return users[0], err
}
