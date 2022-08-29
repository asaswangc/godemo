package validators

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

var CustomValidator *validator.Validate
var CustomValidationErrors = ValidationErrors{}

// ValidationErrors 验证器报错信息收集器
type ValidationErrors map[string]interface{}

func (self ValidationErrors) Get(key string) (value interface{}, ok bool) {
	value, ok = self[key]
	return
}

func (self ValidationErrors) Set(key string, value interface{}) {
	self[key] = value
}

// Init 验证器初始化
func Init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		CustomValidator = v
	} else {
		log.Fatal("初始化验证器失败")
	}
}

// RegisterValidator 注册验证器
func RegisterValidator(tag string, fn validator.Func) {
	if err := CustomValidator.RegisterValidation(tag, fn); err != nil {
		log.Fatal(fmt.Sprintf("注册%s验证器失败", tag))
	}
}

// InterceptError 如果验证器验证失败,这个将配合error_result包拦截,直接触发panic
func InterceptError(errs error) {
	if errors, ok := errs.(validator.ValidationErrors); ok {
		for i := 0; i < len(errors); i++ {
			if value, ok := CustomValidationErrors.Get(errors[i].Tag()); ok {
				panic(value)
			}
		}
	}
}
