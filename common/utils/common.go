package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
	"os"
	"qpgame/config"
	"qpgame/models"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"

	"github.com/kataras/iris"
)

//ip地址数据库保存的位置
var ipDbPath string = ""

//地址转转地象
var ipTransfer *ip2region.Ip2Region

// 必填参数验证
func RequiredParam(c *iris.Context, params []string) bool {
	return requireParamVerify(c, params, "get")
}

// 必填参数验证
func RequiredParamPost(c *iris.Context, params []string) bool {
	return requireParamVerify(c, params, "post")
}

//get,post参数请求验证
func requireParamVerify(c *iris.Context, params []string, requestType string) bool {
	for _, par := range params {
		param := (*c).URLParam(par)
		if requestType == "post" {
			param = (*c).FormValue(par)
		}
		if param == "" {
			ResFaiJSON(c, "缺少参数:"+par, "缺少参数:"+par, config.PARAMERROR)
			return false
		}
	}
	return true
}

func ValidRequiredPostData(ctx iris.Context, data PostData, requireFields []string) bool {
	for _, field := range requireFields {
		if data.Get(field) == "" {
			ResFaiJSON(&ctx, "", "缺少参数:"+field, config.PARAMERROR)
			return false
		}
	}
	return true
}

// json转换
func JsonEncode(content interface{}) (bufInString string, mtsIsOk bool) {
	mapToString, _ := json.Marshal(content)
	//是否压缩成功
	mtsIsOk = true
	var bufIn bytes.Buffer
	gz, _ := gzip.NewWriterLevel(&bufIn, 9)
	if _, err := gz.Write(mapToString); err != nil {
		mtsIsOk = false
	}
	if err := gz.Close(); err != nil {
		mtsIsOk = false
	}
	bufInString = string(bufIn.Bytes())
	return bufInString, mtsIsOk
}

// gzip压缩功能
func ByteToGzip(content []byte) (bufInString string, mtsIsOk bool) {
	//是否压缩成功
	mtsIsOk = true
	var bufIn bytes.Buffer
	gz, _ := gzip.NewWriterLevel(&bufIn, 9)
	if _, err := gz.Write(content); err != nil {
		mtsIsOk = false
	}
	if err := gz.Close(); err != nil {
		mtsIsOk = false
	}
	bufInString = string(bufIn.Bytes())
	return bufInString, mtsIsOk
}

// 根据字符串调用对应对象的方法
func CallMethodReflect(any interface{}, methodName string, args ...interface{}) (reflectValue []reflect.Value, runOk bool) {
	inputs := make([]reflect.Value, len(args))

	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	if v := reflect.ValueOf(any).MethodByName(methodName); v.String() == "<invalid Value>" {
		runOk = false
		return []reflect.Value{}, runOk
	} else {
		runOk = true
		reflectValue = v.Call(inputs)
		return reflectValue, runOk
	}
}

// md5加密
func MD5(message string) string {
	res := md5.Sum([]byte(message))
	return fmt.Sprintf("%x", res)
}

// md5加密16
func MD5By16(message string) string {
	res := md5.Sum([]byte(message))
	return fmt.Sprintf("%x", res)[8:24]
}

// 得到用户的IP
func GetIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

//获取当前时间转换成int
func GetNowTime() int {
	return int(time.Now().Unix())
}

//获取格式化时间yyyyMMddHHmmssSSS
func GetFmtTime() string {
	return strings.Replace(time.Now().Format("20060102150405.000"), ".", "", 1)
}

//由时间字符串生成时间戳
func GetInt64FromTime(value string) int64 {
	var currentTime int64 = 0
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if v, err := time.ParseInLocation("2006-01-02 15:04:05", value, loc); err == nil {
		currentTime = v.Unix()
	}
	return currentTime
}

//由时间字符串生成时间戳
func GetInt64FromDate(value string) int64 {
	var currentTime int64 = 0
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if v, err := time.ParseInLocation("2006-01-02", value, loc); err == nil {
		currentTime = v.Unix()
	}
	return currentTime
}

