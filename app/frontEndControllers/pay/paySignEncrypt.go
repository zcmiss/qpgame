package pay

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"qpgame/common/utils"
	"sort"
	"strconv"
)

/**
 * payKsortAddstr
 * map排序，先遍历map Key至数组，排序。在遍历数组，进行字符串拼接
 * @param 			postData map[string]string 	map数据
 * @isConnector    bool  							判断是否需要添加&符号
 * return string
 */
func payKsortAddstr(postData map[string]string, isConnector int) string {
	sortedKeys := make([]string, 0)
	for k := range postData {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	//排序后Json处理返回
	if isConnector == 5 {
		addressesJson, _ := json.Marshal(postData)
		return string(addressesJson)
	}
	var str string
	keyLength := len(sortedKeys)
	for index, keyVal := range sortedKeys {
		if isConnector == 0 { // 原 False
			str += keyVal + "=" + postData[keyVal]
		} else if isConnector == 1 { // 原 True
			if keyVal != "sign" && postData[keyVal] != "" {
				str += keyVal + "=" + postData[keyVal]
			}
			if index != keyLength-1 {
				str += "&"
			}
		} else if isConnector == 2 { //如果值为空，则不参加签名，最后一位元素 不添加 &
			if postData[keyVal] != "" {
				str += keyVal + "=" + postData[keyVal] + "&"
			}
		} else if isConnector == 3 {
			if postData[keyVal] != "" {
				str += postData[keyVal]
				if index != keyLength-1 {
					str = str + "|"
				}
			}
		} else if isConnector == 4 {
			if postData[keyVal] != "" {
				str += postData[keyVal]
			}
		} else if isConnector == 5 {
			if postData[keyVal] != "" {
				str += postData[keyVal]
				if index != keyLength-1 {
					str = str + "#"
				}
			}
		} else if isConnector == 6 {
			str += utils.Php2GoTrim(keyVal) + utils.Php2GoTrim(postData[keyVal])
		} else if isConnector == 7 {
			if postData[keyVal] != "" {
				str += postData[keyVal]
				if index != keyLength-1 {
					str = str + ""
				}
			}
		} else if isConnector == 8 {
			str += keyVal + "=" + postData[keyVal] + "&"
		}
	}
	return str
}

/**
 * floatToString
 * 浮点数转换字符串 ， 主要用户付款金额转换
 * @param payAmout float64
 * return string
 */
func floatToString(payAmout float64) string {
	return strconv.FormatFloat(payAmout, 'f', -1, 64)
}

/**
 * hmacMd5
 * hmacMd5 加密
 * @param 	message  string	需要加密的字符串
 * @param 	key      string 加密字符串的key
 * return string
 */
func hmacMd5(message string, key string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func payKsortMapInterface(postData map[string]interface{}, isConnector int) string {
	sortedKeys := make([]string, 0)
	for k := range postData {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	//排序后Json处理返回
	var str string
	keyLength := len(sortedKeys)
	for index, keyVal := range sortedKeys {
		if isConnector == 103 {
			if keyVal != "sign" && keyVal != "payUrl" && keyVal != "payHTML" && postData[keyVal] != "" {
				str += keyVal + "=" + postData[keyVal].(string) + "&"
			}
			if index != keyLength-1 {

			}
		}
	}
	return str
}
