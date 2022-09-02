package result

// ProcessFunc 自定义Process函数原型
type ProcessFunc func(err ...*ConstErr)

// ConstErr 自定义Error
type ConstErr struct {
	ErrStr  string // Error消息
	ErrCode int    // Error状态码
}

// Error 实现了原生的Error接口,可用于获取Error信息
func (ce *ConstErr) Error() string {
	return ce.ErrStr
}

// GetCode 用于获取ErrorCode,与项目中定义的Code绑定
func (ce *ConstErr) GetCode() int {
	return ce.ErrCode
}

// Process 当Result捕获到Error时,可使用Result的Process方法来调用本方法,达到使用自定义处理Error的需求
func (ce *ConstErr) Process(pfn ...ProcessFunc) {
	if len(pfn) > 0 {
		pfn[0](ce)
		return
	}
	panic(ce)
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

// Unwrap 自动处理Error, 比较暴躁, 直接panic, 让中间件去兜底
func (er *ErrorResult) Unwrap() interface{} {
	if er.err != nil {
		panic(er.err)
	}
	return er.data
}

// Process 如果Result捕获到了Err,就是用自定义的Err来处理
func (er *ErrorResult) Process(pfn ...ProcessFunc) ProcessFunc {
	return func(ce ...*ConstErr) {
		if er.err != nil {
			ce[0].Process(pfn...)
		}
	}
}
