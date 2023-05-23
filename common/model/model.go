package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	dbUser := "api_mzwjc_com"
	dbPass := "4FWJnsAt5fiwxW4d"
	dbName := "test"
	dbAddr := "81.69.7.120"
	dbPort := 3306
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbAddr, dbPort, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "dan_", // 设置表前缀
			SingularTable: true,   // 禁用复数
		},
		Logger: logger.Default.LogMode(logger.Info), //调试用，输出日志
	})
	if err != nil {
		fmt.Println("数据库初始化失败", err)
	}
}