//时间查询条件转换
func GetQueryTime(timed int) (string, string) {
	started := 0
	ended := 0
	switch timed {
	case 1:
		st, _ := time.Parse("20060102", time.Now().Format("20060102"))
		started = int(st.Unix())
		ended = int(st.Add(time.Hour * 24).Add(time.Second * -1).Unix())
	case 2:
		yesterday := time.Now().Add(time.Hour * -24)
		st, _ := time.Parse("20060102", yesterday.Format("20060102"))
		started = int(st.Unix())
		ended = int(st.Add(time.Hour * 24).Add(time.Second * -1).Unix())
	case 0, 3:
		started = int(time.Now().AddDate(0, 0, -30).Unix())
		ended = int(time.Now().Unix())
	}
	return strconv.Itoa(started), strconv.Itoa(ended)
}

// 是否是日期格式
func IsDate(value string) bool {
	matched, err := regexp.MatchString("^\\d{4}\\-\\d{2}\\-\\d{2}$", value)
	return err == nil && matched
}

//是否是日期时间格式
func IsDatetime(value string) bool {
	matched, err := regexp.MatchString("^\\d{4}\\-\\d{2}\\-\\d{2}\\s+\\d{1,2}:\\d{1,2}:\\d{1,2}$", value)
	return err == nil && matched
}

//得到ip地址相关信息, 格式: 国家-省-城市
func GetIpInfo(ipAddress string) string {
	if ipDbPath == "" {
		pwd, pwdErr := os.Getwd()
		if pwdErr == nil {
			ipDbPath = pwd + "/config/ip2region.db"
		}
	}
	if ipDbPath != "" {
		ipTransfer, _ = ip2region.New(ipDbPath)
		defer ipTransfer.Close()
		ip, ipErr := ipTransfer.BinarySearch(ipAddress)
		if ipErr != nil {
			return "解析失败"
		}
		if ip.Country == "0" && ip.Province == "0" && ip.Country == "0" && ip.Region == "" {
			return "未知地区"
		}
		ipStr := ""
		if ip.Country != "0" {
			ipStr += ip.Country
		}
		if ip.Province != "0" {
			ipStr += "-" + ip.Province
		}
		if ip.City != "0" {
			ipStr += "-" + ip.City
		}
		if ip.Region != "0" {
			ipStr += "-" + ip.Region
		}
		return ipStr
	}
	return "IP加载错误"
}

// 得到xml文本
func GetXmlText(xmlContent []byte, name ...string) map[string]string {
	decoder := xml.NewDecoder(bytes.NewBuffer(xmlContent))
	startname := ""
	res := make(map[string]string)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		if err != nil {
			return nil
		}
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			startname = token.Name.Local
		case xml.CharData:
			for _, s := range name {
				if startname == s {
					content := string([]byte(token))
					res[s] = content
					startname = ""
				}
			}

		}
	}
	return res
}

// 计算日期时间范围，例：GetDatetimeRange(0, 1)，返回当天00:00:00至第二天00:00:00的时间戳
func GetDatetimeRange(start int64, length int64) (int64, int64) {
	var timeLoc, _ = time.LoadLocation("Asia/Shanghai")
	nowTime := time.Now().In(timeLoc)
	sFromTime := nowTime.Format("2006-01-02")
	fromTime, _ := time.ParseInLocation("2006-01-02", sFromTime, timeLoc)
	var iDaySec int64 = 3600 * 24
	iFromTime := fromTime.Unix() + iDaySec*start
	iToTime := iFromTime + iDaySec*length
	return iFromTime, iToTime
}

// 得到当前请求所在的平台识别号
func GetPlatform(ctx *iris.Context) string {
	return (*ctx).Params().Get("platform")
}

// 得到当前平台的数据库连接
func GetDb(ctx *iris.Context) *xorm.Engine {
	return models.MyEngine[GetPlatform(ctx)]
}

// 得到指定平台的数据库连接
func GetDbForPlatform(platform string) *xorm.Engine {
	return models.MyEngine[platform]
}

// 时间戳转字符串
func TimestampToDateStr(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)
	if format == "1" {
		format = "2006-01-02 15:04:05"
	}
	return tm.Format(format)
}
