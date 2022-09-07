package cmdb_handler

import (
	"fmt"
	"goframework/src/app/cmdb/model"
	"goframework/src/app/cmdb/onlineDDL"
	"goframework/src/framework/data/mysql_cli"
	"goframework/src/framework/response"
	"goframework/src/framework/result"
	"goframework/variable"
)

// CreateClass 资源类别创建接口
func CreateClass(mo *model.AdminCmdbClass) *response.JsonResult {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	// 创建资源类别记录
	db := tx.Table(mo.TableName()).Create(mo)
	result.Result(db.Error).Process(result.SqlCrudErr)

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.CreateAudit(tx, mo)).Process(result.SqlCrudErr)

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, db.RowsAffected, nil)
}

// CreateModel 资源模型创建接口
func CreateModel(mo *model.AdminCmdbModel) *response.JsonResult {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	// 创建模型记录(事务操作)
	mo.ModelName = fmt.Sprintf("%d_%s", mo.TenantId, mo.ModelName)
	db := tx.Table(mo.TableName()).Create(mo)
	result.Result(db.Error).Process(result.SqlCrudErr)

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.CreateAudit(tx, mo)).Process(result.SqlCrudErr)

	// 创建模型表(这里不用事务,DDL会隐式提交事务,就DDL如果在begin与rollback之间,会隐式提交,然后事务就嗝屁了...)
	result.Result(onlineDDL.Create(mysql_cli.GormDB, mo.ModelName)).Process(func(err error, data ...interface{}) {
		if err != nil {
			panic(result.NewConstErr(err.Error(), result.CreateErrCode, "DDL创建模型失败"))
		}
	})

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, db.RowsAffected, []struct{}{})
}

// CreateCheck 字段验证器创建接口
func CreateCheck(mo *model.AdminCmdbCheck) *response.JsonResult {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	db := tx.Table(mo.TableName()).Create(mo)
	result.Result(db.Error).Process(result.SqlCrudErr)

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.CreateAudit(tx, mo)).Process(result.SqlCrudErr)

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, db.RowsAffected, nil)
}

// CreateShowF 展示定制列创建接口
func CreateShowF(mo *model.AdminShowField, fields []string) *response.JsonResult {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	// 先删除原来的数据
	result.Result(tx.Table(mo.TableName()).Delete(mo)).Process(result.SqlCrudErr)

	// 构建定制列数据
	var showFs []*model.AdminShowField
	for i := 0; i < len(fields); i++ {
		showFs = append(showFs, &model.AdminShowField{
			UsersId:   mo.UsersId,
			ModelId:   mo.ModelId,
			TenantId:  mo.TenantId,
			FieldName: fields[i],
		})
	}

	// 创建定制列数据
	showFDb := tx.Table(mo.TableName()).Create(showFs)
	result.Result(showFDb.Error).Process(result.SqlCrudErr)

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.CreateAudit(tx, mo)).Process(result.SqlCrudErr)

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, showFDb.RowsAffected, nil)
}

