package enter

import (
	"github.com/gin-gonic/gin"
	"goframework/src/framework/data/mysql"
	"goframework/src/framework/result"
	"gorm.io/gorm"
)

type Interface interface {
	TableName() string
	QueryByAll(GormDB *gorm.DB) *gorm.DB
	QueryByLike(GormDB *gorm.DB) *gorm.DB
}

type BaseModel struct {
	InModel Interface
	GormDB  *gorm.DB
}

func TenantIdFunc(ctx *gin.Context, inModel Interface) *BaseModel {
	// 获取TenantId
	var TenantId = 0
	value, isOK := ctx.Get("tenant_id")
	if isOK && value != nil {
		TenantId = int(value.(float64))
	} else {
		result.Result().Process()(result.AuthFailedErr)
	}

	return &BaseModel{
		InModel: inModel,
		GormDB:  mysql.GormDB.Table(inModel.TableName()).Where("tenant_id = ?", TenantId),
	}
}

// LikeScopes 判断是否使用模糊查询
func LikeScopes(Base *BaseModel, like bool) (Scopes *gorm.DB) {
	if !like {
		Scopes = Base.GormDB.Scopes(Base.InModel.QueryByAll)
	} else {
		Scopes = Base.GormDB.Scopes(Base.InModel.QueryByLike)
	}
	return Scopes
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

func Paginate(Page int, Size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if res := Size - 10; res <= 0 {
			Size = 10
		}
		if res := Size - 100; res >= 0 {
			Size = 100
		}
		offset := (Page - 1) * Size
		return db.Offset(offset).Limit(Size)
	}
}
