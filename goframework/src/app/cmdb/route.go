package cmdb

import (
	"github.com/gin-gonic/gin"
	"goframework/src/api"
	"goframework/src/app/cmdb/handler"
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
		// select
		user := web.Route.Service(ServiceName, "web", "v1")()
		user.Groups.POST("get/audit/list", cmdb_handler.SelectAuditHandler)
		user.Groups.POST("get/class/list", cmdb_handler.SelectClassHandler)
		user.Groups.POST("get/field/list", cmdb_handler.SelectFieldHandler)
		user.Groups.POST("get/model/list", cmdb_handler.SelectModelHandler)
		user.Groups.POST("get/check/list", cmdb_handler.SelectCheckHandler)
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
