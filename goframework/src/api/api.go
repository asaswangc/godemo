package api

import (
	"fmt"
	"goframework/src/middleware"

	"github.com/gin-gonic/gin"
)

type RegisterApi interface {
	HttpRoutes() *HttpRoutes
	// CreateGrpcRoutes() *HttpRoutes
}

type HttpRoutes struct {
	Engine *gin.Engine
	Groups *gin.RouterGroup
}

func NewApi() *HttpRoutes {
	return &HttpRoutes{Engine: gin.Default()}
}

func (r *HttpRoutes) Service(serviceName string, apiType string, version string) SetGinMwsFunc {
	// 设置服务名
	r.Groups = r.Engine.Group(fmt.Sprintf("/%s/%s/%s", serviceName, apiType, version))

	// 加载默认中间件["错误处理中间件","用户鉴权中间件"]
	r.Groups.Use(middleware.ErrorHandler(), middleware.Authenticate())

	return r.SetGinMws
}

type SetGinMwsFunc func(mws ...gin.HandlerFunc) *HttpRoutes

func (r *HttpRoutes) SetGinMws(mws ...gin.HandlerFunc) *HttpRoutes {
	r.Engine.Use(mws...)
	return r
}

func (r *HttpRoutes) LoadApi(reg RegisterApi) {
	reg.HttpRoutes()
}
