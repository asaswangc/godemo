package gin_err

import (
	"errors"
	"goframework/utils/result"
	"goframework/variable"
)

/* 自定义Error */

var (
	GinBindErr = &TGinBindErr{defaultMsg: "参数解析错误", respCode: variable.ParamParseErrCode}
)

// TGinBindErr Gin 绑定参数Error
type TGinBindErr struct {
	respCode   int    // response code
	message    string // 自定义Error信息
	defaultMsg string // 默认的Error信息
}

func (self *TGinBindErr) Error() string {
	if self.message != "" {
		return self.message
	}
	return self.defaultMsg
}

func (self *TGinBindErr) SetMessage(_ result.CheckErr, err error) (through bool) {
	self.message = err.Error()
	return true
}

func (self *TGinBindErr) GetRespCode() int {
	return self.respCode
}

func (self *TGinBindErr) Process(errs ...error) {
	if len(errs) != 0 {
		self.SetMessage(GinCheck, errs[0])
	}
	panic(self)
}

// ErrHook 这里做个处理Error的钩子
func (self *TGinBindErr) ErrHook(hook ...result.ErrHook) *TGinBindErr {
	hook[0]()
	return self
}

// CallBack 回调函数 消化完err后将其恢复原状(为什么重新弄个实例呢：因为频繁的创建实例会影响基准性能)
func (self *TGinBindErr) CallBack() string {
	// 这里有个技巧：defer比return后执行
	defer func() {
		self.SetMessage(GinCheck, errors.New(""))
	}()
	return self.Error()
}
