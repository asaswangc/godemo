package model

import (
	"gorm.io/gorm"
	"time"
)

// AdminCmdbField 资源模型字段
type AdminCmdbField struct {
	Id           int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:自增id" json:"id"`
	TenantId     int       `gorm:"column:tenant_id;type:int(11);comment:租户id;NOT NULL" json:"tenant_id"`
	FieldName    string    `gorm:"column:field_name;type:varchar(50);comment:字段名称;NOT NULL" json:"field_name"`
	FieldNameZh  string    `gorm:"column:field_name_zh;type:varchar(50);comment:字段名称;NOT NULL" json:"field_name_zh"`
	FieldType    string    `gorm:"column:field_type;type:varchar(10);comment:字段类型;NOT NULL" json:"field_type"`
	FieldLength  int       `gorm:"column:field_length;type:int(11);comment:字段长度;NOT NULL" json:"field_length"`
	AllowNotNull string    `gorm:"column:allow_not_null;type:varchar(10);default:true;comment:允许为空;NOT NULL" json:"allow_not_null"`
	CheckId      int       `gorm:"column:verify_id;type:int(11);default:0;comment:验证器id;NOT NULL" json:"verify_id"`
	ModelId      int       `gorm:"column:model_id;type:int(11);comment:资源模型表的id;NOT NULL" json:"model_id"`
	Comments     string    `gorm:"column:comments;type:varchar(200);comment:备注;NOT NULL" json:"comments"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:最后更新时间;NOT NULL" json:"updated_at"`
}

func (m *AdminCmdbField) TableName() string {
	return "admin_cmdb_field"
}

func (m *AdminCmdbField) QueryByAll(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where(m)
}

func (m *AdminCmdbField) QueryById(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where("id = ?", m.Id)
}

func (m *AdminCmdbField) QueryByLike(GormDB *gorm.DB) *gorm.DB {
	// 支持模糊查询的字段
	if m.FieldName != "" {
		GormDB = GormDB.Model(&m).Where("field_name like %?%", m.FieldName)
	}
	if m.FieldNameZh != "" {
		GormDB = GormDB.Model(&m).Where("field_name_zh like %?%", m.FieldNameZh)
	}
	return GormDB
}
