package validations

type BlacklistsValidation struct{}

// 添加/修改动作数据验证
func (self BlacklistsValidation) Validate(ctx *Context) (string, bool) {
	return "", true
	//return mvc.NewValidation(ctx).
	//	IsNumeric("user_id", "用户id必须为数字").
	//	Validate()
}
