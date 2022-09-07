package onlineDDL

import (
	"fmt"
	"goframework/src/framework/data/mysql_cli"
	"goframework/variable"
	"gorm.io/gorm"
)

type Field struct {
	Len          int         `json:"len"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	NotNull      string      `json:"not_null"`
	Comment      string      `json:"comment"`
	OldName      string      `json:"old_name"`
	Operation    string      `json:"operation"`
	DefaultValue interface{} `json:"default_value"`
}

// exec 执行sql
func exec(db *gorm.DB, sql string) error {
	if err := db.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func build(field *Field) (string, error) {
	// 构建建表字段
	var (
		sqlSlice string
	)
	var (
		fLen     int
		fName    string
		fType    string
		fNotNull string
		fDefault string
		fComment string
	)

	// 校验字段长度
	if ConstType[field.Type].Max >= field.Len {
		fLen = field.Len
	} else {
		fLen = field.Len
	}

	// 校验字段名是否为空
	if field.Name == "" {
		return sqlSlice, FieldNameIsNotNull
	} else {
		fName = field.Name
	}

	// 校验数据类型是否合法
	if value := ConstType[field.Type]; value == nil {
		return sqlSlice, FieldTypeIsIllegal
	} else {
		fType = field.Type
	}

	// 校验字段是否可以为空
	if field.NotNull == variable.TRUE {
		fNotNull = "NOT NULL"
	} else {
		fNotNull = ""
	}

	// 校验是否有默认值
	if field.DefaultValue == nil {
		fDefault = ""
	} else {
		fDefault = ConstType[field.Type].Default(field.DefaultValue)
	}

	// 校验是否有描述信息
	if field.Comment == "" {
		fComment = ""
	} else {
		fComment = fmt.Sprintf("comment '%s'", field.Comment)
	}
	return fmt.Sprintf("%s %s(%d) %s %s %s", fName, fType, fLen, fNotNull, fDefault, fComment), nil
}

// Create 创建表
func Create(gdb *gorm.DB, tableName string) (sql string, err error) {
	// 判断表是否存在
	if mysql_cli.GormDB.Migrator().HasTable(tableName) {
		return sql, fmt.Errorf("表%s已经存在", tableName)
	}
	// 开始创建表
	if err = exec(gdb, fmt.Sprintf(CreateTableSql, tableName)); err != nil {
		return CreateTableSql, err
	}
	return CreateTableSql, nil
}

// DropTable 删除表
func DropTable(gdb *gorm.DB, tableName string) (sql string, err error) {
	if tableName == "" {
		return "", fmt.Errorf("删除表失败,表名不可为空")
	}
	sql = fmt.Sprintf("drop table `%s`", tableName)
	if err = exec(gdb, sql); err != nil {
		return sql, err
	}
	return sql, nil
}

const (
	Add    = "add"
	Drop   = "drop"
	Modify = "modify"
	Change = "change"
)

// Alter Alter操作
func Alter(gdb *gorm.DB, field *Field, tableName string, args ...string) (sql string, err error) {
	switch field.Operation {
	case Add:
		if sql, err = build(field); err != nil {
			return sql, err
		}
		execSql := fmt.Sprintf("alter table %s add %s", tableName, sql)
		if err = exec(gdb, execSql); err != nil {
			return execSql, err
		}
	case Drop:
		execSql := fmt.Sprintf("alter table %s drop %s", tableName, field.Name)
		if err = exec(gdb, execSql); err != nil {
			return execSql, err
		}
	case Modify: // 更改列属性 modify: alter table {{.表名}} {{.列名}} {{.类型}}
		if sql, err = build(field); err != nil {
			return sql, err
		}
		execSql := fmt.Sprintf("alter table %s modify %s", tableName, sql)
		if err = exec(gdb, execSql); err != nil {
			return execSql, err
		}
	case Change: // 更改列名 change: alter table {{.表名}} change {{.旧列名}} {{.新列名}} {{.类型}}
		if sql, err = build(field); err != nil {
			return sql, err
		}
		execSql := fmt.Sprintf("alter table %s change %s %s", tableName, args[0], sql)
		if err := exec(gdb, execSql); err != nil {
			return execSql, err
		}
	default:
		return "", fmt.Errorf("OnlineDDL操作类型不合法")
	}
	return sql, nil
}
