package model

import (
	"errors"
	"fmt"
	"goframework/src/framework/data/mysql_cli"
	"goframework/src/framework/result"
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
	CheckId      int       `gorm:"column:check_id;type:int(11);default:0;comment:验证器id;NOT NULL" json:"verify_id"`
	ModelId      int       `gorm:"column:model_id;type:int(11);comment:资源模型表的id;NOT NULL" json:"model_id"`
	Comments     string    `gorm:"column:comments;type:varchar(200);comment:备注;NOT NULL" json:"comments"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:最后更新时间;NOT NULL" json:"updated_at"`
}

func (m *AdminCmdbField) TableName() string {
	return "admin_cmdb_field"
}

func (m *AdminCmdbField) TableNameZh() string {
	return "资源模型字段"
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

// BeforeCreate 钩子
func (m *AdminCmdbField) BeforeCreate(_ *gorm.DB) (err error) {
	// 基本参数校验
	switch {
	case m.ModelId == 0:

	}

	// 校验资源模型是否存在
	var (
		modelFirst = &AdminCmdbModel{}
		modelModel = &AdminCmdbModel{
			Id:       m.ModelId,
			TenantId: m.TenantId,
		}
	)
	if err = mysql_cli.GormDB.Table(modelModel.TableName()).Where(modelModel).First(&modelFirst).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result.NewConstErr(fmt.Sprintf("模型%d不存在", m.ModelId), result.CreateErrCode, err.Error())
		}
		return result.NewConstErr(err.Error(), result.CreateErrCode)
	}

	// 校验字段是否已经存在
	var (
		fieldFirst = &AdminCmdbField{}
		fieldModel = &AdminCmdbField{
			FieldName: m.FieldName,
			TenantId:  m.TenantId,
		}
	)
	if err = mysql_cli.GormDB.Table(fieldModel.TableName()).Where(fieldModel).First(&fieldFirst).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return result.NewConstErr(err.Error(), result.CreateErrCode)
		}
	}
	if fieldFirst.Id != 0 {
		return result.NewConstErr(fmt.Sprintf("字段%s已存在", m.FieldName), result.CreateErrCode)
	}

	// 校验字段验证器是否存在
	var (
		checkFirst = &AdminCmdbCheck{}
		checkModel = &AdminCmdbCheck{
			Id:       m.CheckId,
			TenantId: m.TenantId,
		}
	)
	if m.CheckId != 0 {
		if err = mysql_cli.GormDB.Table(checkModel.TableName()).Where(checkModel).First(&checkFirst).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return result.NewConstErr(fmt.Sprintf("字段验证器%d不存在", m.CheckId), result.CreateErrCode, err.Error())
			}
			return result.NewConstErr(err.Error(), result.CreateErrCode)
		}
	}

	return nil
}

// BeforeDelete 钩子
func (m *AdminCmdbField) BeforeDelete(_ *gorm.DB) (err error) {
	// 校验字段是否存在
	var (
		fieldFirst = &AdminCmdbField{}
		fieldModel = &AdminCmdbField{
			Id:       m.Id,
			TenantId: m.TenantId,
		}
	)
	if err = mysql_cli.GormDB.Table(fieldModel.TableName()).Where(fieldModel).First(&fieldFirst).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return result.NewConstErr(err.Error(), result.DeleteErrCode)
		}
	}
	if fieldFirst.Id == 0 {
		return result.NewConstErr(fmt.Sprintf("字段%d不存在,Error:%s", m.Id, err.Error()), result.DeleteErrCode)
	}

	// 校验字段下面是否有数据
	var (
		dataCount  = int64(0)
		modelFirst = &AdminCmdbModel{}
		modelModel = &AdminCmdbModel{
			Id:       m.Id,
			TenantId: m.TenantId,
		}
	)
	if err = mysql_cli.GormDB.Table(modelModel.TableName()).Where(modelModel).First(&modelFirst).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result.NewConstErr(fmt.Sprintf("查询模型%d失败,Err:%s", modelModel.Id, err.Error()), result.DeleteErrCode)
		}
	}
	if err = mysql_cli.GormDB.Table(modelFirst.ModelName).Where(fmt.Sprintf("%s != ?", fieldFirst.FieldName), "").Count(&dataCount).Error; err != nil {
		return result.NewConstErr(fmt.Sprintf("校验字段%d下面是否有数据失败,Err:%s", m.Id, err.Error()), result.DeleteErrCode)
	}
	if dataCount != 0 {
		return result.NewConstErr(fmt.Sprintf("字段%d下面有数据,不可删除", m.Id), result.DeleteErrCode)
	}

	return nil
}

// BeforeUpdate 钩子
func (m *AdminCmdbField) BeforeUpdate(_ *gorm.DB) (err error) {
	// 校验字段是否存在
	var (
		fieldFirst = &AdminCmdbField{}
		fieldModel = &AdminCmdbField{
			Id:       m.Id,
			TenantId: m.TenantId,
		}
	)
	if err = mysql_cli.GormDB.Table(fieldModel.TableName()).Where(fieldModel).First(&fieldFirst).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		return result.NewConstErr(err.Error(), result.UpdateErrCode)
	}
	if fieldFirst.Id == 0 {
		return result.NewConstErr(fmt.Sprintf("字段%d不存在", m.Id), result.UpdateErrCode)
	}

	// 校验字段是否重复
	var fieldCount int64
	if err = mysql_cli.GormDB.Table(fieldModel.TableName()).Where("field_name = ?", m.FieldName).Count(&fieldCount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		return result.NewConstErr(err.Error(), result.UpdateErrCode)
	}
	if fieldCount != 0 {
		return result.NewConstErr(fmt.Sprintf("字段%s存在重复", m.FieldName), result.UpdateErrCode)
	}

	// 校验字段验证器是否存在
	var (
		checkFirst = &AdminCmdbCheck{}
		checkModel = &AdminCmdbCheck{
			Id:       m.CheckId,
			TenantId: m.TenantId,
		}
	)
	if m.CheckId != 0 {
		if err = mysql_cli.GormDB.Table(checkModel.TableName()).Where(checkModel).First(&checkFirst).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return result.NewConstErr(fmt.Sprintf("字段验证器%d不存在", m.CheckId), result.CreateErrCode, err.Error())
			}
			return result.NewConstErr(err.Error(), result.CreateErrCode)
		}
	}

	return nil
}
