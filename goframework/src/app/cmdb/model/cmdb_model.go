package model

import (
	"errors"
	"fmt"
	"goframework/src/framework/result"
	"gorm.io/gorm"
	"time"
)

// AdminCmdbModel 资源模型
type AdminCmdbModel struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:自增id" json:"id"`
	TenantId    int       `gorm:"column:tenant_id;type:int(11);comment:租户id;NOT NULL" json:"tenant_id"`
	IsEnabled   string    `gorm:"column:is_enabled;type:varchar(10);default:true;comment:是否启用;NOT NULL" json:"is_enabled"`
	ClassId     int       `gorm:"column:class_id;type:int(11);comment:资源类别id;NOT NULL" json:"class_id"`
	ModelName   string    `gorm:"column:model_name;type:varchar(100);comment:资源模型表的名称;NOT NULL" json:"model_name"`
	ModelNameZh string    `gorm:"column:model_name_zh;type:varchar(20);comment:资源模型表的名称(中文);NOT NULL" json:"model_name_zh"`
	Comments    string    `gorm:"column:comments;type:varchar(200);comment:备注;NOT NULL" json:"comments"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:最后更新时间;NOT NULL" json:"updated_at"`
}

func (m *AdminCmdbModel) TableName() string {
	return "admin_cmdb_model"
}

func (m *AdminCmdbModel) TableNameZh() string {
	return "资源模型"
}

func (m *AdminCmdbModel) QueryByAll(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where(m)
}

func (m *AdminCmdbModel) QueryById(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where("id = ?", m.Id)
}

func (m *AdminCmdbModel) QueryByLike(GormDB *gorm.DB) *gorm.DB {
	// 支持模糊查询的字段
	if m.ModelName != "" {
		GormDB = GormDB.Model(&m).Where("model_name like %?%", m.ModelName)
	}
	if m.ModelNameZh != "" {
		GormDB = GormDB.Model(&m).Where("model_name_zh like %?%", m.ModelNameZh)
	}
	return GormDB
}

// BeforeCreate 校验资源类别是否存在
func (m *AdminCmdbModel) BeforeCreate(tx *gorm.DB) (err error) {
	var (
		classFirst = &AdminCmdbClass{}
		classModel = &AdminCmdbClass{
			Id:       m.ClassId,
			TenantId: m.TenantId,
		}
	)
	if err := tx.Session(&gorm.Session{NewDB: true}).Table(classModel.TableName()).Where(classModel).First(&classFirst).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return result.NewConstErr(fmt.Sprintf("id为%d的资源类别不存在", m.ClassId), result.CreateErrCode, err.Error())
	}
	return nil
}
