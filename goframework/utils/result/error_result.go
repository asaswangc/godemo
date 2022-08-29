package result

type ErrHook func()
type CheckErr func(err error) error

// ICustomErr Err消化器需要被消化的自定义Error需要实现此接口
type ICustomErr interface {
	Error() string
	Process(err ...error)
	SetMessage(CheckErr, error) bool
	CallBack() string
}

// ErrorResult Err消化器（不要在意名字）
type ErrorResult struct {
	err  error
	data interface{}
}

// Result 所谓的构造函数
func Result(vs ...interface{}) *ErrorResult {
	var result = new(ErrorResult)
	for i := 0; i < len(vs); i++ {
		// 这里根据类型区分参数是err还是data，这两个参数是互斥的，如果有err那么data就为空，反之亦然
		switch vs[i].(type) {
		case error:
			result.err = vs[i].(error)
		case interface{}:
			result.data = vs[i].(interface{})
		}
	}
	return result
}

// Unwrap 自动处理Error，比较暴躁，直接panic，让中间件去兜底
func (self *ErrorResult) Unwrap() interface{} {
	if self.err != nil {
		panic(self.err)
	}
	// err为空才会将data返回出去
	return self.data
}

// Process 主要处理自定义的Error
func (self *ErrorResult) Process() func(err ICustomErr) {
	return func(err ICustomErr) {
		if self.err != nil {
			err.Process(self.err)
		}
	}
}
