package user_service

import (
	"goframework/src/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebApi struct {
	Route *api.HttpRoutes
}

func NewWebApi(route *api.HttpRoutes) *WebApi {
	return &WebApi{Route: route}
}

func (web *WebApi) HttpRoutes() *api.HttpRoutes {

	{
		// 用户模块 web api
		user := web.Route.Service("user", "web", "v1")()
		user.Groups.GET("get/user/list", func(ctx *gin.Context) {})
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
		user := service.Route.Service("user", "service", "v1")()
		user.Groups.GET("get/user/list", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "hello world")
		})
	}

	return service.Route
}
