package utils

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"reflect"
	"strconv"
	"strings"
)

type SeqMap struct {
	keys   []string
	seqMap map[string]interface{}
	err    string
}

func NewSeqMap() SeqMap {
	var seq SeqMap
	seq.keys = make([]string, 0)
	seq.seqMap = make(map[string]interface{})
	return seq
}

// 通过map主键唯一的特性过滤重复元素
func (seq *SeqMap) removeRepByMap() {
	result := []string{}                    // 存放结果
	tempMap := make(map[string]interface{}) // 存放不重复主键
	for i := len(seq.keys) - 1; i >= 0; i-- {
		e := seq.keys[i]
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	seq.keys = make([]string, 0) // 存放结果
	for i := len(result) - 1; i >= 0; i-- {
		seq.keys = append(seq.keys, result[i])
	}
}
func (seq *SeqMap) Put(key string, value interface{}) {
	seq.keys = append(seq.keys, key)
	seq.seqMap[key] = value
}

func (seq *SeqMap) PutOldIndex(key string, value interface{}) {
	_, ok := seq.seqMap[key]
	if !ok {
		seq.keys = append(seq.keys, key)
	}
	seq.seqMap[key] = value
}

func (seq *SeqMap) Get(key string) interface{} {
	return seq.seqMap[key]
}

func (seq *SeqMap) GetString(key string) string {
	fmt.Println(reflect.TypeOf(seq.seqMap[key]))
	if reflect.TypeOf(seq.seqMap[key]).Name() == "string" {
		return seq.seqMap[key].(string)
	}
	if reflect.TypeOf(seq.seqMap[key]).Name() == "int" {
		return strconv.Itoa(seq.seqMap[key].(int))
	}
	if reflect.TypeOf(seq.seqMap[key]).Name() == "float64" {
		value := fmt.Sprintf("%f", seq.seqMap[key].(float64))
		return value
	}
	if reflect.TypeOf(seq.seqMap[key]).Name() == "SeqMap" {
		return seq.seqMap[key].(*SeqMap).JsonSeq()
	}
	if reflect.TypeOf(seq.seqMap[key]).Name() == "bool" {

		return strconv.FormatBool(seq.seqMap[key].(bool))
	}
	seq.err = "type error!"
	return ""
}

func (seq *SeqMap) GetInt(key string) (int, error) {
	reflect.TypeOf(seq.seqMap[key])
	if reflect.TypeOf(seq.seqMap[key]).Name() == "float64" {
		value := int(seq.seqMap[key].(float64))
		return value, nil
	}
	if reflect.TypeOf(seq.seqMap[key]).Name() == "int" {
		return seq.seqMap[key].(int), nil
	}
	return 0, errors.New("type error!")
}

func (seq *SeqMap) GetFloat(key string) (float64, error) {
	reflect.TypeOf(seq.seqMap[key])
	if reflect.TypeOf(seq.seqMap[key]).Name() == "float64" {
		value := seq.seqMap[key].(float64)
		return value, nil
	}
	if reflect.TypeOf(seq.seqMap[key]).Name() == "int" {
		return float64(seq.seqMap[key].(int)), nil
	}
	return 0, errors.New("type error!")
}

func (seq *SeqMap) GetMap(key string) (SeqMap, error) {
	reflect.TypeOf(seq.seqMap[key])
	if reflect.TypeOf(seq.seqMap[key]).Name() == "SeqMap" {
		return seq.seqMap[key].(SeqMap), nil
	} else {
		return NewSeqMap(), errors.New("the value is not a SeqMap!")
	}
}

func (seq *SeqMap) Remove(key string) {
	for i, k := range seq.keys {
		if k == key {
			seq.keys = append(seq.keys[:i], seq.keys[i+1:]...)
			break
		}
	}
	delete(seq.seqMap, key)
}

func (seq *SeqMap) Keys() []string {
	seq.removeRepByMap()
	return seq.keys
}

func (seq *SeqMap) Values() map[string]interface{} {
	return seq.seqMap
}

func (seq *SeqMap) JsonSeq() string {
	seq.removeRepByMap()
	resStr := "{"
	for _, key := range seq.keys {
		var sumMap SeqMap
		if reflect.TypeOf(seq.seqMap[key]).Name() == "SeqMap" {
			sumMap = seq.seqMap[key].(SeqMap)
			str := (&sumMap).JsonSeq()
			resStr += fmt.Sprintf("\"%v\":%v,", key, str)
		} else {
			jsonstr, err := json.Marshal(seq.seqMap[key])
			if err != nil {
				fmt.Println(err.Error())
			}
			resStr += fmt.Sprintf("\"%v\":%v,", key, string(jsonstr))
		}
	}
	resStr = resStr[0 : len(resStr)-1]
	resStr += "}"
	return strings.Replace(strings.Replace(strings.Replace(resStr, "\"{", "{", -1), "}\"", "}", -1), "\\", "", -1)
}
