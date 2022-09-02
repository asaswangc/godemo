package result

// 数据操作类Code
const (
	SelectErrCode      = 400400 // 数据查询失败
	SelectEmptyErrCode = 400401 // 数据查询为空
	SelectExistErrCode = 400402 // 数据已存在
	UpdateErrCode      = 400403 // 数据更新失败
	DeleteErrCode      = 400404 // 数据删除失败
	CreateErrCode      = 400405 // 数据存储失败
	ReqSendErrCode     = 400406 // 请求失败
)

var (
	SelectErr      = &ConstErr{ErrStr: "数据查询失败", ErrCode: SelectErrCode}
	UpdateErr      = &ConstErr{ErrStr: "数据更新失败", ErrCode: UpdateErrCode}
	DeleteErr      = &ConstErr{ErrStr: "数据删除失败", ErrCode: DeleteErrCode}
	CreateErr      = &ConstErr{ErrStr: "数据删除失败", ErrCode: CreateErrCode}
	ReqSendErr     = &ConstErr{ErrStr: "数据删除失败", ErrCode: ReqSendErrCode}
	SelectExistErr = &ConstErr{ErrStr: "数据已经存在", ErrCode: SelectExistErrCode}
	SelectEmptyErr = &ConstErr{ErrStr: "数据查询为空", ErrCode: SelectEmptyErrCode}
)

// 参数校验类Code
const (
	ParamCheckErrCode = 400200 // 参数检验失败
	ParamTypeErrCode  = 400201 // 参数类型错误
	ParamParseErrCode = 400202 // 参数解析失败
)

var (
	ParamTypeErr  = &ConstErr{ErrStr: "参数类型错误", ErrCode: ParamTypeErrCode}
	ParamCheckErr = &ConstErr{ErrStr: "参数检验失败", ErrCode: ParamCheckErrCode}
	ParamParseErr = &ConstErr{ErrStr: "参数解析失败", ErrCode: ParamParseErrCode}
)

// AuthFailCode 用户响应类Code
const (
	AuthFailedCode        = 400100 // 租户认证失败
	LoadCookiesFailedCode = 400101 // 解析Cookie失败
)

var (
	AuthFailedErr     = &ConstErr{ErrStr: "认证失败,用户未登录或session已过期", ErrCode: AuthFailedCode}
	LoadCookiesFailed = &ConstErr{ErrStr: "解析Cookie失败,用户未登录或session已过期", ErrCode: LoadCookiesFailedCode}
)
