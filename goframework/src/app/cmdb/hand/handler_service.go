package cmdb_handler

import (
	"github.com/gin-gonic/gin"
	"goframework/src/framework/data/mysql_cli"
	"goframework/src/framework/result"
	"gorm.io/gorm"
)

type ddlResults struct {
	ddlResults []*ddlResult
}

func newDdlResults() *ddlResults {
	return &ddlResults{ddlResults: []*ddlResult{}}
}

type ddlResult struct {
	State      bool   `json:"state"`
	Field      string `json:"fields"`
	ExecSql    string `json:"exec_sql"`
	OptType    string `json:"opt_type"`
	ErrComment string `json:"err_comment"`
}

func (ddl *ddlResults) Get() []*ddlResult {
	return ddl.ddlResults
}

func (ddl *ddlResults) Add(state bool, field string, execSql string, optType string, errComment string) {
	ddl.ddlResults = append(ddl.ddlResults, &ddlResult{
		State:      state,
		Field:      field,
		ExecSql:    execSql,
		OptType:    optType,
		ErrComment: errComment,
	})
}

// Transaction 事务操作
func Transaction() *gorm.DB {
	return mysql_cli.GormDB.Begin()
}

// GetTenantId 获取TenantId
func GetTenantId(ctx *gin.Context) int {
	// 获取TenantId
	var TenantId int
	userId, ok := ctx.Get("tenant_id")
	if ok && userId != nil {
		TenantId = userId.(int)
	} else {
		result.Result(result.AuthFailedErr).Process(func(err error, data ...interface{}) {
			panic(&result.ConstErr{
				ErrStr:  "获取TenantId失败",
				ErrCode: result.AuthErrCode,
			})
		})
	}
	return TenantId
}

// GetUserId 获取UserId
func GetUserId(ctx *gin.Context) int {
	// 获取UserId
	var UserId int
	userId, ok := ctx.Get("user_id")
	if ok && userId != nil {
		UserId = userId.(int)
	} else {
		result.Result(result.AuthFailedErr).Process(func(err error, data ...interface{}) {
			panic(&result.ConstErr{
				ErrStr:  "获取UserId失败",
				ErrCode: result.AuthErrCode,
			})
		})
	}
	return UserId
}

// Namespaces 数据隔离函数原型
type Namespaces func(ctx *gin.Context) func(db *gorm.DB) *gorm.DB

// UserIdSpace 用户id隔离数据
func UserIdSpace(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	// 获取UserId
	var UserId int
	userId, ok := ctx.Get("user_id")
	if ok && userId != nil {
		UserId = userId.(int)
	} else {
		result.Result(result.AuthFailedErr).Process(func(err error, data ...interface{}) {
			panic(&result.ConstErr{
				ErrStr:  "获取UserId失败",
				ErrCode: result.AuthErrCode,
			})
		})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", UserId)
	}
}

// TenantIdSpace 租户id隔离数据
func TenantIdSpace(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var TenantId = 0
		value, ok := ctx.Get("tenant_id")
		if ok && value != nil {
			TenantId = value.(int)
		} else {
			result.Result(result.AuthFailedErr).Process(func(err error, data ...interface{}) {
				panic(&result.ConstErr{
					ErrStr:  "获取TenantId失败",
					ErrCode: result.AuthErrCode,
				})
			})
		}
		return db.Where("tenant_id = ?", TenantId)
	}
}

// Paginate 分页器
func Paginate(Page int, Size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((Page - 1) * Size).Limit(Size)
	}
}

// CheckPageSize 校验分页参数
func CheckPageSize(Size int) int {
	if res := Size - 10; res <= 0 {
		Size = 10
	}
	if res := Size - 100; res >= 0 {
		Size = 100
	}
	return Size
}
