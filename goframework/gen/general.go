package gen

import (
	"flag"
	"github.com/gin-gonic/gin"
	"goframework/src/framework/data/mysql_cli"
	"goframework/src/framework/utils/cfg"
	"goframework/src/framework/utils/logger"
	"goframework/src/framework/validators"
	"goframework/variable"
	"log"
)

var (
	configPath  = "" // 配置文件目录
	serviceMode = "" // 服务运行模式
	servicePort = "" // 服务监听端口
)

func Init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	flag.StringVar(&configPath, "cfg", "", "配置文件目录")
	flag.StringVar(&servicePort, "port", "8081", "服务监听端口")
	flag.StringVar(&serviceMode, "mode", "debug", "服务运行模式:默认值:debug")
	flag.Parse()

	// 设定gin模式
	if serviceMode == variable.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 配置文件写入全局变量
	variable.Global.Set(variable.ConfPath, configPath)

	// 与运行模式写入全局变量
	variable.Global.Set(variable.RunMode, serviceMode)

	// 依赖加载
	{
		// 加载配置文件
		cfg.Init()

		// 加载日志
		logger.Init()

		// 加载Redis数据库
		//redis.Init()

		// 加载Mysql数据库
		mysql_cli.Init()

		// 加载验证器
		validators.Init("zh")
	}
}
