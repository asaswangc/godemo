package onlineDDL

import (
	"bytes"
	"fmt"
	"goframework/src/framework/data/mysql"
	"goframework/variable"
	"gorm.io/gorm"
	"text/template"
)

type OptRes struct {
	OptSql    string
	OptName   string
	OptResult interface{}
}

type Table struct {
	Name        string
	Fields      []*Field
	SqlSlice    []string
	OptResSlice []OptRes
}

type Field struct {
	Len       int
	Name      string
	Type      string
	NotNull   string
	Default   interface{}
	Comment   string
	Operation string
}

func New(name string, fields []*Field) *Table {
	return &Table{Name: name, Fields: fields, OptResSlice: []OptRes{}}
}

// SetResult 设置Sql执行/生成结果
func (table *Table) SetResult(optName string, optResult interface{}, optSql string) {
	table.OptResSlice = append(table.OptResSlice, OptRes{
		OptSql:    optSql,
		OptName:   optName,
		OptResult: optResult,
	})
}

// SqlExec 执行sql
func (table *Table) SqlExec(db *gorm.DB) *Table {
	for i := 0; i < len(table.SqlSlice); i++ {
		if err := db.Exec(table.SqlSlice[i]).Error; err != nil {
			table.SetResult("SqlExec", err, table.SqlSlice[i])
		}
	}
	return table
}

func (table *Table) build() (fSqlSlice []string, err error) {
	// 构建建表字段
	for i := 0; i < len(table.Fields); i++ {
		var (
			fLen     int
			fName    string
			fType    string
			fNotNull string
			fDefault string
			fComment string
		)

		// 校验字段长度
		if ConstType[table.Fields[i].Type].Max >= table.Fields[i].Len {
			fLen = table.Fields[i].Len
		} else {
			fLen = table.Fields[i].Len
		}

		// 校验字段名是否为空
		if table.Fields[i].Name == "" {
			return nil, FieldNameIsNotNull
		} else {
			fName = table.Fields[i].Name
		}

		// 校验数据类型是否合法
		if value := ConstType[table.Fields[i].Type]; value == nil {
			return nil, FieldTypeIsIllegal
		} else {
			fType = table.Fields[i].Type
		}

		// 校验字段是否可以为空
		if table.Fields[i].NotNull == variable.TRUE {
			fNotNull = "NOT NULL"
		} else {
			fNotNull = ""
		}

		// 校验是否有默认值
		if table.Fields[i].Default == nil {
			fDefault = ""
		} else {
			fDefault = ConstType[table.Fields[i].Type].Default(table.Fields[i].Default)
		}

		// 校验是否有描述信息
		if table.Fields[i].Comment == "" {
			fComment = ""
		} else {
			fComment = fmt.Sprintf("comment '%s'", table.Fields[i].Comment)
		}
		fSqlSlice = append(fSqlSlice, fmt.Sprintf("%s %s(%d) %s %s %s", fName, fType, fLen, fNotNull, fDefault, fComment))
	}
	return fSqlSlice, nil
}

// Create 创建表
func (table *Table) Create() (*Table, error) {
	// 判断表是否存在
	if mysql.GormDB.Migrator().HasTable(table.Name) {
		return table, fmt.Errorf("表%s已经存在", table.Name)
	}

	// 构建字段
	fSqlSlice, err := table.build()
	if err != nil {
		return table, err
	}

	// 构建建表语句
	var (
		bufSql bytes.Buffer
	)
	tp, err := template.New("tp").Parse(CreateTableSql)
	if err != nil {
		return table, fmt.Errorf("构建建表语句失败,Err:%s", err.Error())
	}
	if err := tp.Execute(&bufSql, struct {
		TableName   string
		TableFields []string
	}{
		TableName:   table.Name,
		TableFields: fSqlSlice,
	}); err != nil {
		return table, fmt.Errorf("渲染建表语句失败,Err:%s", err.Error())
	}
	table.SqlSlice = append(table.SqlSlice, bufSql.String())
	return table, nil
}

// AlterRename 修改表名
func (table *Table) AlterRename(newName string) (*Table, error) {
	if newName == "" {
		return nil, fmt.Errorf("修改表名失败,新表名不可为空")
	}
	table.SqlSlice = append(table.SqlSlice, fmt.Sprintf("alter table `%s` rename to `%s;", table.Name, newName))
	return table, nil
}

const (
	Add    = "add"
	Drop   = "drop"
	Modify = "modify"
	Change = "change"
)

// Alter Alter操作
func (table *Table) Alter(oldName ...string) (*Table, error) {
	for i := 0; i < len(table.Fields); i++ {
		// 构建sql
		fSqlSlice, err := New(table.Name, []*Field{table.Fields[i]}).build()
		if err != nil {
			return table, err
		}
		switch table.Fields[i].Operation {
		case Add:
			table.SqlSlice = append(table.SqlSlice, fmt.Sprintf("alter table %s add %s", table.Name, fSqlSlice[0]))
		case Drop:
			table.SqlSlice = append(table.SqlSlice, fmt.Sprintf("alter table %s drop %s", table.Name, table.Fields[i].Name))
		case Modify:
			table.SqlSlice = append(table.SqlSlice, fmt.Sprintf("alter table %s modify %s", table.Name, fSqlSlice[0]))
		case Change:
			if len(oldName) == 1 {
				table.SqlSlice = append(table.SqlSlice, fmt.Sprintf("alter table %s change %s %s", table.Name, oldName[0], fSqlSlice[0]))
			} else {
				return table, fmt.Errorf("修改字段名需要提供原字段名")
			}
		default:
			return table, fmt.Errorf("操作不合法")
		}
	}
	return table, nil
}
