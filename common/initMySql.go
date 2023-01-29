package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	driverName = "mysql"
	port       = "3306"
	host       = "49.232.185.187"
	database   = "douyinApp"
	username   = "root"
	password   = "gegege"
	charset    = "utf8"
)

var (
	DB *gorm.DB
)

// InitMysql 数据库初始化
func InitMysql() {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DriverName:                driverName,
		DSN:                       args,
		DefaultStringSize:         512,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("MySQL LOAD SUCCESS>>>>>>>>>>")
}

func GetDB() *gorm.DB {
	return DB
}