// CreateField 资源模型字段创建接口
func CreateField(mos map[string][]*model.AdminCmdbField) *response.JsonResult {
	// 构建DDL字段结构体
	ddlFieldFunc := func(mo *model.AdminCmdbField, opt string) *onlineDDL.Field {
		return &onlineDDL.Field{
			Len:       mo.FieldLength,
			Name:      mo.FieldName,
			Type:      mo.FieldType,
			NotNull:   mo.AllowNotNull,
			Comment:   mo.Comments,
			Operation: opt,
		}
	}

	// 操作计数/操作结果
	var (
		Rows          int64
		returnResults = newDdlResults()
	)

	// 新增字段
	for i := 0; i < len(mos["create"]); i++ {
		// 构建成函数的原因是为把多字段操作构建成独立的一个个事务单元
		// 需要注意的是这里的defer操控着Rollback,数据的一致性就靠它
		createFunc := func() {
			// 事务操作
			tx := Transaction()
			defer tx.Rollback()

			// 为了简化代码
			var createReturnFunc = func(errComment string) {
				returnResults.Add(false, mos["create"][i].FieldName, "", "create", errComment)
				return
			}

			// 新增字段落库
			db := tx.Table(mos["create"][i].TableName()).Create(mos["create"][i])
			if db.Error != nil {
				createReturnFunc(fmt.Sprintf("字段落库失败,原因:%s", db.Error.Error()))
				return
			}

			// 审计落库
			err := model.AdminCmdbAudit{}.CreateAudit(tx, mos["create"][i])
			if err != nil {
				createReturnFunc(fmt.Sprintf("审计落库失败,原因:%s", err.Error()))
				return
			}

			// 查询模型名称
			var modelScan = &model.AdminCmdbModel{}
			err = tx.Model(&model.AdminCmdbModel{}).Where("id = ?", mos["create"][i].ModelId).First(modelScan).Error
			if err != nil {
				createReturnFunc(fmt.Sprintf("查询模型名称失败,原因:%s", err.Error()))
				return
			}

			// 进行DDL操作,开始创建字段
			field := ddlFieldFunc(mos["create"][i], onlineDDL.Add)
			// 这里不用事务,DDL会隐式提交事务,就DDL如果在begin与rollback之间,会隐式提交,然后事务就嗝屁了...
			if sql, err := onlineDDL.Alter(mysql_cli.GormDB, field, modelScan.ModelName); err != nil {
				returnResults.Add(false, field.Name, sql, "create", fmt.Sprintf("DDL操作失败,原因:%s", err.Error()))
				return
			}

			// Rows Count
			Rows = Rows + db.RowsAffected

			// 提交一次
			tx.Commit()
		}
		// Run
		createFunc()
	}

	// 更新字段
	for i := 0; i < len(mos["update"]); i++ {
		// 构建成函数的原因是为把多字段操作构建成独立的一个个事务单元
		// 需要注意的是这里的defer操控着Rollback,数据的一致性就靠它
		updateFunc := func() {
			// 事务操作
			tx := Transaction()
			defer tx.Rollback()

			// 为了简化代码
			var updateReturnFunc = func(errComment string) {
				returnResults.Add(false, mos["update"][i].FieldName, "", "update", errComment)
			}

			// 查询旧的字段名
			var oldFieldName string
			if db := tx.Model(mos["update"][i]).Where("id = ?", mos["update"][i].Id).Select("field_name").First(&oldFieldName); db.Error != nil {
				updateReturnFunc(fmt.Sprintf("查询旧的字段名失败,原因:%s", db.Error.Error()))
				return
			}

			// 更新字段记录落库
			db := tx.Model(mos["update"][i]).Where("id = ?", mos["update"][i].Id).Updates(mos["update"][i])
			if db.Error != nil {
				updateReturnFunc(fmt.Sprintf("落库失败,原因:%s", db.Error.Error()))
				return
			}

			// 审计记录
			err := model.AdminCmdbAudit{}.CreateAudit(tx, mos["update"][i])
			if err != nil {
				updateReturnFunc(fmt.Sprintf("审计落库失败,原因:%s", err.Error()))
				return
			}

			// 查询模型名称
			var modelScan = &model.AdminCmdbModel{}
			err = tx.Model(&model.AdminCmdbModel{}).Where("id = ?", mos["update"][i].ModelId).First(modelScan).Error
			if err != nil {
				updateReturnFunc(fmt.Sprintf("查询模型名称失败,原因:%s", err.Error()))
				return
			}

			// 进行DDL操作,开始更新字段
			field := ddlFieldFunc(mos["update"][i], onlineDDL.Change)
			// 这里不用事务,DDL会隐式提交事务,就DDL如果在begin与rollback之间,会隐式提交,然后事务就嗝屁了...
			sql, err := onlineDDL.Alter(mysql_cli.GormDB, field, modelScan.ModelName, oldFieldName)
			if err != nil {
				returnResults.Add(false, field.Name, sql, "update", fmt.Sprintf("DDL操作失败:%s", err.Error()))
				return
			}

			// Rows Count
			Rows = Rows + db.RowsAffected

			// 提交一次
			tx.Commit()
		}
		// Run
		updateFunc()
	}

	// 删除字段
	for i := 0; i < len(mos["delete"]); i++ {
		// 构建成函数的原因是为把多字段操作构建成独立的一个个事务单元
		// 需要注意的是这里的defer操控着Rollback,数据的一致性就靠它
		deleteFunc := func() {
			// 事务操作
			tx := Transaction()
			defer tx.Rollback()

			// 为了简化代码
			var deleteReturnFunc = func(errComment string) {
				returnResults.Add(false, mos["delete"][i].FieldName, "", "delete", errComment)
			}

			// 删除字段记录落库
			db := tx.Model(mos["delete"][i]).Where("id = ?", mos["delete"][i].Id).Delete(mos["delete"][i])
			if db.Error != nil {
				deleteReturnFunc(fmt.Sprintf("删除字段失败,原因:%s", db.Error.Error()))
				return
			}

			// 审计记录
			err := model.AdminCmdbAudit{}.CreateAudit(tx, mos["delete"][i])
			if err != nil {
				deleteReturnFunc(fmt.Sprintf("审计落库失败,原因:%s", err.Error()))
				return
			}

			// 查询模型名称
			var modelScan = &model.AdminCmdbModel{}
			err = tx.Model(&model.AdminCmdbModel{}).Where("id = ?", mos["delete"][i].ModelId).First(modelScan).Error
			if err != nil {
				deleteReturnFunc(fmt.Sprintf("查询模型名称失败,原因:%s", err.Error()))
				return
			}

			// 进行DDL操作,开始删除字段
			field := ddlFieldFunc(mos["delete"][i], onlineDDL.Drop)
			// 这里不用事务,DDL会隐式提交事务,就DDL如果在begin与rollback之间,会隐式提交,然后事务就嗝屁了...
			sql, err := onlineDDL.Alter(mysql_cli.GormDB, field, modelScan.ModelName)
			if err != nil {
				returnResults.Add(false, field.Name, sql, "delete", fmt.Sprintf("DDL操作失败,原因:%s", err.Error()))
				return
			}

			// Rows Count
			Rows = Rows + db.RowsAffected

			// 提交一次
			tx.Commit()
		}
		// Run
		deleteFunc()
	}

	return response.NewJsonResult(result.CreateErrCode, Rows, returnResults.Get())
}
