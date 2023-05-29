package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var Sqlx *gorm.DB

type DBConn struct {
	Conn sqlx.SqlConn
}

func Connect(dataSource string) *DBConn {
	dbUser := "api_mzwjc_com"
	dbPass := "4FWJnsAt5fiwxW4d"
	dbName := "test"
	dbAddr := "81.69.7.120"
	dbPort := 3306
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", dbUser, dbPass, dbAddr, dbPort, dbName)
	sqlConn := sqlx.NewMysql(dsn)
	return &DBConn{
		Conn: sqlConn,
	}
}
func init() {

}
