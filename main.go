package main

import (
	"time"
	"titok_v1/dao"
	"titok_v1/models"
)

func main() {
	// r := gin.Default()
	// common.InitMysql()
	// common.InitMinio()
	// routers.InitRouter(r)
	// err := r.Run() // 使用gin默认端口 8080
	// if err != nil {
	// 	panic(err)
	// }

	var user = models.User{
		ID:            1,
		UserName:      "jack",
		Password:      "123",
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	var to_user = models.User{
		ID:            2,
		UserName:      "tom",
		Password:      "456",
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}

	// dao.InsertUser(&user)
	// dao.InsertUser(&to_user)
	dao.FollowUser(&user, &to_user)
}
