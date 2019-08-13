package utils

import (
	"encoding/json"
	"html"
	"qpgame/models"
	"strconv"
	"strings"
)

//批量更新相同属性，相同的值
func UpdateBatchSame(tableName string, params map[string]string, ids []int) string {
	sqlStr := "update " + tableName + " set "
	for k, v := range params {
		sqlStr += k + "='" + v + "',"
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	sqlStr += " where id in ("
	for _, id := range ids {
		sqlStr += strconv.Itoa(id) + ","
	}
	sqlStr = sqlStr[0:len(sqlStr)-1] + ")"
	return sqlStr
}

func QueryString(platform string, sqlString string) ([]map[string]interface{}, error) {
	res, err := models.MyEngine[platform].Query(sqlString)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	var response = make([]map[string]interface{}, 0)
	for i := range res {
		resMap := make(map[string]interface{})
		user := res[i]
		for k, v := range user {
			var temp interface{}
			json.Unmarshal(v, &temp)
			//所有字符串查询出来都为空，需要在这里判断
			temp = string(v)
			resMap[k] = temp
		}
		response = append(response, resMap)
	}

	return response, nil
}

func Query(platform string, sqlString string, stringArgs ...string) ([]map[string]interface{}, error) {
	res, err := models.MyEngine[platform].Query(sqlString)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	var response = make([]map[string]interface{}, 0)
	for i := range res {
		resMap := make(map[string]interface{})
		user := res[i]
		for k, v := range user {
			var temp interface{}
			json.Unmarshal(v, &temp)
			//所有字符串查询出来都为空，需要在这里判断
			if temp == nil {
				temp = string(v)
			}
			for _, s := range stringArgs {
				if s == k {
					temp = string(v)
				}
			}
			resMap[k] = temp
		}
		response = append(response, resMap)
	}

	return response, nil
}

func BuildQueryCondition(filterArr []map[string]string) string {
	whereStr := ""
	var whereArr []string
	if len(filterArr) > 0 {
		for _, filterItem := range filterArr {
			key, keyOk := filterItem["key"]
			val, valOk := filterItem["val"]
			prefix, prefixOk := filterItem["prefix"]
			if keyOk && valOk && val != "" {
				val = html.EscapeString(val)
				key = "`" + key + "`"
				if prefixOk {
					key = prefix + "." + key
				}
				if connType, ok := filterItem["condition"]; ok {
					switch connType {
					case "gt":
						whereArr = append(whereArr, key+">'"+val+"'")
					case "ge":
						whereArr = append(whereArr, key+">='"+val+"'")
					case "lt":
						whereArr = append(whereArr, key+"<'"+val+"'")
					case "le":
						whereArr = append(whereArr, key+"<='"+val+"'")
					case "like_a":
						whereArr = append(whereArr, key+" LIKE '%"+val+"'")
					case "like_b":
						whereArr = append(whereArr, key+" LIKE '"+val+"%'")
					case "like_ab":
						whereArr = append(whereArr, key+" LIKE '%"+val+"%'")
					case "in":
						whereArr = append(whereArr, key+" IN ("+val+")")
					}
				} else {
					whereArr = append(whereArr, key+"='"+val+"'")
				}
			} else {
				continue
			}
		}
		if len(whereArr) > 0 {
			whereStr = strings.Join(whereArr, " AND ")
		}
	}
	return whereStr
}

func BuildWhere(filterArr []map[string]string) string {
	whereStr := BuildQueryCondition(filterArr)
	if whereStr != "" {
		whereStr = " WHERE " + whereStr
	}
	return whereStr
}

func GetDefaultPageInfo(page, size int) (int, int) {
	if size <= 0 || size > 500 {
		size = 20
	}
	if page <= 0 {
		page = 1
	}
	return page, size
}

func BuildXormLimit(page, size int) (int, int) {
	if page < 1 || size <= 0 {
		page, size = GetDefaultPageInfo(page, size)
	}
	from := (page - 1) * size
	return size, from
}

func BuildLimit(page, size int) string {
	var from int
	size, from = BuildXormLimit(page, size)
	return " LIMIT " + strconv.Itoa(from) + "," + strconv.Itoa(size)
}
