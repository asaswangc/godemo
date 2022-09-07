package mysql_cli

import (
	"fmt"
	"goframework/src/framework/utils/cfg"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	GormDB *gorm.DB
)

func Init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)

	var ConfigMySql = mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.T.Mysql.UserName,
			cfg.T.Mysql.Password,
			cfg.T.Mysql.Host,
			cfg.T.Mysql.Port,
			cfg.T.Mysql.DbName,
		),
		DefaultStringSize:         1024, // string 类型字段的默认长度
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}

	// Gorm配置
	var ConfigGorm = &gorm.Config{
		Logger:                 loggerFunc(),
		SkipDefaultTransaction: true,
	}

	if db, err := gorm.Open(mysql.New(ConfigMySql), ConfigGorm); err != nil {
		log.Fatalf("连接Mysql数据库失败，%s", err)
	} else {
		GormDB = db
	}

	sqlDB, err := GormDB.DB()
	if err != nil {
		log.Fatalf("连接Mysql数据库失败，%s", err)
	}

	sqlDB.SetMaxIdleConns(cfg.T.Mysql.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.T.Mysql.MaxIdleTime))
	sqlDB.SetMaxOpenConns(cfg.T.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(cfg.T.Mysql.MaxLifeTime))
}
