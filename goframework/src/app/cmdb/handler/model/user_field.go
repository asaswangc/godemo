package model

import (
	"gorm.io/gorm"
	"time"
)

// AdminUserField 资源模型数据定制显示列
type AdminUserField struct {
	Id        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:自增id" json:"id"`
	UsersId   int       `gorm:"column:users_id;type:int(11);comment:用户id;NOT NULL" json:"users_id"`
	TenantId  int       `gorm:"column:tenant_id;type:int(11);comment:租户id;NOT NULL" json:"tenant_id"`
	ModelId   int       `gorm:"column:model_id;type:int(11);comment:模型id;NOT NULL" json:"model_id"`
	FieldName string    `gorm:"column:field_name;type:varchar(20);comment:字段id;NOT NULL" json:"field_name"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:最后更新时间;NOT NULL" json:"updated_at"`
}

func (m *AdminUserField) TableName() string {
	return "admin_user_field"
}

func (m *AdminUserField) QueryAll(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where(m)
}

func (m *AdminUserField) QueryById(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where("id = ?", m.Id)
}
