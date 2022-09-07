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

// 参数校验类Code
const (
	ParamCheckErrCode = 400200 // 参数检验失败
	ParamParseErrCode = 400202 // 参数解析失败
)

// AuthFailCode 用户响应类Code
const (
	AuthFailCode      = 400100 // 认证失败
	AuthErrCode       = 400110 // 租户用户权限验证失败
	UserCreateErrCode = 400101 // 用户注册失败
	UserUpdateErrCode = 400102 // 用户更新失败
	UserDeleteErrCode = 400103 // 用户删除失败
	UserNotFoundCode  = 400104 // 用户不存在
	UserLoginErrCode  = 400105 // 用户登录失败
	UserLogoutErrCode = 400106 // 用户退出登录失败
)

var (
	AuthFailedErr = &ConstErr{ErrStr: "认证失败", ErrCode: AuthFailCode}
)
