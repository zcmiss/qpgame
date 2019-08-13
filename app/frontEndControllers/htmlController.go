package frontEndControllers

import (
	"encoding/base64"
	"fmt"
	"github.com/kataras/iris"
	"github.com/skip2/go-qrcode"
	"net/url"
)

type HtmlController struct{}

//单例对象
var HtmlC *HtmlController

func init() {
	HtmlC = new(HtmlController)
}

//二维码生成,性能很差1s,270多 这个接口是为了兼容老的版本,新的使用客户端本地生成
func (cthis *HtmlController) QrcodeText() (handler iris.Handler) {
	return func(ctx iris.Context) {
		text := ctx.Params().Get("text")
		textByte,_ := base64.StdEncoding.DecodeString(text)
		text,_ = url.QueryUnescape(string(textByte))
		png, _ := qrcode.Encode(text, qrcode.Medium, 128)
		ctx.Header("Content-Type", "image/png")
		ctx.Header("Content-Length", fmt.Sprintf("%d", len(png)))
		ctx.Write(png)
	}
}
