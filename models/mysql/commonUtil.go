package mysql

import (
	"database/sql"
	"qpgame/common/log"
	"strconv"
)

//mysql用到的一些共用处理放入这里,
//比如遍历等其他

//获取结果集
func GetMysqlResultSets(rows *sql.Rows) (container []map[string]interface{}) {
	log.DeferRecover()
	columnTypes, _ := rows.ColumnTypes()
	container = make([]map[string]interface{}, 0)
	//values是每个列的值，这里获取到byte里
	var values = make([]sql.RawBytes, len(columnTypes))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	var scans = make([]interface{}, len(columnTypes))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}
	for rows.Next() { //循环，让游标往下推
		if err := rows.Scan(scans...); err != nil {
			//query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			rows.Close()
			panic(err)
		}
		var parseRow = make(map[string]interface{})
		for k, val := range values { //每行数据是放在values里面，现在把它挪到row里
			fieldObject := *columnTypes[k]
			columnName := fieldObject.Name()
			columnType := fieldObject.DatabaseTypeName()
			var parseValue interface{}
			valString := string(val)
			switch columnType {
			case "INT":
				parseValue, _ = strconv.Atoi(valString)
			case "VARCHAR":
				parseValue = valString
			case "FLOAT":
				parseValue, _ = strconv.ParseFloat(valString, 64)
			case "CHAR":
				parseValue = valString
			case "TINYINT":
				parseValue, _ = strconv.Atoi(valString)
			case "MEDIUMINT":
				parseValue, _ = strconv.Atoi(valString)
			case "DECIMAL":
				parseValue, _ = strconv.ParseFloat(valString, 64)
			default:
				parseValue = valString
			}
			parseRow[columnName] = parseValue
		}
		container = append(container, parseRow)
	}
	return container
}
