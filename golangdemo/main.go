package main

import (
	"golangdemo/src/api"
	"golangdemo/src/app/user_service"
)

func main() {
	route := api.NewApi()

	// 用户模块路由
	route.LoadApi(user_service.NewWebApi(route))
	route.LoadApi(user_service.NewServiceApi(route))

	route.Engine.Run()
}
