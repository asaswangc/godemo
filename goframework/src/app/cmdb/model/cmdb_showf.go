package model

import (
	"gorm.io/gorm"
	"time"
)

// AdminShowField 资源模型数据定制显示列
type AdminShowField struct {
	Id        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:自增id" json:"id"`
	UsersId   int       `gorm:"column:users_id;type:int(11);comment:用户id;NOT NULL" json:"users_id"`
	TenantId  int       `gorm:"column:tenant_id;type:int(11);comment:租户id;NOT NULL" json:"tenant_id"`
	ModelId   int       `gorm:"column:model_id;type:int(11);comment:模型id;NOT NULL" json:"model_id"`
	FieldName string    `gorm:"column:field_name;type:varchar(20);comment:字段id;NOT NULL" json:"field_name"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:最后更新时间;NOT NULL" json:"updated_at"`
}

func (m *AdminShowField) TableName() string {
	return "admin_show_field"
}

func (m *AdminShowField) TableNameZh() string {
	return "资源模型数据定制显示列"
}

func (m *AdminShowField) QueryByAll(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where(m)
}

// QueryByLike 不支持模糊查询
func (m *AdminShowField) QueryByLike(GormDB *gorm.DB) *gorm.DB {
	return GormDB
}
