package model

import (
	"gorm.io/gorm"
	"time"
)

// AdminCmdbAudit 审计表
type AdminCmdbAudit struct {
	Id                  int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:自增id" json:"id"`
	TenantId            int       `gorm:"column:tenant_id;type:int(11);comment:租户id;NOT NULL" json:"tenant_id"`
	OperateType         string    `gorm:"column:operate_type;type:varchar(20);comment:操作动作;NOT NULL" json:"operate_type"`
	OperateAfterString  string    `gorm:"column:operate_after_string;type:varchar(500);comment:操作后数据;NOT NULL" json:"operate_after_string"`
	OperateBeforeString string    `gorm:"column:operate_before_string;type:varchar(500);comment:操作前数据;NOT NULL" json:"operate_before_string"`
	OperateObjectName   string    `gorm:"column:operate_object_name;type:varchar(20);comment:操作对象名称;NOT NULL" json:"operate_object_name"`
	OperateInstanceName string    `gorm:"column:operate_instance_name;type:varchar(20);comment:操作实例名称;NOT NULL" json:"operate_instance_name"`
	Comments            string    `gorm:"column:comments;type:varchar(200);comment:备注(这个字段可以拼装一下字段,拼装成一句话);NOT NULL" json:"comments"`
	CreatedAt           time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:最后更新时间;NOT NULL" json:"updated_at"`
}

func (m *AdminCmdbAudit) TableName() string {
	return "admin_cmdb_audit"
}

func (m *AdminCmdbAudit) QueryByAll(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where(m)
}

func (m *AdminCmdbAudit) QueryByLike(GormDB *gorm.DB) *gorm.DB {
	// 支持模糊查询的字段
	if m.OperateObjectName != "" {
		GormDB = GormDB.Model(&m).Where("operate_object_name like %?%", m.OperateObjectName)
	}
	if m.OperateInstanceName != "" {
		GormDB = GormDB.Model(&m).Where("operate_instance_name like %?%", m.OperateInstanceName)
	}
	return GormDB
}
