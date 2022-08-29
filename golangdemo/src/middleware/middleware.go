package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"golangdemo/src/framework/database/redis"
	"golangdemo/src/response"
	"golangdemo/utils"
	"golangdemo/utils/result/custom_err/gin_err"
	"golangdemo/utils/result/custom_err/mysql_err"
	"golangdemo/variable"
	"strings"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 错误处理中间件，拦截所有panic，将其处理为http响应
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				// 过滤Mysql相关的Err
				case *mysql_err.TMysqlErr:
					response.Resp(ctx)(err.(*mysql_err.TMysqlErr).GetRespCode(), err.(*mysql_err.TMysqlErr).CallBack(), nil)(response.ERR)

				// 过滤Gin参数绑定的Err
				case *gin_err.TGinBindErr:
					response.Resp(ctx)(err.(*gin_err.TGinBindErr).GetRespCode(), err.(*gin_err.TGinBindErr).CallBack(), nil)(response.ERR)

				// 这个属于未知的错误，一旦发现应该立马将其建立自定义Error类型，然后在上面过滤掉
				default:
					response.Resp(ctx)(variable.ResponseInternalErrCode, fmt.Sprintf("err:%v", err), nil)(response.SystemERR)
				}
			}
		}()
		ctx.Next()
	}
}

// Authenticate 接口session鉴权中间件
func Authenticate() gin.HandlerFunc {
	switch variable.Global.Get(variable.RunMode) {
	case variable.ReleaseMode:
		return func(ctx *gin.Context) {
			if strings.Contains(ctx.FullPath(), "/send/message") ||
				strings.Contains(ctx.FullPath(), "/send/verify/code") ||
				strings.Contains(ctx.FullPath(), "/tenant/code") {
				ctx.Next()
			} else {
				cookie, err := ctx.Cookie("monitor_session")
				if err != nil || cookie == "" {
					response.Resp(ctx)(variable.AuthFailCode, "", "用户未登录或session已过期")(response.ERR)
					ctx.Abort()
					return
				}

				// 解密Cookie
				value, err := utils.DecryptWithRSA(cookie, "./private_key.pem")
				if err != nil {
					response.Resp(ctx)(variable.AuthFailCode, "", "用户未登录或session已过期")(response.ERR)
					ctx.Abort()
					return
				}

				// 去redis中校验session
				data, _ := redis.RDB.Get(context.TODO(), "tenant:"+value).Result()
				if data == "" {
					response.Resp(ctx)(variable.AuthFailCode, "", "用户未登录或session已过期")(response.ERR)
					ctx.Abort()
					return
				}

				// 从value中取出tenantId
				var res map[string]interface{}
				if err := json.Unmarshal([]byte(data), &res); err != nil {
					response.Resp(ctx)(variable.AuthFailCode, "", "用户未登录或session已过期")(response.ERR)
					ctx.Abort()
					return
				}

				// 将其写到context中
				ctx.Set("tenant_id", fmt.Sprintf("%v", res["tenantInfo"].(map[string]interface{})["tenantId"].(float64)))
				ctx.Next()
			}
		}
	default:
		return func(c *gin.Context) {
			c.Set("tenant_id", float64(10018))
			c.Next()
		}
	}
}
