package cmdb

import (
	"github.com/gin-gonic/gin"
	"goframework/src/api"
	"goframework/src/app/cmdb/hand"
)

// ServiceName 定义Service名称
const ServiceName = "cmdb"

type WebApi struct {
	Route *api.HttpRoutes
}

func NewWebApi(route *api.HttpRoutes) *WebApi {
	return &WebApi{Route: route}
}

func (web *WebApi) HttpRoutes() *api.HttpRoutes {
	// cmdb web api
	{
		web := web.Route.Service(ServiceName, "web", "v1")()

		// select
		web.Groups.POST("select/cmdb/home", cmdb_handler.SelectHomeHandler)
		web.Groups.POST("select/audit/list", cmdb_handler.SelectAuditHandler)
		web.Groups.POST("select/class/list", cmdb_handler.SelectClassHandler)
		web.Groups.POST("select/field/list", cmdb_handler.SelectFieldHandler)
		web.Groups.POST("select/model/list", cmdb_handler.SelectModelHandler)
		web.Groups.POST("select/check/list", cmdb_handler.SelectCheckHandler)
		web.Groups.POST("select/show/fields/list", cmdb_handler.SelectShowFHandler)

		// create
		web.Groups.POST("create/class", cmdb_handler.CreateClassHandler)
		web.Groups.POST("create/model", cmdb_handler.CreateModelHandler)
		web.Groups.POST("create/field", cmdb_handler.CreateFieldHandler)
		web.Groups.POST("create/check", cmdb_handler.CreateCheckHandler)
		web.Groups.POST("create/show/fields", cmdb_handler.CreateShowFHandler)

		// update
		web.Groups.POST("update/class", cmdb_handler.UpdateClassHandler)
		web.Groups.POST("update/model", cmdb_handler.UpdateModelHandler)
		web.Groups.POST("update/check", cmdb_handler.UpdateCheckHandler)

		// delete
		web.Groups.POST("delete/check", cmdb_handler.DeleteCheckHandler)
		web.Groups.POST("delete/class", cmdb_handler.DeleteClassHandler)
		web.Groups.POST("delete/model", cmdb_handler.DeleteModelHandler)

	}

	return web.Route
}

type ServiceApi struct {
	Route *api.HttpRoutes
}

func NewServiceApi(route *api.HttpRoutes) *ServiceApi {
	return &ServiceApi{Route: route}
}

func (service *ServiceApi) HttpRoutes() *api.HttpRoutes {
	{
		// 用户模块 service api
		user := service.Route.Service(ServiceName, "service", "v1")()
		user.Groups.GET("get/user/list", func(ctx *gin.Context) {})
	}
	return service.Route
}
