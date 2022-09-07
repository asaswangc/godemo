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

// DeleteClass 资源类别删除接口
func DeleteClass(id int) interface{} {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	// 校验资源类是否存在
	classScan := &model.AdminCmdbClass{}
	result.Result(tx.Model(&model.AdminCmdbClass{}).Where("id = ?", id).First(&classScan).Error).Process(result.SqlCrudErr)
	if classScan == nil {
		panic(result.NewConstErr(fmt.Sprintf("资源类别%d不存在", id), result.DeleteErrCode))
	}

	// 删除前需要校验(校验是否被关联, 资源类别 ——> 资源模型)
	modelCount := int64(0)
	result.Result(tx.Model(&model.AdminCmdbModel{}).Where("class_id = ?", id).Count(&modelCount).Error).Process(result.SqlCrudErr)
	if modelCount != 0 {
		panic(result.NewConstErr(fmt.Sprintf("资源类别%d被使用,不可删除", id), result.DeleteErrCode))
	}

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.DeleteAudit(tx, classScan)).Process(result.SqlCrudErr)

	// 删除数据
	classDb := tx.Table(classScan.TableName()).Delete(&model.AdminCmdbClass{Id: id})
	result.Result(classDb.Error).Process(result.SqlCrudErr)

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, classDb.RowsAffected, nil)
}

// DeleteModel 资源模型删除接口
func DeleteModel(id int, tenantId int) interface{} {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	// 校验模型是否存在
	modelScan := &model.AdminCmdbModel{}
	result.Result(tx.Model(&model.AdminCmdbModel{}).Where("id = ?", id).First(&modelScan).Error).Process(result.SqlCrudErr)
	if modelScan == nil {
		panic(result.NewConstErr(fmt.Sprintf("资源模型%d不存在", id), result.DeleteErrCode))
	}

	// 删除前需要校验(校验模型下是否有数据,如果有数据不可删)
	dataCount := int64(0)
	result.Result(tx.Table(fmt.Sprintf("%d_%s", tenantId, modelScan.ModelName)).Count(&dataCount)).Process(result.SqlCrudErr)
	if dataCount != 0 {
		panic(result.NewConstErr(fmt.Sprintf("资源模型%d有数据,不可删除", id), result.DeleteErrCode))
	}

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.DeleteAudit(tx, modelScan)).Process(result.SqlCrudErr)

	// 删除数据(表记录)
	modelDb := tx.Table(modelScan.TableName()).Delete(&model.AdminCmdbModel{Id: id})
	result.Result(modelDb.Error).Process(result.SqlCrudErr)

	// 删除实体表(DDL操作)
	// 这里不用事务,DDL会隐式提交事务,就DDL如果在begin与rollback之间,会隐式提交,然后事务就嗝屁了...
	result.Result(onlineDDL.DropTable(mysql_cli.GormDB, modelScan.ModelName)).Process(func(err error, data ...interface{}) {
		if err != nil {
			panic(result.NewConstErr(err.Error(), result.DeleteErrCode, "删除模型失败"))
		}
	})

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, modelDb.RowsAffected, []struct{}{})
}

// DeleteCheck 字段验证器删除接口
func DeleteCheck(id int) interface{} {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	// 校验字段验证器是否存在
	checkScan := &model.AdminCmdbCheck{}
	result.Result(tx.Model(&model.AdminCmdbCheck{}).Where("id = ?", id).First(&checkScan).Error).Process(result.SqlCrudErr)
	if checkScan == nil {
		panic(result.NewConstErr(fmt.Sprintf("字段验证器%d不存在", id), result.DeleteErrCode))
	}

	// 删除前需要校验(校验是否被关联, 资源验证器 ——> 模型字段)
	fieldCount := int64(0)
	result.Result(tx.Model(&model.AdminCmdbField{}).Where("check_id = ?", id).Count(&fieldCount).Error).Process(result.SqlCrudErr)
	if fieldCount != 0 {
		panic(result.NewConstErr(fmt.Sprintf("字段验证器%d被使用,不可删除", id), result.DeleteErrCode))
	}

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.DeleteAudit(tx, checkScan)).Process(result.SqlCrudErr)

	// 删除数据
	checkDb := tx.Table(checkScan.TableName()).Delete(&model.AdminCmdbCheck{Id: id})
	result.Result(checkDb.Error).Process(result.SqlCrudErr)

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, checkDb.RowsAffected, nil)
}
