package model

import (
	"gorm.io/gorm"
	"time"
)

// AdminCmdbCheck 资源模型字段验证表
type AdminCmdbCheck struct {
	Id        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:自增id" json:"id"`
	TenantId  int       `gorm:"column:tenant_id;type:int(11);comment:租户id;NOT NULL" json:"tenant_id"`
	IsEnabled string    `gorm:"column:is_enabled;type:varchar(10);default:true;comment:是否启用;NOT NULL" json:"is_enabled"`
	CheckName string    `gorm:"column:verify_name;type:varchar(50);comment:验证器名称;NOT NULL" json:"verify_name"`
	CheckType string    `gorm:"column:verify_type;type:varchar(50);comment:验证器类型;NOT NULL" json:"verify_type"`
	CheckBody string    `gorm:"column:verify_body;type:varchar(100);comment:验证器内容;NOT NULL" json:"verify_body"`
	Comments  string    `gorm:"column:comments;type:varchar(200);comment:备注;NOT NULL" json:"comments"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:最后更新时间;NOT NULL" json:"updated_at"`
}

func (m *AdminCmdbCheck) TableName() string {
	return "admin_cmdb_verify"
}

func (m *AdminCmdbCheck) QueryByAll(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where(m)
}

func (m *AdminCmdbCheck) QueryById(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where("id = ?", m.Id)
}

func (m *AdminCmdbCheck) QueryByLike(GormDB *gorm.DB) *gorm.DB {
	// 支持模糊查询的字段
	if m.CheckName != "" {
		GormDB = GormDB.Model(&m).Where("verify_name like %?%", m.CheckName)
	}
	if m.CheckBody != "" {
		GormDB = GormDB.Model(&m).Where("verify_body like %?%", m.CheckBody)
	}
	return GormDB
}
