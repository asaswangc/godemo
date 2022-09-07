package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// AdminCmdbClass 资源模型类别
type AdminCmdbClass struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`            // 自增id
	TenantId  int       `gorm:"column:tenant_id;NOT NULL" json:"tenant_id"`                // 租户id
	IsEnabled string    `gorm:"column:is_enabled;default:true;NOT NULL" json:"is_enabled"` // 是否启用
	ClassName string    `gorm:"column:class_name;NOT NULL" json:"class_name"`              // 资源类名称
	Comments  string    `gorm:"column:comments;NOT NULL" json:"comments"`                  // 备注
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:最后更新时间;NOT NULL" json:"updated_at"`
}

func (m *AdminCmdbClass) TableName() string {
	return "admin_cmdb_class"
}

func (m *AdminCmdbClass) TableNameZh() string {
	return "资源模型类别"
}

func (m *AdminCmdbClass) QueryByAll(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where(m)
}

func (m *AdminCmdbClass) QueryById(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where("id = ?", m.Id)
}

func (m *AdminCmdbClass) QueryByLike(GormDB *gorm.DB) *gorm.DB {
	// 支持模糊查询的字段
	if m.ClassName != "" {
		GormDB = GormDB.Model(&m).Where("class_name like ?", fmt.Sprintf("%%%s%%", m.ClassName))
	}
	return GormDB
}
