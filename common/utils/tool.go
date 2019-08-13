package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//检测int数组中是否存在某个元素
func InArrayInt(ele int, arr []int) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}

//检测字符串数组中是否存在某个元素
func InArrayString(ele string, arr []string) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}

//php的range函数
func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

//int数组合并成字符串
func IntJoin(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return fmt.Sprintf("%v", a[0])
	}

	buffer := &bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("%v", a[0]))
	for i := 1; i < len(a); i++ {
		buffer.WriteString(sep)
		buffer.WriteString(fmt.Sprintf("%v", a[i]))
	}
	return buffer.String()
}

// NumberFormat number_format()
// decimals: Sets the number of decimal points.
// decPoint: Sets the separator for the decimal point.
// thousandsSep: Sets the thousands separator.
func NumberFormat(number float64, decimals uint, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

// 多个数组进行合并, ArrayMerge array_merge()
func ArrayMergeString(ss ...[]string) []string {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]string, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

// Trim trim()
func Php2GoTrim(str string, characterMask ...string) string {
	mask := ""
	if len(characterMask) == 0 {
		mask = " \\t\\n\\r\\0\\x0B"
	} else {
		mask = characterMask[0]
	}
	return strings.Trim(str, mask)
}

// Md5 md5()
func Php2GoMd5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

//创建订单
func CreationOrder(prefix string, userIdS string) string {
	//生成订单
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(999999-100000) + 100000
	str := strconv.Itoa(randNum) + userIdS
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	str = string([]byte(md5str)[:6])
	orderId := prefix + time.Now().Format("20060102150405") + str
	return orderId
}

// Sha1 sha1()
func StrSha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
