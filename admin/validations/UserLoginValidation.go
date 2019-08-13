package validations

import (
	"qpgame/admin/common"
	"qpgame/common/mvc"
	"qpgame/common/utils"
	"qpgame/models"
	"strings"

	"github.com/mojocn/base64Captcha"
)

type UserLoginValidation struct{}

// 添加/修改动作数据验证
func (self UserLoginValidation) Validate(ctx *Context, row *map[string]string) (string, bool) {
	platform := (*ctx).Params().Get("platform")
	//1. 对于提交的验证码的处理
	getVerifyCode := func(ctx *Context) string {
		postData := utils.GetPostData(ctx)
		submittedCode := postData.Get("code")       //提交的验证码
		submittedCodeKey := postData.Get("codeKey") //提交的验证码标识
		if submittedCodeKey == "" {
			return ""
		}
		key := platform + "-" + submittedCodeKey
		verifyCode, exists := common.AdminVerifyCodes[key]
		if !exists {
			return ""
		}
		success := base64Captcha.VerifyCaptcha(verifyCode[0], submittedCode)
		delete(common.AdminVerifyCodes, key)
		if success {
			return submittedCode
		}
		return ""
	}
	errMessage, pass := mvc.NewValidation(ctx).
		StringLength("username", "用户名称长度必须在5-20之间", 5, 20).
		StringLength("password", "用户密码长度必须在5-20之间", 6, 20).
		StringEquals("code", "请输入正确格式的验证码", getVerifyCode(ctx)).
		Validate()
	if !pass {
		return errMessage, pass
	}

	form := utils.GetPostData(ctx) //提交的post数据

	//2. 用户名密码验证
	userName := form.Get("username")
	sql := "SELECT * FROM admins WHERE name = '" + userName + "' LIMIT 1"
	conn := models.MyEngine[platform]        //数据库连接
	rows, err := conn.SQL(sql).QueryString() //拿到登录用户信息
	if err != nil || len(rows) == 0 {
		return "登录失败: 用户不存在", false
	}

	*row = rows[0]
	password := form.Get("password")
	if strings.Compare((*row)["password"], utils.MD5(password)) != 0 {
		return "登录失败: 用户密码错误", false
	}

	if strings.Compare((*row)["role_id"], "") == 0 || strings.Compare((*row)["role_id"], "0") == 0 {
		return "登录失败: 当前用户角色的权限不足", false
	}

	return "", true
}
