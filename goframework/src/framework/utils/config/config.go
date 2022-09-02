package config

import (
	"fmt"
	"goframework/src/framework/result"
	"goframework/variable"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var Toml TomlConf

type (
	Mysql struct {
		Host         string      `json:"host" toml:"host"`
		Port         string      `json:"port" toml:"port"`
		DbName       string      `json:"db_name" toml:"db_name"`
		UserName     string      `json:"user_name" toml:"user_name"`
		Password     string      `json:"password" toml:"password"`
		MaxLifeTime  int         `json:"max_life_time" toml:"max_life_time"`
		MaxIdleTime  int         `json:"max_idle_time" toml:"max_idle_time"`
		MaxIdleConns int         `json:"max_idle_conns" toml:"max_idle_conns"`
		MaxOpenConns int         `json:"max_open_conns" toml:"max_open_conns"`
		Logger       MysqlLogger `json:"logger" toml:"logger"`
	}

	MysqlLogger struct {
		MaxAge     int    `json:"max_age" toml:"max_age"`
		Maxsize    int    `json:"max_size" toml:"max_size"`
		LogPath    string `json:"log_path" toml:"log_path"`
		Compress   bool   `json:"compress" toml:"compress"`
		MaxBackups int    `json:"max_backups" toml:"max_backups"`
	}

	Redis struct {
		Host        []string `json:"host"`
		Password    string   `json:"password" toml:"password"`
		PoolSize    int      `json:"pool_size" toml:"pool_size"`
		PoolTimeout int      `json:"pool_timeout" toml:"pool_timeout"`
	}

	ZapLog struct {
		MaxAge      int    `json:"max_age" toml:"max_age"`
		Maxsize     int    `json:"max_size" toml:"max_size"`
		ErrLog      string `json:"err_log" toml:"err_log"`
		WarnLog     string `json:"warn_log" toml:"warn_log"`
		InfoLog     string `json:"info_log" toml:"info_log"`
		MaxBackups  int    `json:"max_backups" toml:"max_backups"`
		Compress    bool   `json:"compress" toml:"compress"`
		JsonEncoder bool   `json:"json_encoder" toml:"json_encoder"`
	}

	TomlConf struct {
		Mysql  Mysql  `json:"mysql" toml:"mysql"`
		Redis  Redis  `json:"redis" toml:"redis"`
		ZapLog ZapLog `json:"zap_log" toml:"zap_log"`
	}
)

func Init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	// 根据启动模式获取配置文件
	var path string
	if variable.Global.Get(variable.ConfPath) == "" {
		path = result.Result(os.Getwd()).Unwrap().(string)
	} else {
		path = variable.Global.Get(variable.ConfPath).(string)
	}
	var file = fmt.Sprintf("cfg%s.toml", variable.Global.Get(variable.RunMode))
	if _, err := toml.DecodeFile(fmt.Sprintf("%s/%s", path, file), &Toml); err != nil {
		log.Fatal("加载配置文件失败")
	}
}
