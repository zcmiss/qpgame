package utils

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
)

// 实现 Unmarshal
type PostData struct {
	Type int //0: form, 1: body
	Ctx  *iris.Context
	Data *[]byte
	//body string
	//Items map[string]interface{}
}

//错误定义
//type PostDataError struct { }
////错误显示
//func (self *PostDataError) Error() string {
//	return "从json转化时出现错误"
//}

// 封装get方法, 默认取得是字符串
func (self *PostData) Get(field string) string {
	if self.Type == 0 {
		return (*self.Ctx).FormValue(field)
	}

	v, _, _, err := jsonparser.Get(*self.Data, field)
	if err != nil { //再次尝试从 FormValues当中读取
		return ""
	}

	return string(v)
}

// 封装GetInt方法
func (self *PostData) GetInt(field string) int64 {
	if self.Type == 0 {
		value := (*self.Ctx).FormValue(field)
		if strings.Compare(value, "") == 0 {
			return 0
		}

		if num, err := strconv.ParseInt(value, 10, 64); err == nil {
			return num
		}
		return 0
	}

	v, _, _, err := jsonparser.Get(*self.Data, field)
	if err != nil {
		return 0
	}
	if num, numErr := strconv.Atoi(string(v)); numErr == nil {
		return int64(num)
	}

	return 0
}

// 获取key-value映射的map
func (self *PostData) GetMap() map[string]string {
	if self.Type == 1 { //json-body模式
		data := map[string]string{}
		tmp := map[string]interface{}{}
		json.Unmarshal(*self.Data, &tmp)

		for k, v := range tmp {
			value := ""
			switch reflect.TypeOf(v).String() {
			case "int32":
				value = strconv.Itoa(int(v.(int32)))
			case "int64":
				value = strconv.Itoa(int(v.(int64)))
			case "float32":
				value = decimal.NewFromFloat32(v.(float32)).String()
			case "float64":
				value = decimal.NewFromFloat(v.(float64)).String()
			case "bool":
				value = "true"
				if v.(bool) {
					value = "false"
				}
			case "string":
				value = v.(string)
			default:
				if strings.HasPrefix(reflect.TypeOf(v).String(), "map[") {
					bs, err := json.Marshal(v)
					if err == nil {
						value = string(bs)
					}
				}
			}
			data[k] = value
		}
		return data
	}

	//form模式
	values := (*self.Ctx).FormValues()
	data := map[string]string{}
	for k, v := range values {
		data[k] = strings.Join(v, ",")
	}

	return data
}
