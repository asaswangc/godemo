package cmdb_handler

import (
	"github.com/gin-gonic/gin"
	model2 "goframework/src/app/cmdb/model"
	"goframework/src/framework/response"
	"goframework/src/framework/result"
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
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(GetList(&model2.AdminCmdbAudit{
		OperateType:         request.OperateType,
		OperateObjectName:   request.OperateObjectName,
		OperateInstanceName: request.OperateInstanceName,
	}, request.PageNum, CheckPageSize(request.PageSize), request.Like, TenantIdSpace(ctx)))(response.OK)
}

// SelectClassHandler 查询类别
func SelectClassHandler(ctx *gin.Context) {
	var request struct {
		Id        int    `json:"id" binding:"omitempty"`                          // 资源类id
		Like      bool   `json:"like" binding:"omitempty"`                        // 是否模糊查询
		IsEnabled string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		ClassName string `json:"class_name" binding:"omitempty,min=1,max=20"`     // 资源类名称
		PageNum   int    `json:"page_num" binding:"omitempty"`                    // 分页页码
		PageSize  int    `json:"page_size" binding:"omitempty"`                   // 分页数量
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(GetList(&model2.AdminCmdbClass{
		Id:        request.Id,
		IsEnabled: request.IsEnabled,
		ClassName: request.ClassName,
	}, request.PageNum, CheckPageSize(request.PageSize), request.Like, TenantIdSpace(ctx)))(response.OK)
}

// SelectFieldHandler 查询字段
func SelectFieldHandler(ctx *gin.Context) {
	var request struct {
		Id        int    `json:"id" binding:"omitempty"`                      // 资源模型字段id
		Like      bool   `json:"like" binding:"omitempty"`                    // 是否模糊查询
		CheckId   int    `json:"verify_id" binding:"omitempty"`               // 验证器id
		ModelId   int    `json:"model_id" binding:"omitempty"`                // 模型表的id
		FieldName string `json:"field_name" binding:"omitempty,min=3,max=20"` // 字段名称
		FieldType string `json:"field_type" binding:"omitempty,min=3,max=20"` // 字段类型
		PageNum   int    `json:"page_num" binding:"omitempty"`                // 分页页码
		PageSize  int    `json:"page_size" binding:"omitempty"`               // 分页数量
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(GetList(&model2.AdminCmdbField{
		Id:        request.Id,
		CheckId:   request.CheckId,
		ModelId:   request.ModelId,
		FieldName: request.FieldName,
		FieldType: request.FieldType,
	}, request.PageNum, CheckPageSize(request.PageSize), request.Like, TenantIdSpace(ctx)))(response.OK)
}

// SelectModelHandler 查询模型
func SelectModelHandler(ctx *gin.Context) {
	var request struct {
		Id          int    `json:"id" binding:"omitempty"`                          // 资源模型id
		Like        bool   `json:"like" binding:"omitempty"`                        // 是否模糊查询
		ClassId     int    `json:"class_id" binding:"omitempty"`                    // 资源类别id
		IsEnabled   string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		ModelName   string `json:"model_name" binding:"omitempty,min=3,max=20"`     // 资源模型表的名称(前端字段叫标识)
		ModelNameZh string `json:"model_name_zh" binding:"omitempty,min=3,max=20"`  // 资源模型表的名称(中文)(前端字段叫名称)
		PageNum     int    `json:"page_num" binding:"omitempty"`                    // 分页页码
		PageSize    int    `json:"page_size" binding:"omitempty"`                   // 分页数量
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(GetList(&model2.AdminCmdbModel{
		Id:          request.Id,
		ClassId:     request.ClassId,
		IsEnabled:   request.IsEnabled,
		ModelName:   request.ModelName,
		ModelNameZh: request.ModelNameZh,
	}, request.PageNum, CheckPageSize(request.PageSize), request.Like, TenantIdSpace(ctx)))(response.OK)
}

// SelectCheckHandler 查询验证器
func SelectCheckHandler(ctx *gin.Context) {
	var request struct {
		Id        int    `json:"id" binding:"omitempty"`                          // 验证器id
		Like      bool   `json:"like" binding:"omitempty"`                        // 是否模糊查询
		IsEnabled string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		CheckName string `json:"verify_name" binding:"omitempty,min=3,max=50"`    // 验证器名称
		CheckType string `json:"verify_type" binding:"omitempty"`                 // 验证器类型
		PageNum   int    `json:"page_num" binding:"omitempty"`                    // 分页页码
		PageSize  int    `json:"page_size" binding:"omitempty"`                   // 分页数量
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(GetList(&model2.AdminCmdbCheck{
		Id:        request.Id,
		IsEnabled: request.IsEnabled,
		CheckName: request.CheckName,
		CheckType: request.CheckType,
	}, request.PageNum, CheckPageSize(request.PageSize), request.Like, TenantIdSpace(ctx)))(response.OK)
}

// SelectShowFHandler 查询定制列
func SelectShowFHandler(ctx *gin.Context) {
	var request struct {
		ModelId int `json:"model_id" binding:"required"`
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(GetList(&model2.AdminShowField{
		ModelId: request.ModelId,
	}, 0, 0, false, TenantIdSpace(ctx), UserIdSpace(ctx)))(response.OK)
}

// SelectHomeHandler 查询资源目录
func SelectHomeHandler(ctx *gin.Context) {
	response.Resp(ctx)(GetHomePage(TenantIdSpace(ctx)))(response.OK)
}

// CreateClassHandler 创建资源类型
func CreateClassHandler(ctx *gin.Context) {
	var request struct {
		IsEnabled string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		ClassName string `json:"class_name" binding:"required,min=1,max=20"`      // 资源类名称
		Comments  string `json:"comments" binding:"omitempty,min=1,max=200"`      // 备注
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(CreateClass(&model2.AdminCmdbClass{
		ClassName: request.ClassName,
		Comments:  request.Comments,
		TenantId:  GetTenantId(ctx),
	}))(response.OK)
}

// CreateModelHandler 创建资源模型
func CreateModelHandler(ctx *gin.Context) {
	var request struct {
		ClassId     int    `json:"class_id" binding:"required"`                     // 资源类id
		IsEnabled   string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		Comments    string `json:"comments" binding:"omitempty,min=1,max=200"`      // 备注
		ModelName   string `json:"model_name" binding:"required,min=1,max=20"`      // 资源模型表的名称(前端字段叫标识)
		ModelNameZh string `json:"model_name_zh" binding:"required,min=1,max=20"`   // 资源模型表的名称(中文)(前端字段叫名称)
	}
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(CreateModel(&model2.AdminCmdbModel{
		ClassId:     request.ClassId,
		IsEnabled:   request.IsEnabled,
		Comments:    request.Comments,
		ModelName:   request.ModelName,
		ModelNameZh: request.ModelNameZh,
		TenantId:    GetTenantId(ctx),
	}))(response.OK)
}

// CreateFieldHandler 创建模型字段
func CreateFieldHandler(ctx *gin.Context) {
	var request struct {
		Fields []struct {
			Id            int    `json:"id" binding:"omitempty"`                                       // 旧数据的id
			ModelId       int    `json:"model_id" binding:"omitempty"`                                 // 模型id
			FieldName     string `json:"field_name" binding:"omitempty,min=1,max=20"`                  // 字段名称
			FieldNameZh   string `json:"field_name_zh" binding:"omitempty,min=1,max=20"`               // 字段名称
			FieldType     string `json:"field_type" binding:"omitempty,min=3,max=20"`                  // 字段类型
			FieldLength   int    `json:"field_length" binding:"omitempty,min=1,max=10000"`             // 字段长度
			AllowNotNull  string `json:"allow_not_null" binding:"omitempty,oneof=true false"`          // 允许为空
			CheckId       int    `json:"check_id" binding:"omitempty"`                                 // 验证器id
			Comments      string `json:"comments" binding:"omitempty"`                                 // 备注
			OperationType string `json:"operation_type" binding:"required,oneof=update delete create"` // 操作类型
		} `json:"fields" binding:"omitempty"` // 字段
	}
	// 参数解析
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)

	// 构建数据
	var mos = make(map[string][]*model2.AdminCmdbField)
	for i := 0; i < len(request.Fields); i++ {
		mo := &model2.AdminCmdbField{
			Id:           request.Fields[i].Id,
			ModelId:      request.Fields[i].ModelId,
			FieldName:    request.Fields[i].FieldName,
			FieldNameZh:  request.Fields[i].FieldNameZh,
			FieldType:    request.Fields[i].FieldType,
			FieldLength:  request.Fields[i].FieldLength,
			AllowNotNull: request.Fields[i].AllowNotNull,
			CheckId:      request.Fields[i].CheckId,
			Comments:     request.Fields[i].Comments,
			TenantId:     GetTenantId(ctx),
		}

		// 参数校验结果数据
		returnResults := newDdlResults()

		switch request.Fields[i].OperationType {
		case "create":
			// 参数校验
			switch {
			case request.Fields[i].FieldName == "":
				returnResults.Add(false, request.Fields[i].FieldNameZh, "", "create", "新增字段时字段标识参数必填")
				continue
			case request.Fields[i].FieldNameZh == "":
				returnResults.Add(false, request.Fields[i].FieldNameZh, "", "create", "新增字段时字段中文名参数必填")
				continue
			}
			mos["create"] = append(mos["create"], mo)
		case "update":
			// 参数校验
			switch {
			case request.Fields[i].Id == 0:
				returnResults.Add(false, request.Fields[i].FieldNameZh, "", "update", "修改字段时字段id参数必填")
				continue
			}
			mos["update"] = append(mos["update"], mo)
		case "delete":
			// 参数校验
			switch {
			case request.Fields[i].Id == 0:
				returnResults.Add(false, request.Fields[i].FieldNameZh, "", "delete", "删除字段时字段id参数必填")
				continue
			}
			mos["delete"] = append(mos["delete"], mo)
		}
	}
	response.Resp(ctx)(CreateField(mos))(response.OK)
}

// CreateCheckHandler 创建字段验证器
func CreateCheckHandler(ctx *gin.Context) {
	var request struct {
		IsEnabled string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		CheckName string `json:"check_name" binding:"required,min=1,max=50"`      // 验证器名称
		CheckType string `json:"check_type" binding:"required"`                   // 验证器类型
		CheckBody string `json:"check_body" binding:"required"`                   // 验证器内容
		Comments  string `json:"comments" binding:"omitempty"`                    // 备注
	}
	// 参数解析
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(CreateCheck(&model2.AdminCmdbCheck{
		IsEnabled: request.IsEnabled,
		CheckName: request.CheckName,
		CheckType: request.CheckType,
		CheckBody: request.CheckBody,
		Comments:  request.Comments,
		TenantId:  GetTenantId(ctx),
	}))(response.OK)
}

// CreateShowFHandler 创建展示定制列
func CreateShowFHandler(ctx *gin.Context) {
	var request struct {
		Fields  []string `json:"fields" binding:"required"`   // 字段列表
		ModelId int      `json:"model_id" binding:"required"` // 资源模型id
	}
	// 参数解析
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(CreateShowF(&model2.AdminShowField{
		UsersId:  GetUserId(ctx),
		TenantId: GetTenantId(ctx),
		ModelId:  request.ModelId,
	}, request.Fields))(response.OK)
}

// UpdateClassHandler 修改资源类型
func UpdateClassHandler(ctx *gin.Context) {
	var request struct {
		Id        int    `json:"id" binding:"required"`                           // 资源类id
		IsEnabled string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		ClassName string `json:"class_name" binding:"omitempty,min=1,max=20"`     // 资源类名称
		Comments  string `json:"comments" binding:"omitempty,min=1,max=200"`      // 备注
	}
	// 参数解析
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(UpdateClass(&model2.AdminCmdbClass{
		Id:        request.Id,
		IsEnabled: request.IsEnabled,
		ClassName: request.ClassName,
		Comments:  request.Comments,
	}))(response.OK)
}

// UpdateModelHandler 修改资源模型
func UpdateModelHandler(ctx *gin.Context) {
	var request struct {
		Id          int    `json:"id" binding:"required"`                           // 资源模型id
		IsEnabled   string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		Comments    string `json:"comments" binding:"omitempty,min=3,max=200"`      // 备注
		ModelNameZh string `json:"model_name_zh" binding:"omitempty,min=3,max=20"`  // 资源模型表的名称(中文)(前端字段叫名称)
	}
	// 参数解析
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)

	response.Resp(ctx)(UpdateModel(&model2.AdminCmdbModel{
		Id:          request.Id,
		IsEnabled:   request.IsEnabled,
		ModelNameZh: request.ModelNameZh,
		Comments:    request.Comments,
	}))(response.OK)
}

// UpdateCheckHandler 修改字段验证器
func UpdateCheckHandler(ctx *gin.Context) {
	var request struct {
		Id        int    `json:"id" binding:"required"`                           // 验证器id
		IsEnabled string `json:"is_enabled" binding:"omitempty,oneof=true false"` // 是否启用
		CheckName string `json:"check_name" binding:"omitempty,min=3,max=50"`     // 验证器名称
		CheckType string `json:"check_type" binding:"omitempty"`                  // 验证器类型
		CheckBody string `json:"check_body" binding:"omitempty"`                  // 验证器内容
		Comments  string `json:"comments" binding:"omitempty"`                    // 备注
	}
	// 参数解析
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(UpdateCheck(&model2.AdminCmdbCheck{
		Id:        request.Id,
		IsEnabled: request.IsEnabled,
		CheckName: request.CheckName,
		CheckType: request.CheckType,
		CheckBody: request.CheckBody,
		Comments:  request.Comments,
	}))(response.OK)
}

// DeleteClassHandler 删除资源类型
func DeleteClassHandler(ctx *gin.Context) {
	var request struct {
		Id int `json:"id" binding:"required"` // 资源类id
	}
	// 参数解析
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(DeleteClass(request.Id))(response.OK)
}

// DeleteModelHandler 删除资源模型
func DeleteModelHandler(ctx *gin.Context) {
	var request struct {
		Id int `json:"id" binding:"required"` // 资源类id
	}
	// 参数解析
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(DeleteModel(request.Id, GetTenantId(ctx)))(response.OK)
}

// DeleteCheckHandler 删除字段验证器
func DeleteCheckHandler(ctx *gin.Context) {
	var request struct {
		Id int `json:"id" binding:"required"` // 资源类id
	}
	// 参数解析
	result.Result(ctx.ShouldBindJSON(&request)).Process(result.GinBindErr)
	response.Resp(ctx)(DeleteCheck(request.Id))(response.OK)
}
