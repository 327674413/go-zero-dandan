package dao

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	/*dbUser := ""
	dbPass := ""
	dbName := ""
	dbAddr := ""
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
	*/
}
