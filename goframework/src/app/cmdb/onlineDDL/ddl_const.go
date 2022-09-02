package onlineDDL

import (
	"errors"
	"fmt"
)

var (
	FieldNameIsNotNull = errors.New("字段名称不可为空")
	FieldTypeIsIllegal = errors.New("字段类型非法")
)

// 定义支持的数据类型
const (
	Bool    = "bool"
	Varchar = "varchar"
	Nummber = "nummber"
)

// DefaultFunc 函数原形
type DefaultFunc func(args interface{}) string

type Type struct {
	Max     int
	Type    string
	Default DefaultFunc
}

func (t Type) VfMax(max int) int {
	if max >= t.Max {
		return t.Max
	}
	return max
}

// ConstType 数值类型
var ConstType = map[string]*Type{
	Bool: &Type{
		Max:     5,
		Type:    "char",
		Default: func(args interface{}) string { return fmt.Sprintf("default '%v'", args) },
	},

	Nummber: &Type{
		Max:     11,
		Type:    "integer",
		Default: func(args interface{}) string { return fmt.Sprintf("default '%v'", args) },
	},

	Varchar: &Type{
		Max:     65533,
		Type:    "varchar",
		Default: func(args interface{}) string { return fmt.Sprintf("default '%v'", args) },
	},
}
