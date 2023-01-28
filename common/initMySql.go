package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"titok_v1/models"
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
	DB  *gorm.DB
	err error
)

// 所有DB相关的初始化操作放在这里
func Init() {
	if err := InitMysql(); err != nil {
		panic(err)
	}

	// 因为之前建表使用了外键，修改表会失败，暂时不管吧
	//if err := InitUserTable(); err != nil {
	//	panic(err)
	//}
	//if err := InitVideoTable(); err != nil {
	//	panic(err)
	//}
	//if err := InitFavTable(); err != nil {
	//	panic(err)
	//}
	//if err := InitFollowTable(); err != nil {
	//	panic(err)
	//}
	//if err := InitCommentTable(); err != nil {
	//	panic(err)
	//}

	log.Println("The database is initialized successful.")
}

// InitMysql 数据库初始化
func InitMysql() error {
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

	return err
}

func InitUserTable() error {
	err := GetDB().AutoMigrate(&models.User{})
	return err
}

func InitVideoTable() error {
	err := GetDB().AutoMigrate(&models.Video{})
	return err
}

func InitFavTable() error {
	err := GetDB().AutoMigrate(&models.Favourite{})
	return err
}

func InitFollowTable() error {
	err := GetDB().AutoMigrate(&models.Follow{})
	return err
}

func InitCommentTable() error {
	err := GetDB().AutoMigrate(&models.Comment{})
	return err
}

func GetDB() *gorm.DB {
	return DB
}
