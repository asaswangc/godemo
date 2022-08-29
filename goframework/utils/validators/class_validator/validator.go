package class_validator

import (
	"goframework/utils/validators"

	"github.com/go-playground/validator/v10"
)

// UserNameValidator 定义验证器名称
const UserNameValidator = "UserName"

// init 注册UserNameValidator验证器
func init() {
	validators.CustomValidationErrors.Set(UserNameValidator, "user_name 字段必须传,并且长度要在2～10之间")
	validators.RegisterValidator(UserNameValidator, UserName("required,min=2,max=10").ToFunc())
}

// UserName 验证器自定义类型
type UserName string

// ToFunc 生成自定义验证器的方法
func (u UserName) ToFunc() validator.Func {
	return func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(string)
		if ok {
			return u.validator(data)
		}
		return false
	}
}

// validator 自定义验证器核心验证器
func (u UserName) validator(data string) bool {
	if err := validators.CustomValidator.Var(data, string(u)); err != nil {
		return false
	}
	return true
}
