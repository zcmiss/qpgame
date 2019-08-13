package websockets

//上传的命令
type WsCommand struct {
	Platform string `json:"platform"` //平台标识号
	Type     string `json:"type"`     //类型/动作
	Id       int    `json:"id"`       //后台用户编号
}

//带id的上传的命令
type WsCommandNoId struct {
	Platform string `json:"platform"` //平台标识号
	Type     string `json:"type"`     //类型/动作
}

//用于返回的结果
type WsResult struct {
	Type         string      `json:"type"`         //类型
	ClientMsg    string      `json:"clientMsg"`    //外部错误信息
	Data         interface{} `json:"data"`         //结果数据
	TimeConsumed int64       `json:"timeConsumed"` //花费时间
}
