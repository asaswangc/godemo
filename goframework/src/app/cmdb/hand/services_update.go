package cmdb_handler

import (
	"fmt"
	"goframework/src/app/cmdb/model"
	"goframework/src/framework/response"
	"goframework/src/framework/result"
	"goframework/variable"
)

// UpdateClass 资源类别更新接口
func UpdateClass(mo *model.AdminCmdbClass) interface{} {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	// 校验资源类是否存在
	classScan := &model.AdminCmdbClass{}
	result.Result(tx.Model(&model.AdminCmdbClass{}).Where("id = ?", mo.Id).First(&classScan).Error).Process(result.SqlCrudErr)
	if classScan == nil {
		panic(result.NewConstErr(fmt.Sprintf("资源类别%d不存在", mo.Id), result.DeleteErrCode))
	}

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.UpdateAudit(tx, mo, classScan)).Process(result.SqlCrudErr)

	// 更新数据
	classDb := tx.Table(classScan.TableName()).Updates(mo)
	result.Result(classDb.Error).Process(result.SqlCrudErr)

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, classDb.RowsAffected, nil)
}

// UpdateModel 资源模型更新接口
func UpdateModel(mo *model.AdminCmdbModel) interface{} {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	// 校验资源模型是否存在
	modelScan := &model.AdminCmdbModel{}
	result.Result(tx.Model(&model.AdminCmdbModel{}).Where("id = ?", mo.Id).First(&modelScan).Error).Process(result.SqlCrudErr)
	if modelScan == nil {
		panic(result.NewConstErr(fmt.Sprintf("资源模型%d不存在", mo.Id), result.DeleteErrCode))
	}

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.UpdateAudit(tx, mo, modelScan)).Process(result.SqlCrudErr)

	// 更新数据
	classDb := tx.Table(modelScan.TableName()).Updates(mo)
	result.Result(classDb.Error).Process(result.SqlCrudErr)

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, classDb.RowsAffected, nil)
}

// UpdateCheck 字段验证器更新接口
func UpdateCheck(mo *model.AdminCmdbCheck) interface{} {
	// 事务操作
	tx := Transaction()
	defer tx.Rollback()

	// 校验资源模型是否存在
	checkScan := &model.AdminCmdbCheck{}
	result.Result(tx.Model(&model.AdminCmdbCheck{}).Where("id = ?", mo.Id).First(&checkScan).Error).Process(result.SqlCrudErr)
	if checkScan == nil {
		panic(result.NewConstErr(fmt.Sprintf("字段验证器%d不存在", mo.Id), result.DeleteErrCode))
	}

	// 审计记录
	result.Result(model.AdminCmdbAudit{}.UpdateAudit(tx, mo, checkScan)).Process(result.SqlCrudErr)

	// 更新数据
	checkDb := tx.Table(checkScan.TableName()).Updates(mo)
	result.Result(checkDb.Error).Process(result.SqlCrudErr)

	// 事务提交
	tx.Commit()
	return response.NewJsonResult(variable.ResponseOkCode, checkDb.RowsAffected, nil)
}
