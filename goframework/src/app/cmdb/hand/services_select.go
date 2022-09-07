package cmdb_handler

import (
	"goframework/src/app/cmdb/model"
	"goframework/src/framework/data/mysql_cli"
	"goframework/src/framework/response"
	"goframework/src/framework/result"
	"goframework/variable"
	"gorm.io/gorm"
)

// GetList 通用查询接口
func GetList(mo interface{}, page int, size int, like bool, funcs ...func(*gorm.DB) *gorm.DB) interface{} {
	var (
		total  int64
		gormDB *gorm.DB
		scans  = make([]map[string]interface{}, 0)
	)

	switch mo.(type) {
	case *model.AdminCmdbClass:
		gormDB = mysql_cli.GormDB.Table(mo.(*model.AdminCmdbClass).TableName()).Scopes(funcs...)
		if like {
			gormDB.Scopes(mo.(*model.AdminCmdbClass).QueryByLike)
		} else {
			gormDB.Scopes(mo.(*model.AdminCmdbClass).QueryByAll)
		}
	case *model.AdminCmdbModel:
		gormDB = mysql_cli.GormDB.Table(mo.(*model.AdminCmdbModel).TableName()).Scopes(funcs...)
		if like {
			gormDB.Scopes(mo.(*model.AdminCmdbModel).QueryByLike)
		} else {
			gormDB.Scopes(mo.(*model.AdminCmdbModel).QueryByAll)
		}
	case *model.AdminCmdbField:
		gormDB = mysql_cli.GormDB.Table(mo.(*model.AdminCmdbField).TableName()).Scopes(funcs...)
		if like {
			gormDB.Scopes(mo.(*model.AdminCmdbField).QueryByLike)
		} else {
			gormDB.Scopes(mo.(*model.AdminCmdbField).QueryByAll)
		}
	case *model.AdminCmdbCheck:
		gormDB = mysql_cli.GormDB.Table(mo.(*model.AdminCmdbCheck).TableName()).Scopes(funcs...)
		if like {
			gormDB.Scopes(mo.(*model.AdminCmdbCheck).QueryByLike)
		} else {
			gormDB.Scopes(mo.(*model.AdminCmdbCheck).QueryByAll)
		}
	case *model.AdminCmdbAudit:
		gormDB = mysql_cli.GormDB.Table(mo.(*model.AdminCmdbAudit).TableName()).Scopes(funcs...)
		if like {
			gormDB.Scopes(mo.(*model.AdminCmdbAudit).QueryByLike)
		} else {
			gormDB.Scopes(mo.(*model.AdminCmdbAudit).QueryByAll)
		}
	case *model.AdminShowField:
		gormDB = mysql_cli.GormDB.Table(mo.(*model.AdminShowField).TableName()).Scopes(funcs...)
		if like {
			gormDB.Scopes(mo.(*model.AdminShowField).QueryByLike)
		} else {
			gormDB.Scopes(mo.(*model.AdminShowField).QueryByAll)
		}
	}

	// 不需要分页的查询
	if page+size == 0 {
		result.Result(gormDB.Count(&total)).Process(result.SqlCrudErr)
		result.Result(gormDB.Omit("tenant_id").Find(&scans)).Process(result.SqlCrudErr)
		return response.NewJsonResult(variable.ResponseOkCode, "Success", scans)
	}

	// 需要分页的查询
	result.Result(gormDB.Count(&total)).Process(result.SqlCrudErr)
	result.Result(gormDB.Scopes(Paginate(page, size)).Omit("tenant_id").Find(&scans)).Process(result.SqlCrudErr)
	return response.NewPageResult(variable.ResponseOkCode, total, page, size, scans)
}

// GetHomePage 查询CMDB主页数据
func GetHomePage(funcs ...func(*gorm.DB) *gorm.DB) interface{} {
	var (
		RespScans = make([]map[string]interface{}, 0)
	)

	// 查询资源类
	var classScans []*model.AdminCmdbClass
	result.Result(mysql_cli.GormDB.Model(&model.AdminCmdbClass{}).Scopes(funcs...).Find(&classScans).Error).Process(result.SqlCrudErr)

	// 构建响应数据结构
	for i := 0; i < len(classScans); i++ {
		// 查询模型类
		var modelScans []*model.AdminCmdbModel
		result.Result(mysql_cli.GormDB.Model(&model.AdminCmdbModel{}).Scopes(funcs...).
			Where("class_id = ?", classScans[i].Id).Find(&modelScans).Error).Process(result.SqlCrudErr)
		for j := 0; j < len(modelScans); j++ {
			// 查询模型实例数据数量
			var instanceCount int64
			result.Result(mysql_cli.GormDB.Table(modelScans[j].TableName()).Scopes(funcs...).Count(&instanceCount).Error).Process(result.SqlCrudErr)

			// 构建数据
			RespScans = append(RespScans, map[string]interface{}{
				"class_id":   classScans[i].Id,
				"class_name": classScans[i].ClassName,
				"model_info": map[string]interface{}{
					"model_id":         modelScans[j].Id,
					"model_name":       modelScans[j].ModelName,
					"model_name_zh":    modelScans[j].ModelNameZh,
					"model_inst_count": instanceCount,
				},
			})
		}
	}

	return response.NewJsonResult(variable.ResponseOkCode, "Success", RespScans)
}
