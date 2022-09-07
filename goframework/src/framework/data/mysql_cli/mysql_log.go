package mysql_cli

import (
	"goframework/src/framework/utils/cfg"
	"goframework/variable"
	"io"
	"log"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"gorm.io/gorm/logger"
)

func loggerFunc() logger.Interface {
	// Sql日志文件配置
	var LogWriter = &lumberjack.Logger{
		MaxAge:     cfg.T.Mysql.Logger.MaxAge,
		Filename:   cfg.T.Mysql.Logger.LogPath,
		MaxSize:    cfg.T.Mysql.Logger.Maxsize,
		Compress:   cfg.T.Mysql.Logger.Compress,
		MaxBackups: cfg.T.Mysql.Logger.MaxBackups,
	}

	// Sql日志配置
	var Logger = func() logger.Interface {
		var (
			IoWriter io.Writer
			LogLevel logger.LogLevel
		)
		switch variable.Global.Get(variable.RunMode) {
		case variable.ReleaseMode:
			IoWriter = LogWriter
			LogLevel = logger.Error
		default:
			IoWriter = os.Stdout
			LogLevel = logger.Info
		}
		return logger.New(log.New(IoWriter, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢查询阀值
			LogLevel:                  LogLevel,               // 日志级别
			IgnoreRecordNotFoundError: false,                  // 是否忽略 RecordNotFoundError
			Colorful:                  false,                  // 是否启用日志颜色
		})
	}

	return Logger()
}
