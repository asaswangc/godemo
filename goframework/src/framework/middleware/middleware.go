package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"goframework/src/framework/data/redis"
	"goframework/src/framework/response"
	"goframework/src/framework/result"
	"goframework/src/framework/utils"
	"goframework/variable"
	"strings"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 错误处理中间件，拦截所有panic，将其处理为http响应
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch ty := err.(type) {
				case result.ConstErr:
					response.Resp(ctx)(response.NewJsonResult(ty.GetCode(), fmt.Sprintf("%s", ty.Error()), nil))(response.ERR)
				default:
					response.Resp(ctx)(response.NewJsonResult(variable.ResponseInternalErrCode, fmt.Sprintf("%v", ty), nil))(response.SystemERR)
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
					result.Result().Process()(result.LoadCookiesFailed)
					ctx.Abort()
					return
				}

				// 解密Cookie
				value, err := utils.DecryptWithRSA(cookie, "./private_key.pem")
				if err != nil {
					result.Result().Process()(result.LoadCookiesFailed)
					ctx.Abort()
					return
				}

				// 去redis中校验session
				data, _ := redis.RDB.Get(context.TODO(), "tenant:"+value).Result()
				if data == "" {
					result.Result().Process()(result.LoadCookiesFailed)
					ctx.Abort()
					return
				}

				// 从value中取出tenantId
				var res map[string]interface{}
				if err := json.Unmarshal([]byte(data), &res); err != nil {
					result.Result().Process()(result.LoadCookiesFailed)
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
			c.Set("tenant_id", float64(10013))
			c.Next()
		}
	}
}
