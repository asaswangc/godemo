package variable

import "reflect"

// 服务启动模式
const (
	RunMode     = "mode"
	ConfPath    = "conf"
	ReleaseMode = "release"
)

// TimeFormat 时间格式化
const TimeFormat = "2006-01-02 15:04:05"

const (
	ResponseOkCode          = 400000 // 常规响应Code，所有动作，成功之后的返回码
	RespinseLimitErr        = 200503 // 超载 服务器暂时无法处理客户端的请求
	ResponseInternalErrCode = 500000 // 系统内部错误响应Code
)

// true false
const (
	TRUE  = "true"
	FALSE = "false"
)

// REGULAR 资源模型字段验证器类型
const (
	REGULAR = "regular" // 正则校验
)

// FieldTypes 资源模型数据类型字段类型定义
var FieldTypes = map[string]interface{}{
	"smallint":  "2 Bytes 大整数值",
	"mediumint": "3 Bytes 大整数值",
	"integer":   "4 Bytes 大整数值",
	"bigint":    "8 Bytes 极大整数值",
	"float":     "4 Bytes 单精度浮点数值",
	"double":    "8 Bytes 双精度浮点数值",
	//"date":       "3 Bytes YYYY-MM-DD 日期值",
	//"time":       "3 Bytes HH:MM:SS 时间值或持续时间",
	//"year":       "1 Bytes YYYY 年份值",
	//"datetime":   "8 Bytes YYYY-MM-DD HH:MM:SS 混合日期和时间值",
	//"timestamp":  "4 Bytes YYYYMMDD HHMMSS 混合日期和时间值，时间戳",
	"char":       "0-255 Bytes 定长字符串",
	"varchar":    "0-65535 Bytes 变长字符串",
	"tinyblob":   "0-255 Bytes 不超过 255 个字符的二进制字符串",
	"tinytext":   "0-255 Bytes 短文本字符串",
	"blob":       "0-65535 Bytes 二进制形式的长文本数据",
	"text":       "0-65535 Bytes 长文本数据",
	"mediumblob": "0-16777215 Bytes 二进制形式的中等长度文本数据",
	"mediumtext": "0-16777215 Bytes 中等长度文本数据",
	"longblob":   "0-294967295 Bytes 二进制形式的极大文本数据",
	"longtext":   "0-294967295 Bytes 极大文本数据",
}

// FieldGoTypes 资源模型数据类型字段类型对应Go类型的定义
var FieldGoTypes = map[string]interface{}{
	"smallint":   reflect.Float64,
	"mediumint":  reflect.Float64,
	"integer":    reflect.Float64,
	"bigint":     reflect.Float64,
	"float":      reflect.Float64,
	"double":     reflect.Float64,
	"date":       reflect.String,
	"time":       reflect.String,
	"year":       reflect.String,
	"datetime":   reflect.String,
	"timestamp":  reflect.String,
	"char":       reflect.String,
	"varchar":    reflect.String,
	"tinyblob":   reflect.String,
	"tinytext":   reflect.String,
	"blob":       reflect.String,
	"text":       reflect.String,
	"mediumblob": reflect.String,
	"mediumtext": reflect.String,
	"longblob":   reflect.String,
	"longtext":   reflect.String,
}

// 用户响应类Code
const (
	AuthFailCode      = 400100 // 认证失败
	UserCreateErrCode = 400101 // 用户注册失败
	UserUpdateErrCode = 400102 // 用户更新失败
	UserDeleteErrCode = 400103 // 用户删除失败
	UserNotFoundCode  = 400104 // 用户不存在
	UserLoginErrCode  = 400105 // 用户登录失败
	UserLogoutErrCode = 400106 // 用户退出登录失败
)

// 参数校验类Code
const (
	ParamCheckErrCode = 400200 // 参数检验失败
	ParamTypeErrCode  = 400201 // 参数类型错误
	ParamParseErrCode = 400202 // 参数解析失败
)

// 配置文件类Code
const (
	ConfigFileNotFoundCode = 400300 // 配置文件不存在
	ConfigFileReadErrCode  = 400301 // 配置文件读取失败
)

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

// 主机类消息
const (
	HostNodeStateUnknownMsg  = "主机状态未知"
	HostNodeStateActiveMsg   = "主机状态正常"
	HostNodeStateInactiveMsg = "主机状态异常"
)

// 服务节点类消息
const (
	ServNodeStateUnknownMsg  = "节点状态未知"
	ServNodeStateActiveMsg   = "节点状态在线"
	ServNodeStateInactiveMsg = "节点状态离线"
)
