package model

import (
	"fmt"
	"goframework/src/framework/result"
	"goframework/src/framework/utils"
	"gorm.io/gorm"
	"time"
)

const (
	CREATE = "新增"
	DELETE = "删除"
	UPDATE = "修改"
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

func (m *AdminCmdbAudit) TableNameZh() string {
	return "审计表"
}

func (m *AdminCmdbAudit) QueryByAll(GormDB *gorm.DB) *gorm.DB {
	return GormDB.Where(m)
}

func (m *AdminCmdbAudit) QueryByLike(GormDB *gorm.DB) *gorm.DB {
	// 支持模糊查询的字段
	if m.OperateObjectName != "" {
		GormDB = GormDB.Model(&m).Where("operate_object_name like ?", fmt.Sprintf("%%%s%%", m.OperateObjectName))
	}
	if m.OperateInstanceName != "" {
		GormDB = GormDB.Model(&m).Where("operate_instance_name like %?%", m.OperateInstanceName)
	}
	return GormDB
}

func (self AdminCmdbAudit) Create(gdb *gorm.DB) error {
	db := gdb.Model(&self).Create(&self)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// CreateAudit 创建数据审计 参数model是创建的数据
func (AdminCmdbAudit) CreateAudit(gdb *gorm.DB, model interface{}) error {
	switch model.(type) {
	case *AdminCmdbClass:
		value := model.(*AdminCmdbClass)
		NewData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         CREATE,
			OperateAfterString:  NewData,
			OperateBeforeString: "",
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.ClassName,
			Comments:            fmt.Sprintf("%s%s:%s", CREATE, value.TableNameZh(), value.ClassName),
		}.Create(gdb)

	case *AdminCmdbModel:
		value := model.(*AdminCmdbModel)
		NewData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         CREATE,
			OperateAfterString:  NewData,
			OperateBeforeString: "",
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.ModelNameZh,
			Comments:            fmt.Sprintf("%s%s:%s", CREATE, value.TableNameZh(), value.ModelNameZh),
		}.Create(gdb)

	case *AdminCmdbField:
		value := model.(*AdminCmdbField)
		NewData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         CREATE,
			OperateAfterString:  NewData,
			OperateBeforeString: "",
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.FieldName,
			Comments:            fmt.Sprintf("%s%s:%s", CREATE, value.TableNameZh(), value.FieldName),
		}.Create(gdb)

	case *AdminCmdbCheck:
		value := model.(*AdminCmdbCheck)
		NewData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         CREATE,
			OperateAfterString:  NewData,
			OperateBeforeString: "",
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.CheckName,
			Comments:            fmt.Sprintf("%s%s:%s", CREATE, value.TableNameZh(), value.CheckName),
		}.Create(gdb)
	}
	return nil
}

// DeleteAudit 删除数据审计 参数delete是被删除的数据
func (AdminCmdbAudit) DeleteAudit(gdb *gorm.DB, delete interface{}) error {
	switch delete.(type) {
	case *AdminCmdbClass:
		value := delete.(*AdminCmdbClass)
		OldData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         DELETE,
			OperateAfterString:  "",
			OperateBeforeString: OldData,
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.ClassName,
			Comments:            fmt.Sprintf("%s%s:%s", DELETE, value.TableNameZh(), value.ClassName),
		}.Create(gdb)

	case *AdminCmdbModel:
		value := delete.(*AdminCmdbModel)
		OldData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         DELETE,
			OperateAfterString:  "",
			OperateBeforeString: OldData,
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.ModelNameZh,
			Comments:            fmt.Sprintf("%s%s:%s", DELETE, value.TableNameZh(), value.ModelNameZh),
		}.Create(gdb)

	case *AdminCmdbField:
		value := delete.(*AdminCmdbField)
		OldData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         DELETE,
			OperateAfterString:  "",
			OperateBeforeString: OldData,
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.FieldName,
			Comments:            fmt.Sprintf("%s%s:%s", DELETE, value.TableNameZh(), value.FieldName),
		}.Create(gdb)

	case *AdminCmdbCheck:
		value := delete.(*AdminCmdbCheck)
		OldData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         DELETE,
			OperateAfterString:  "",
			OperateBeforeString: OldData,
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.CheckName,
			Comments:            fmt.Sprintf("%s%s:%s", DELETE, value.TableNameZh(), value.CheckName),
		}.Create(gdb)
	}
	return nil
}

// UpdateAudit 更新数据审计 参数model是更新后的数据 参数update是更新之前的数据
func (AdminCmdbAudit) UpdateAudit(gdb *gorm.DB, model interface{}, update interface{}) error {
	switch model.(type) {
	case *AdminCmdbClass:
		value := model.(*AdminCmdbClass)
		NewData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		OldData := result.Result(utils.StructToJson(update.(*AdminCmdbClass))).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         UPDATE,
			OperateAfterString:  NewData,
			OperateBeforeString: OldData,
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.ClassName,
			Comments:            fmt.Sprintf("%s%s:%s", UPDATE, value.TableNameZh(), value.ClassName),
		}.Create(gdb)

	case *AdminCmdbModel:
		value := model.(*AdminCmdbModel)
		NewData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		OldData := result.Result(utils.StructToJson(update.(*AdminCmdbModel))).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         UPDATE,
			OperateAfterString:  NewData,
			OperateBeforeString: OldData,
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.ModelNameZh,
			Comments:            fmt.Sprintf("%s%s:%s", UPDATE, value.TableNameZh(), value.ModelNameZh),
		}.Create(gdb)

	case *AdminCmdbField:
		value := model.(*AdminCmdbField)
		NewData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		OldData := result.Result(utils.StructToJson(update.(*AdminCmdbField))).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         UPDATE,
			OperateAfterString:  NewData,
			OperateBeforeString: OldData,
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.FieldName,
			Comments:            fmt.Sprintf("%s%s:%s", UPDATE, value.TableNameZh(), value.FieldName),
		}.Create(gdb)

	case *AdminCmdbCheck:
		value := model.(*AdminCmdbCheck)
		NewData := result.Result(utils.StructToJson(value)).Unwrap().(string)
		OldData := result.Result(utils.StructToJson(update.(*AdminCmdbCheck))).Unwrap().(string)
		return AdminCmdbAudit{
			TenantId:            value.TenantId,
			OperateType:         UPDATE,
			OperateAfterString:  NewData,
			OperateBeforeString: OldData,
			OperateObjectName:   value.TableNameZh(),
			OperateInstanceName: value.CheckName,
			Comments:            fmt.Sprintf("%s%s:%s", UPDATE, value.TableNameZh(), value.CheckName),
		}.Create(gdb)
	}
	return nil
}
