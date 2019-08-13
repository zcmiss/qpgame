package ramcache

import (
	"qpgame/config"
	db "qpgame/models"
	"strings"

	"github.com/kataras/iris"
)

// 加载所有平台/数据库/表/字段信息
func LoadTableFields() {

	for platform := range config.PlatformCPs {
		conn, dbName := db.MyEngine[platform], ""
		sArr := strings.Split(conn.DataSourceName(), "/")
		if strings.Compare(sArr[1], "") == 0 {
			return
		}
		qArr := strings.Split(sArr[1], "?")
		if strings.Compare(qArr[0], "") == 0 {
			return
		}
		dbName = qArr[0]
		if strings.Compare(dbName, "") == 0 {
			return
		}
		sqlAll := "SELECT GROUP_CONCAT(COLUMN_NAME) fields,TABLE_NAME tbl FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA='" + dbName + "' GROUP BY tbl"
		rows, err := conn.SQL(sqlAll).QueryString()
		if err != nil {
			continue
		}
		for _, row := range rows {
			key := platform + "-" + dbName + "-" + row["tbl"]
			if _, ok := TableFields[key]; ok {
				delete(TableFields, key)
			}
			fields, allFields := strings.Split(row["fields"], ","), map[string]int{}
			for _, field := range fields {
				allFields[field] = 1
			}
			TableFields[key] = allFields
		}
		// 加载配置信息表, 前台有加载, 但是后台userFund.go会需要
		loadTableConfigs(platform, conn)
	}
}

//得到某个平台/数据库/表的的字段信息
func GetTableFields(ctx *iris.Context, tableName string) map[string]int {
	result := map[string]int{}
	platform := (*ctx).Params().Get("platform")
	conn := db.MyEngine[platform]
	dbName := ""
	{
		sArr := strings.Split(conn.DataSourceName(), "/")
		if strings.Compare(sArr[1], "") == 0 {
			return result
		}
		qArr := strings.Split(sArr[1], "?")
		if strings.Compare(qArr[0], "") == 0 {
			return result
		}
		dbName = qArr[0]
	}

	if strings.Compare(dbName, "") == 0 { //获取数据库名称失败
		return result
	}

	key := platform + "-" + dbName + "-" + tableName
	return TableFields[key]
}
