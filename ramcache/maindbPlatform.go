package ramcache

import (
	"encoding/json"
	"fmt"
	"qpgame/models"
	"qpgame/models/mainxorm"
)

//主库platform表
func mainDbPlatform(data []byte) {
	platformTable := make([]mainxorm.Platform, 0)
	val, _ := MainTablePlatform.Load("platform")
	platformTable = append(platformTable, val.([]mainxorm.Platform)...)
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	oper := rowMap["tableOperation"]
	//获取数据
	rowData := rowMap["data"].(map[string]interface{})
	id := int(rowData["id"].(float64))
	//新数据
	var ele = mainxorm.Platform{}
	ele.Id = id
	ele.Code = rowData["code"].(string)
	ele.Name = rowData["name"].(string)
	ele.AdminAddress = rowData["admin_address"].(string)
	ele.ApiAddress = rowData["api_address"].(string)
	ele.PicAddress = rowData["pic_address"].(string)
	ele.PayBackAddress = rowData["pay_back_address"].(string)

	ele.SilverMerchantAddress = rowData["silver_merchant_address"].(string)
	ele.SilverMerchantApiAddress = rowData["silver_merchant_api_address"].(string)
	//新增操作
	if oper == "insert" {
		platformTable = append(platformTable, ele)
		MainTablePlatform.Store("platform", platformTable)
		return
	}

	updateRowTable := make([]mainxorm.Platform, 0)
	//删除操作
	if oper == "delete" {
		for _, v := range platformTable {
			if v.Id != id {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	//更新操作
	if oper == "update" {
		for _, v := range platformTable {
			if v.Id == id {
				updateRowTable = append(updateRowTable, ele)
			} else {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	MainTablePlatform.Store("platform", updateRowTable)
}

//加载主库platform表
func loadMainDbTablePlatform() bool {
	engine := models.MyEngineMainDb
	platform := make([]mainxorm.Platform, 0)
	errRes := engine.Find(&platform)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	//这里的下标是表名
	MainTablePlatform.Store("platform", platform)
	return true
}

func update() {

}
