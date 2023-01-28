package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"titok_v1/common"
	"titok_v1/dao"
	"titok_v1/routers"
)

func main() {
	r := gin.Default()
	common.InitMysql()
	r = routers.InitRouter(r)
	video, err := dao.GetVideo("")
	fmt.Println(err)
	fmt.Println(video)
	err = r.Run() // 使用gin默认端口 8080
	if err != nil {
		panic(err)
	}
}
