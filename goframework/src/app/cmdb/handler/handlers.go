package cmdb_handler

import (
	"github.com/gin-gonic/gin"
	"goframework/src/app/cmdb/handler/enter"
	"goframework/src/app/cmdb/handler/model"
	"goframework/src/app/cmdb/handler/service/select"
	"goframework/src/framework/response"
	"goframework/src/framework/result"
	"goframework/variable"
)

// SelectAuditHandler 查询审计
func SelectAuditHandler(ctx *gin.Context) {
	var request struct {
		Like                bool   `json:"like" binding:"omitempty"`                  // 是否模糊查询
		OperateType         string `json:"operate_type" binding:"omitempty"`          // 操作类型
		OperateObjectName   string `json:"operate_object_name" binding:"omitempty"`   // 操作中模型名称
		OperateInstanceName string `json:"operate_instance_name" binding:"omitempty"` // 操作实例名称
		PageNum             int    `json:"page_num" binding:"omitempty"`              // 分页页码
		PageSize            int    `json:"page_size" binding:"omitempty"`             // 分页数量
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process()(result.ParamParseErr)

	rows, total := select_service.GetList(enter.TenantIdFunc(ctx, &model.AdminCmdbAudit{
		OperateType:         request.OperateType,
		OperateObjectName:   request.OperateObjectName,
		OperateInstanceName: request.OperateInstanceName,
	}), request.PageNum, enter.CheckPageSize(request.PageSize), request.Like)

	response.Resp(ctx)(response.NewPageResult(variable.ResponseOkCode, total, rows, request.PageNum, enter.CheckPageSize(request.PageSize)))(response.OK)
}

// SelectClassHandler 查询类别
func SelectClassHandler(ctx *gin.Context) {
	var request struct {
		Id        int    `json:"id" binding:"omitempty"`                          // 资源类id
		IsEnabled string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		ClassName string `json:"class_name" binding:"omitempty,min=1,max=20"`     // 资源类名称
		PageNum   int    `json:"page_num" binding:"omitempty"`                    // 分页页码
		PageSize  int    `json:"page_size" binding:"omitempty"`                   // 分页数量
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process()(result.ParamParseErr)
	rows, total := select_service.GetList(enter.TenantIdFunc(ctx, &model.AdminCmdbClass{
		Id:        request.Id,
		IsEnabled: request.IsEnabled,
		ClassName: request.ClassName,
	}), request.PageNum, enter.CheckPageSize(request.PageSize), true)
	response.Resp(ctx)(response.NewPageResult(variable.ResponseOkCode, total, rows, request.PageNum, enter.CheckPageSize(request.PageSize)))(response.OK)
}

// SelectFieldHandler 查询字段
func SelectFieldHandler(ctx *gin.Context) {
	var request struct {
		Id        int    `json:"id" binding:"omitempty"`                      // 资源模型字段id
		CheckId   int    `json:"verify_id" binding:"omitempty"`               // 验证器id
		ModelId   int    `json:"model_id" binding:"omitempty"`                // 模型表的id
		FieldName string `json:"field_name" binding:"omitempty,min=3,max=20"` // 字段名称
		FieldType string `json:"field_type" binding:"omitempty,min=3,max=20"` // 字段类型
		PageNum   int    `json:"page_num" binding:"omitempty"`                // 分页页码
		PageSize  int    `json:"page_size" binding:"omitempty"`               // 分页数量
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process()(result.ParamParseErr)
	rows, total := select_service.GetList(enter.TenantIdFunc(ctx, &model.AdminCmdbField{
		Id:        request.Id,
		CheckId:   request.CheckId,
		ModelId:   request.ModelId,
		FieldName: request.FieldName,
		FieldType: request.FieldType,
	}), request.PageNum, enter.CheckPageSize(request.PageSize), true)
	response.Resp(ctx)(response.NewPageResult(variable.ResponseOkCode, total, rows, request.PageNum, enter.CheckPageSize(request.PageSize)))(response.OK)
}

// SelectModelHandler 查询模型
func SelectModelHandler(ctx *gin.Context) {
	var request struct {
		Id          int    `json:"id" binding:"omitempty"`                          // 资源模型id
		ClassId     int    `json:"class_id" binding:"omitempty"`                    // 资源类别id
		IsEnabled   string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		ModelName   string `json:"model_name" binding:"omitempty,min=3,max=20"`     // 资源模型表的名称(前端字段叫标识)
		ModelNameZh string `json:"model_name_zh" binding:"omitempty,min=3,max=20"`  // 资源模型表的名称(中文)(前端字段叫名称)
		PageNum     int    `json:"page_num" binding:"omitempty"`                    // 分页页码
		PageSize    int    `json:"page_size" binding:"omitempty"`                   // 分页数量
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process()(result.ParamParseErr)
	rows, total := select_service.GetList(enter.TenantIdFunc(ctx, &model.AdminCmdbModel{
		Id:          request.Id,
		ClassId:     request.ClassId,
		IsEnabled:   request.IsEnabled,
		ModelName:   request.ModelName,
		ModelNameZh: request.ModelNameZh,
	}), request.PageNum, enter.CheckPageSize(request.PageSize), true)
	response.Resp(ctx)(response.NewPageResult(variable.ResponseOkCode, total, rows, request.PageNum, enter.CheckPageSize(request.PageSize)))(response.OK)
}

// SelectCheckHandler 查询验证器
func SelectCheckHandler(ctx *gin.Context) {
	var request struct {
		Id        int    `json:"id" binding:"omitempty"`                          // 验证器id
		IsEnabled string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		CheckName string `json:"verify_name" binding:"omitempty,min=3,max=50"`    // 验证器名称
		CheckType string `json:"verify_type" binding:"omitempty"`                 // 验证器类型
		PageNum   int    `json:"page_num" binding:"omitempty"`                    // 分页页码
		PageSize  int    `json:"page_size" binding:"omitempty"`                   // 分页数量
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process()(result.ParamParseErr)
	rows, total := select_service.GetList(enter.TenantIdFunc(ctx, &model.AdminCmdbCheck{
		Id:        request.Id,
		IsEnabled: request.IsEnabled,
		CheckName: request.CheckName,
		CheckType: request.CheckType,
	}), request.PageNum, enter.CheckPageSize(request.PageSize), true)
	response.Resp(ctx)(response.NewPageResult(variable.ResponseOkCode, total, rows, request.PageNum, enter.CheckPageSize(request.PageSize)))(response.OK)
}
