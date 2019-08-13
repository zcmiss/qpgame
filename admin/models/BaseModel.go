package models

import (
	"qpgame/admin/common"
	"qpgame/common/utils"
	"qpgame/config"
	db "qpgame/models"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 处理日期时间类型的字段
func processDatetime(fields *[]string, row *map[string]string) {
	for _, field := range *fields {
		value, exists := (*row)[field]
		if !exists {
			return
		}
		if strings.Compare(value, "") == 0 || strings.Compare(value, "0") == 0 {
			(*row)[field] = ""
		} else if valueTime, valueErr := strconv.ParseInt(value, 10, 64); valueErr == nil {
			(*row)[field] = time.Unix(valueTime, 0).Format("2006-01-02 15:04:05") //修改时间
		}
	}
}

// 依据可选项来处理字段显示
func processOptions(field string, options *map[string]string, row *map[string]string) {
	value, exists := (*row)[field]
	if !exists {
		return
	}

	option, hasOption := (*options)[value]
	if !hasOption {
		(*row)[field] = "未定义"
	} else {
		(*row)[field] = option
	}
}

// 依据可查询的字段, 来生成新的字段
func processOptionsFor(field string, fieldMap string, options *map[string]string, row *map[string]string) {
	value, exists := (*row)[field]
	if !exists {
		return
	}

	option, hasOption := (*options)[value]
	if !hasOption {
		(*row)[fieldMap] = "未定义"
	} else {
		(*row)[fieldMap] = option
	}
}

// 生成要查询的字段
// data: map[string]string { "字段名称": "操作符"}, 如: {"id": "="}
func getQueryFields(ctx *Context, fields *map[string]string) []string {
	var cond []string
	//fieldsArr := strings.Split(fields, ",")
	for field, op := range *fields {
		value := (*ctx).URLParam(strings.TrimSpace(field))
		if value != "" {
			if op == "%" { //如果是 like
				cond = append(cond, "`"+field+"` LIKE '%"+value+"%'")
			} else if op == "=" {
				cond = append(cond, "`"+field+"` = '"+value+"'")
			} else if op == "in" {
				values := "('" + strings.Join(strings.Split(value, ","), "','") + "')"
				cond = append(cond, "`"+field+"` IN "+values)
			} else {
				cond = append(cond, "`"+field+"` "+value)
			}
		}
	}
	return cond
}

// 得到查询的时间相关字段
func getQueryTime(ctx *Context, field string) int64 {
	value := (*ctx).URLParam(field)
	if value == "" {
		return 0
	}
	if !utils.IsDate(value) {
		return 0
	}

	return utils.GetInt64FromDate(value)
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	//if v, err := time.ParseInLocation("2006-01-02", value, loc); err == nil {
	//	currentTime = v.Unix()
	//}
}

// 依据多个时间相关字段, 来生成相应的sql
func getQueryFieldByTimes(ctx *Context, startField string, endField string) string {
	startTime := getQueryTime(ctx, startField)
	endTime := getQueryTime(ctx, endField)
	if startTime <= 0 && endTime <= 0 {
		return ""
	}
	if startTime <= 0 && endTime != 0 {
		return endField + " < " + strconv.FormatInt(endTime, 10)
	} else if startTime != 0 && endTime <= 0 {
		return startField + " >= " + strconv.FormatInt(startTime, 10)
	} else {
		return "(" + startField + " >= " + strconv.FormatInt(startTime, 10) + " AND " + endField + " < " + strconv.FormatInt(endTime+86400, 10) + ")"
	}
}

// 依据时间得到关于区间的时间查询
func getQueryFieldByTime(ctx *Context, field string, startField string, endField string) string {
	startTime := getQueryTime(ctx, startField)
	endTime := getQueryTime(ctx, endField)
	if startTime <= 0 && endTime <= 0 {
		return ""
	}
	if startTime <= 0 && endTime != 0 {
		return field + " < " + strconv.FormatInt(endTime, 10)
	} else if startTime != 0 && endTime <= 0 {
		return field + " >= " + strconv.FormatInt(startTime, 10)
	} else {
		return "(" + field + " BETWEEN " + strconv.FormatInt(startTime, 10) + " AND " + strconv.FormatInt(endTime+86400, 10) + ")"
	}
}

// 依据时间字段拿到查询信息
func getQueryFieldByDate(ctx *Context, field string, startField string, endField string) string {
	ymdStart := 0
	ymdEnd := 0
	dateStart := (*ctx).URLParam(startField)
	if utils.IsDate(dateStart) {
		dateStr := strings.Replace(dateStart, "-", "", 2)
		if v, err := strconv.Atoi(dateStr); err == nil {
			ymdStart = v
		}
	}
	dateEnd := (*ctx).URLParam(endField)
	if utils.IsDate(dateStart) {
		dateStr := strings.Replace(dateEnd, "-", "", 2)
		if v, err := strconv.Atoi(dateStr); err == nil {
			ymdEnd = v
		}
	}
	if ymdStart == 0 && ymdEnd == 0 {
		return ""
	}
	if ymdStart <= 0 && ymdEnd > 0 {
		return field + " < " + strconv.Itoa(ymdEnd)
	} else if ymdStart > 0 && ymdEnd <= 0 {
		return field + " >= " + strconv.Itoa(ymdStart)
	} else {
		return "(" + field + " BETWEEN " + strconv.Itoa(ymdStart) + " AND " + strconv.Itoa(ymdEnd) + ")"
	}
}

//将datetime转化为时间戳
func turnDatetimeFields(fields *[]string, data *map[string]string) {
	for _, field := range *fields {
		val, exists := (*data)[field]
		if !exists {
			continue
		}

		if !utils.IsDatetime(val) { //必须是日期时间格式
			continue
		}

		value := utils.GetInt64FromTime((*data)[field])
		(*data)[field] = strconv.FormatInt(value, 10)
	}
}

//将datetime转化为时间戳
func turnDateFields(fields *[]string, data *map[string]string) {
	for _, field := range *fields {
		val, exists := (*data)[field]
		if !exists {
			continue
		}

		if !utils.IsDate(val) { //必须是日期时间格式
			continue
		}

		value := utils.GetInt64FromDate((*data)[field])
		(*data)[field] = strconv.FormatInt(value, 10)
	}
}

// 得到关于user_id的查询自动生成sql语句
func getQueryOfUserId(ctx *Context) string {
	idStr := (*ctx).URLParam("user_id")
	if idStr != "" {
		return "(user_id IN (SELECT id FROM users WHERE id LIKE '%" + idStr + "%' OR user_name LIKE '%" + idStr + "%'))"
	}
	return ""
}

// 返回: (Pager, bool) => (分页, 执行是否成功)
// 注意: 如果正常执行，但是没有数据，返回也是成功
func getRecords(ctx *Context,
	model IAdminModel,
	fields string,
	getQueryCond func(ctx *Context) []string, //查询条件
	processRecord func(ctx *Context, row *map[string]string),
	//返回值: (orderBy, groupBy, limit)
	getOrderBy func(ctx *Context) (string, string, int)) (Pager, error) { //对于每个记录的处理

	tableName := model.GetTableName(ctx)
	sql := "SELECT " + fields + " FROM " + tableName          //生成sql语句
	where := ""                                               //" WHERE id > " + strconv.Itoa(lastId) + " " //where查询条件
	isUnionAll := strings.Index(tableName, " UNION ALL ") > 0 //是否是联合查询
	if !isUnionAll {
		if getQueryCond != nil { //对于查询条件进行利累加
			queryConditions := getQueryCond(ctx)
			for _, v := range queryConditions {
				if v != "" { //必须不能为空
					where += " AND " + v //对查询语句进行累积
				}
			}
		}
		idQuery := getQueryOfUserId(ctx)
		if idQuery != "" {
			where += " AND " + idQuery
		}
	}

	page := (*ctx).URLParamIntDefault("page", 1)
	limit := (*ctx).URLParamIntDefault("page_size", 20)
	groupBy := ""
	orderBy := " ORDER BY id DESC " //ORDER BY id DESC" //排序字段

	if getOrderBy != nil {
		limitNum := 0 //每页记录数
		groupByStr, orderByStr, limitNum := getOrderBy(ctx)
		if groupByStr != "" {
			groupBy = " GROUP BY " + groupByStr
		}
		if orderByStr != "" {
			orderBy = " ORDER BY " + orderByStr
		}
		if limitNum > 0 { //小于等于0表示对limit不做限制使用get来的默认结果
			limit = limitNum
		}
	}

	if limit > 500 { //最大允许500条记录
		limit = 500
	}
	sql += " WHERE 1 = 1 " + where + " " + groupBy + " " + orderBy + " LIMIT " + strconv.Itoa(limit*(page-1)) + ", " + strconv.Itoa(limit) //拼装sql语句
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	rows, err := conn.SQL(sql).QueryString()
	if err != nil {
		return Pager{}, Error{What: "执行查询失败"}
	}

	rowsCount := len(rows)
	if rowsCount < 1 {
		return newPager(), nil
	}

	if processRecord != nil { //如果要对于每条记录进行处理
		for _, r := range rows {
			processRecord(ctx, &r)
			processDatetime(&[]string{"created", "updated"}, &r)
		}
	} else {
		for _, r := range rows {
			processDatetime(&[]string{"created", "updated"}, &r)
		}
	}

	sql = "SELECT COUNT(*) AS total FROM " + tableName + " WHERE 1 = 1 " + where + " " + groupBy
	countRows, err := conn.SQL(sql).QueryString()
	if err != nil || len(countRows) == 0 {
		return Pager{}, Error{What: "查询统计信息失败"}
	}
	totalCount := len(countRows)
	if groupBy == "" {
		totalCount, err = strconv.Atoi(countRows[0]["total"])
		if err != nil {
			return Pager{}, Error{What: "获取统计信息失败"}
		}
	}
	pageCount := 0 //总的页数
	if totalCount > 0 {
		num := totalCount % limit
		if num == 0 {
			pageCount = totalCount / limit
		} else {
			pageCount = totalCount/limit + 1
		}
	}

	pager := Pager{
		PageCount: pageCount,  //总页数
		Page:      page,       //当前页数
		TotalRows: totalCount, //总记录数
		PageSize:  limit,      //每页记录数
		Rows:      rows,       //总记录
	}

	return pager, nil
}

// 保存数据
func saveRecord(ctx *Context,
	model IAdminModel,
	beforeSave func(ctx *Context, data *map[string]string) bool,
	beforeCreate func(ctx *Context, data *map[string]string) bool,
	beforeUpdate func(ctx *Context, data *map[string]string) bool,
	afterSave func(ctx *Context, data *map[string]string, isCreating bool)) (int64, error) {

	postData := utils.GetPostData(ctx)
	idStr := postData.Get("id") //默认传递的参数是id
	id, err := strconv.Atoi(idStr)
	isCreating := false //是否是正在创建数据
	if err != nil || id == 0 {
		isCreating = true
	}

	data := postData.GetMap() //make(map[string]string)
	if beforeSave != nil && !beforeSave(ctx, &data) {
		return 0, Error{What: "保存前数据校验失败"}
	}

	var affected int64
	if isCreating { // 如果是创建记录
		if beforeCreate != nil && !beforeCreate(ctx, &data) { //如果有前置操作, 并且前置操作执行不成功
			return 0, Error{What: "添加数据前校验失败"}
		}
		if _, ok := data["id"]; ok {
			delete(data, "id")
		}
		affected, err = createRecord(ctx, &data, model.GetTableName(ctx))
	} else { //如果是修改记录
		conditions := " id = " + idStr
		if beforeUpdate != nil && !beforeUpdate(ctx, &data) {
			return 0, Error{What: "修改数据前校验失败"}
		}
		if _, ok := data["id"]; ok {
			delete(data, "id")
		}
		affected, err = updateRecord(ctx, &data, model.GetTableName(ctx), conditions)
	}
	if err != nil {
		return affected, err
	}
	if afterSave != nil {
		afterSave(ctx, &data, isCreating)
	}
	return affected, nil
}

// 更新记录
func updateRecord(ctx *Context,
	data *map[string]string,
	tableName string,
	conditions string) (int64, error) {

	sql := "UPDATE " + tableName + " SET "
	index := 0
	delete(*data, "id")
	realFields := ramcache.GetTableFields(ctx, tableName)
	_, hasUpdated := realFields["updated"]
	if hasUpdated {
		(*data)["updated"] = strconv.Itoa(utils.GetNowTime())
	}
	for k, v := range *data {
		realField, exists := realFields[k] //获取现有的表的所有字段
		if !exists || realField != 1 {     //跳过所有在数据表中不存在的字段
			continue
		}
		if index == 0 {
			sql += "`" + k + "` = '" + v + "'"
		} else {
			sql += ", `" + k + "` = '" + v + "'"
		}
		index += 1
	}
	if conditions != "" {
		sql += " WHERE " + conditions
	}

	result, err := db.MyEngine[(*ctx).Params().Get("platform")].Exec(sql)
	if err != nil {
		return 0, Error{What: "执行修改操作失败"}
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, Error{What: "数据未经修改"}
	}
	return affectedRows, nil
}

// 添加记录
func createRecord(ctx *Context, data *map[string]string, tableName string) (int64, error) {
	fields := ""
	values := ""
	sql := "INSERT INTO " + tableName + " ("
	index := 0
	realFields := ramcache.GetTableFields(ctx, tableName) //获取现有的表的所有字段
	_, hasCreated := realFields["created"]
	if hasCreated {
		(*data)["created"] = strconv.Itoa(utils.GetNowTime())
	}
	_, hasUpdated := realFields["updated"]
	if hasUpdated {
		(*data)["updated"] = strconv.Itoa(utils.GetNowTime())
	}
	for k, v := range *data {
		realField, exists := realFields[k]
		if !exists || realField != 1 { //跳过所有在数据表中不存在的字段
			continue
		}
		if index == 0 {
			fields += "`" + k + "`"
			values += "'" + v + "'"
		} else {
			fields += ", `" + k + "`"
			values += ", '" + v + "'"
		}
		index += 1
	}
	sql += fields + ") VALUES (" + values + ")"
	result, err := db.MyEngine[(*ctx).Params().Get("platform")].Exec(sql)
	if err != nil { //判断是否执行成功
		return 0, Error{What: "执行添加操作失败"}
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, Error{What: "数据未添加成功"}
	}
	return insertId, nil
}

// 是否是删除操作
func isDelete(ctx *Context) bool {
	path := (*ctx).Path()
	return strings.LastIndex(path, "/delete") > 0
}

// 删除记录
func deleteRecord(ctx *Context,
	model IAdminModel,
	beforeDelete func(ctx *Context, ids string) bool,
	afterDelete func(ctx *Context, ids string)) error {

	idStr := (*ctx).URLParam("id")
	if strings.Compare(idStr, "") == 0 {
		return Error{What: "无法获取ID"}
	}
	idAll := strings.Split(idStr, ",")
	var ids []int
	idCount := 0
	for _, id := range idAll {
		idReal, idErr := strconv.Atoi(id)
		if idErr == nil {
			ids = append(ids, idReal)
			idCount += 1
		}
	}
	if idCount == 0 { //表示没有一个合法的数字
		return Error{What: "要删除的数据ID不能为空"}
	}

	var idArr []string
	for _, id := range ids {
		idArr = append(idArr, strconv.Itoa(id))
	}
	idsAll := strings.Join(idArr, ",")
	if beforeDelete != nil && !beforeDelete(ctx, idsAll) {
		return Error{What: "删除前校验失败"}
	}

	sql := "DELETE FROM " + model.GetTableName(ctx) + " WHERE id IN (" + idsAll + ")"
	result, err := db.MyEngine[(*ctx).Params().Get("platform")].Exec(sql)
	if err != nil { //当中是否有错误
		return Error{What: "执行删除操作时出现错误"}
	}

	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows <= 0 {
		return Error{What: "未知错误, 删除失败"}
	}

	if afterDelete != nil {
		afterDelete(ctx, idsAll)
	}
	return nil
}

// 得到单条记录
func getRecordDetail(ctx *Context,
	model IAdminModel,
	fields string,
	processRecord func(ctx *Context, record *map[string]string)) (map[string]string, error) {

	id, err := (*ctx).URLParamInt("id")
	if err != nil {
		return map[string]string{}, nil
	}

	columns := "*"
	if fields != "" {
		columns = fields
	}

	sql := "SELECT " + columns + " FROM " + model.GetTableName(ctx) + " WHERE id = " + strconv.Itoa(id)
	rows, err := db.MyEngine[(*ctx).Params().Get("platform")].SQL(sql).QueryString()
	if err != nil || len(rows) < 1 {
		return map[string]string{}, nil
	}

	row := rows[0]
	if processRecord != nil {
		processRecord(ctx, &row)
	}
	processDatetime(&[]string{"created", "updated"}, &row)

	return row, nil
}

// 拒绝执行getRecordDetail()操作
func denyDetail() (map[string]string, error) {
	return nil, Error{What: "拒绝获取详情操作,仅用于接口"}
}

// 拒绝执选择save()操作
func denySave() (int64, error) {
	return 0, Error{What: "拒绝保存数据操作,仅用于接口"}
}

// 拒绝删除数据操作
func denyDelete() error {
	return Error{What: "拒绝删除数据操作,仅用于接口"}
}

//写后台的操作日志, 从 (*ctx).Params().Get("log") 当中获取 ,格式: 日志类型|日志内容
func WriteAdminLog(ctx *Context) {
	logStr := (*ctx).Params().Get("log")
	if logStr == "" {
		return
	}
	logArr := strings.Split(logStr, "|")
	if len(logArr) < 2 {
		return
	}

	//1. 解析出管理员编号
	platform := (*ctx).Params().Get("platform")     //平台标识
	authString := (*ctx).GetHeader("Authorization") //提交過來的授權字符串
	tokenString := strings.Replace(authString, "bearer ", "", 7)
	requestToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenKey), nil
	})
	if err != nil { //解析的token不正确
		return
	}
	claim, _ := requestToken.Claims.(jwt.MapClaims) //解码
	adminIdStr := claim["id"].(string)

	key := platform + "-" + adminIdStr
	admin, exists := common.Admins[key]
	if !exists {
		return
	}

	//2. 解析出菜单url
	menus, menuExists := common.AdminMenus[platform]
	if !menuExists || len(menus) == 0 {
		return
	}

	menuName := ""
	pathArr := strings.Split((*ctx).Path(), "/v1/")
	if len(pathArr) < 2 {
		return
	}
	path := "/" + pathArr[1]
	for _, menu := range menus {
		if menu.Url == path {
			menuName = menu.Title
			break
		}
	}
	if menuName == "" { //没有找到相对应的路由
		return
	}

	logType := logArr[0]    //日志类型
	logContent := logArr[1] //日志内容
	//3. 写日志
	data := map[string]string{
		"admin_id":   strconv.Itoa(admin.Id),           //管理员编号
		"admin_name": admin.Name,                       //管理员名称
		"type":       logType,                          //类型
		"node":       menuName,                         //菜单名称/路由名称
		"content":    logContent,                       //操作日志内容
		"created":    strconv.Itoa(utils.GetNowTime()), //日志创建时间
	}
	affected, err := createRecord(ctx, &data, "admin_logs")
	if err != nil || affected <= 0 { //可能要写的其他的额外的日志
	}
}

// 默认的用于保存日志的保存后回调函数
func getSavedFunc(name string, key string) func(ctx *Context, data *map[string]string, isCreating bool) {
	return func(ctx *Context, data *map[string]string, isCreating bool) {
		if isCreating {
			(*ctx).Params().Set("log", "添加|添加"+name+": "+(*data)[key])
		} else {
			post := utils.GetPostData(ctx)
			(*ctx).Params().Set("log", "修改|修改(id:"+post.Get("id")+")"+name+": "+(*data)[key])
		}
	}
}

//默认的用于删除记录后的写日期志的回调函数
func getDeletedFunc(name string) func(ctx *Context, ids string) {
	return func(ctx *Context, ids string) {
		idStr := (*ctx).URLParam("id")
		if idStr == "" {
			return
		}
		(*ctx).Params().Set("log", "删除|删除编号为"+idStr+"的"+name)
	}
}
