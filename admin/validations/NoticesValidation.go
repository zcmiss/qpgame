package validations

import (
	"qpgame/admin/common"
	"qpgame/common/mvc"
	"qpgame/common/utils"
)

type NoticesValidation struct{}

// 添加/修改动作数据验证
func (self NoticesValidation) Validate(ctx *Context) (string, bool) {
	platform := (*ctx).Params().Get("platform")
	post := utils.GetPostData(ctx)
	userName := post.Get("user_id")
	realName := "NoUser"
	//检测此用户是否存在
	idStr := common.GetIdFromUserName(platform, userName)
	if idStr != "" {
		(*ctx).Params().Set("user_id", idStr)
		realName = userName
	}
	return mvc.NewValidation(ctx).
		StringLength("title", "标题长度需要介于2-30之间", 2, 30).
		NotNull("content", "请输入内容的内容").
		IsNumeric("status", "状态必须为数字").
		StringEquals("user_id", "用户不存在", realName).
		Validate()
}
