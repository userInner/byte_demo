package controllers

import (
	// "fmt"
	"net/http"
	"strconv"
	// "time"
	"titok_v1/dao"
	"titok_v1/dto"
	"titok_v1/middleware"
	"titok_v1/models"
	"titok_v1/response"
	// "titok_v1/service"
	// "titok_v1/utils"

	"github.com/gin-gonic/gin"
)

// 根据传进的用户数据，获取该用户的关注列表
func FriendsList(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	integerUserID, err := strconv.Atoi(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildCommentListDto("1001", response.InvaildParame+err.Error(), nil))
		return
	}
	friends, err := dao.GetFriendsList(&models.User{ID: int64(integerUserID)})
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildFriendsListDto("2001", response.MySqlDataInsertErrorCode+err.Error(), nil))
		return
	}

	_, err = middleware.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildFriendsListDto("1001", response.MySqlDataInsertErrorCode+err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, dto.BuildFriendsListDto("0", response.OKMsg, friends))
}
