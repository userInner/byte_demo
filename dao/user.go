package dao

import (
	"fmt"
	"titok_v1/common"
	"titok_v1/models"
)

func GetUser() (*models.User, error) {
	var user models.User
	common.DB.First(&user)
	fmt.Println(user.UserName)
	return nil, nil
}
