package main

import (
	"goframework/gen"
	"goframework/src/api"
	"goframework/src/app/cmdb"
	"goframework/src/framework/result"
)

// 项目初始化
func init() {
	gen.Init()
}

func main() {

	route := api.NewApi()

	// Gin信任所有代理,这是不安全的
	result.Result(route.Engine.SetTrustedProxies(nil)).Unwrap()

	// 用户模块路由
	route.LoadApi(cmdb.NewWebApi(route))
	route.LoadApi(cmdb.NewServiceApi(route))

	err := route.Engine.Run()
	if err != nil {
		return
	}
}
