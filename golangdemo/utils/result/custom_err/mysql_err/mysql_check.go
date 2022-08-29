package mysql_err

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// CheckSqlCode 这里可以判断MYSQL报错的code，从而可以更加清晰的将报错原因响应回去
func CheckSqlCode(err error) error {
	MysqlErr, ok := err.(*mysql.MySQLError)
	switch ok {
	case MysqlErr.Number == 1062:
		return fmt.Errorf("数据已存在，%s", err.Error())
	default:
		return err
	}
}

// CheckGeneral 判断Error类型是否是Gorm的
func CheckGeneral(err error) error {
	switch err {
	// 需要将 ErrRecordNotFound 忽略
	case gorm.ErrRecordNotFound:
		return nil
	default:
		return err
	}
}

// IgnoreError 忽略Error
func IgnoreError(_ error) error {
	return nil
}
