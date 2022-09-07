package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goframework/src/framework/middleware/session"
	"goframework/src/framework/response"
	"goframework/src/framework/result"
	"goframework/variable"
	"strings"
)

// ErrorHandler 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch ty := err.(type) {
				case *result.ConstErr:
					response.Resp(ctx)(response.NewJsonResult(ty.GetCode(), fmt.Sprintf("%s", ty.Error()), ty.ErrComment))(response.ERR)
				default:
					response.Resp(ctx)(response.NewJsonResult(variable.ResponseInternalErrCode, fmt.Sprintf("%v", ty), nil))(response.SystemERR)
				}
			}
		}()
		ctx.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		switch {
		case strings.Contains(ctx.FullPath(), "/cmdb/service"):
			ctx.Next()
		case variable.Global.Get(variable.RunMode) != variable.ReleaseMode:
			ctx.Set("user_id", 1001)
			ctx.Set("tenant_id", 10013)
			ctx.Next()
		default:
			se, err := session.GetSessionData(ctx)
			if err != nil {
				response.Resp(ctx)(response.NewJsonResult(result.AuthFailedErr.ErrCode, err.Error(), nil))(response.ERR)
			}
			if userId, ok := se["userId"].(float64); !ok {
				response.Resp(ctx)(response.NewJsonResult(result.AuthFailedErr.ErrCode, "用户未登录或session已过期", nil))(response.ERR)
				return
			} else {
				ctx.Set("user_id", int(userId))
			}
			if tenantId, ok := se["tenantId"].(float64); !ok {
				response.Resp(ctx)(response.NewJsonResult(result.AuthFailedErr.ErrCode, "用户未登录或session已过期", nil))(response.ERR)
				return
			} else {
				ctx.Set("tenant_id", int(tenantId))
			}
			ctx.Next()
		}
	}
}
