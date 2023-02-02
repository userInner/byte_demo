package common

import (
	"fmt"
	"log"
	"titok_v1/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	driverName = "mysql"
	port       = "3306"
	host       = "49.232.185.187"
	database   = "douyinApp"
	username   = "root"
	password   = "gegege"
	charset    = "utf8mb4"
)

var (
	DB *gorm.DB
)

// 所有DB相关的初始化操作放在这里
func init() {
	if err := InitMysql(); err != nil {
		panic(err)
	}

	// // 表都已经创建好了，所以不用执行以下代码
<<<<<<< HEAD
	// if err := InitUserTable(); err != nil {
	// 	panic(err)
	// }
	// if err := InitVideoTable(); err != nil {
	// 	panic(err)
	// }
	// if err := InitFavTable(); err != nil {
	// 	panic(err)
	// }
	// if err := InitFollowTable(); err != nil {
	// 	panic(err)
	// }
	// if err := InitCommentTable(); err != nil {
	// 	panic(err)
	// }
	// if err := InitFriendsTable(); err != nil {
	// 	panic(err)
	// }
	// if err := InitMessageTable(); err != nil {
	// 	panic(err)
	// }
=======
	if err := InitUserTable(); err != nil {
		panic(err)
	}
	if err := InitVideoTable(); err != nil {
		panic(err)
	}
	if err := InitFavTable(); err != nil {
		panic(err)
	}
	if err := InitFollowTable(); err != nil {
		panic(err)
	}
	if err := InitCommentTable(); err != nil {
		panic(err)
	}
	if err := InitFriendsTable(); err != nil {
		panic(err)
	}
	if err := InitMessageTable(); err != nil {
		panic(err)
	}
>>>>>>> b2cb668523580da494ed0f502e9f763dc42b5086

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
	fmt.Println("database connect success!")
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

func InitFriendsTable() error {
	err := GetDB().AutoMigrate(&models.Friends{})
	return err
}

func InitMessageTable() error {
	err := GetDB().AutoMigrate(&models.Message{})
	return err
}

func GetDB() *gorm.DB {
	return DB
}
