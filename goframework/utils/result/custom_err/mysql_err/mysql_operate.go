package mysql_err

import (
	"goframework/utils/result"
	"goframework/variable"

	"github.com/go-sql-driver/mysql"
)

/* 自定义Error */

var (
	CreateErr = &TMysqlErr{defaultMsg: "创建数据失败", respCode: variable.CreateErrCode}
	DeleteErr = &TMysqlErr{defaultMsg: "删除数据失败", respCode: variable.DeleteErrCode}
	UpdateErr = &TMysqlErr{defaultMsg: "修改数据失败", respCode: variable.UpdateErrCode}
	SelectErr = &TMysqlErr{defaultMsg: "查询数据失败", respCode: variable.SelectErrCode}
)

// TMysqlErr Mysql Error
type TMysqlErr struct {
	respCode   int    // response code
	message    string // 自定义Error信息
	defaultMsg string // 默认的Error信息
}

func (self *TMysqlErr) Error() string {
	if self.message != "" {
		return self.message
	}
	return self.defaultMsg
}

func (self *TMysqlErr) GetRespCode() int {
	return self.respCode
}

// SetMessage 这里会将接收到的Error进行检查（check函数），如果检查结果是nil
func (self *TMysqlErr) SetMessage(check result.CheckErr, err error) (through bool) {
	res := check(err)
	if res != nil {
		through = true
		self.message = res.Error()
	}
	return through
}

func (self *TMysqlErr) Process(errs ...error) {
	// 如果是自定了Error就走第定义Error的处理流程，否则直接panic
	if len(errs) != 0 {
		switch errs[0].(type) {
		// 如果Err的类型是MySQLError，那就要去把它的Number换算成Error消息
		case *mysql.MySQLError:
			// 这里判断Error是否会被CheckErrCode函数拦截，拦截成功的话就不触发下面的panic
			if !self.SetMessage(CheckSqlCode, errs[0]) {
				return
			}
		default:
			// 这里判断Error是否会被CheckGeneral函数拦截，拦截成功的话就不触发下面的panic
			if !self.SetMessage(CheckGeneral, errs[0]) {
				return
			}
		}
	}
	panic(self)
}

// ErrHook 这里做个处理Error的钩子
func (self *TMysqlErr) ErrHook(hook ...result.ErrHook) *TMysqlErr {
	hook[0]()
	return self
}

// CallBack 回调函数 消化完err后将其恢复原状(为什么不重新弄个实例呢：因为这个方法可能会频繁的调用描绘导致频繁创建实例，这样会影响基准性能)
func (self *TMysqlErr) CallBack() string {
	// 这里有个技巧：defer比return后执行
	defer func() {
		self.message = ""
	}()
	return self.Error()
}
