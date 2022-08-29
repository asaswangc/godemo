package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJsonResult(0, "", nil)
		},
	}
}

type PageRes struct {
	Rows     interface{} `json:"rows"`
	Count    int64       `json:"total"`
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
}

type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func NewJsonResult(code int, message string, result interface{}) *JsonResult {
	return &JsonResult{Code: code, Message: message, Data: result}
}

type output func(ctx *gin.Context, v interface{})
type ResultFunc func(code int, message string, result interface{}) func(output output)

func Resp(ctx *gin.Context) ResultFunc {
	return func(code int, message string, result interface{}) func(output output) {
		switch result.(type) {
		case PageRes:
		case map[string]interface{}:
		default:
		}
		ret := ResultPool.Get().(*JsonResult)
		defer ResultPool.Put(ret)
		ret.Code = code
		ret.Data = result
		ret.Message = message
		return func(output output) {
			output(ctx, ret)
		}
	}
}

func OK(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, v)
}

func ERR(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusBadRequest, v)
}

func SystemERR(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusInternalServerError, v)
}
