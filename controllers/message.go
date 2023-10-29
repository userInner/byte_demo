package controllers

import (
	// "fmt"
	"log"
	"net/http"
	"strconv"
	// "time"
	"titok_v1/dao"
	"titok_v1/dto"
	"titok_v1/middleware"
	"titok_v1/models"
	response "titok_v1/response"
	"titok_v1/service"
	"titok_v1/utils"

	"github.com/gin-gonic/gin"

	// "github.com/gorilla/websocket"
	// "gopkg.in/fatih/set.v0"
)

// 这个变量好像是全局的，我先注释掉了，sendmessage里有用掉，变量在comment.go里被定义了
// var (
// 	InvaildMsg = "参数错误"
// )

//查看消息列表
func MessageList(c *gin.Context) {
	token := c.Query("token")
	to_user_id:=c.Query("to_user_id")
	integerToUserID, err := strconv.Atoi(to_user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildMessageListDto("1001", response.InvaildParame+err.Error(), nil))
		return
	}
	integerUserID, err := middleware.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildMessageListDto("1001", response.MySqlDataInsertErrorCode+err.Error(), nil))
		return
	}
	messages, err := dao.GetMessageList(&models.User{ID: int64(integerUserID)},&models.User{ID: int64(integerToUserID)})
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildMessageListDto("2001", response.MySqlDataInsertErrorCode+err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, dto.BuildMessageListDto("0", response.OKMsg, messages))
}


// 发送消息,初步写完，未测试
func SendMessage(c *gin.Context) {
	messageServ := &service.SendMessageService{}
	err := c.ShouldBindQuery(messageServ)
	if err != nil {
		// log.Printf("c.ShouldBind(&messageServ): %s\n", err.Error())
		c.JSON(http.StatusBadRequest, dto.BuildMessageDto(response.InvaildParameCode, InvaildMsg+err.Error(), nil))
		return
	}
	log.Printf("messageServ: %v\n", messageServ)

	// 验证token
	u_id, err := middleware.VerifyToken(messageServ.Token)
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildMessageDto(response.InvaildParameCode, InvaildMsg+err.Error(), nil))
		return
	}
	user, err := dao.GetUserByID(u_id)
	if err != nil { //  数据库错误
		c.JSON(404, dto.BuildMessageDto(response.MysqlDataGetErrorCode, response.MySqlDataGetError+err.Error(), nil))
		return
	}
	if user.ID == 0 { // 数据库无该用户
		c.JSON(http.StatusBadRequest, dto.BuildMessageDto(response.InvaildParameCode, InvaildMsg, nil))
		return
	}

	if(messageServ.Content==""){
		response.Fail(c, nil, InvalidParams)
		return
	}

	message := &models.Message{
		FromUserID:  u_id,
		ToUserID:    messageServ.ToUserID,
		Content:       messageServ.Content,
		CreateTime:   utils.GetTime(),
	}

	if messageServ.ActionType == 1 { // 发送聊天

		err := dao.SendMessage(message)
		if err != nil {
			c.JSON(404, dto.BuildMessageDto(response.MySqlDataInsertErrorCode, response.MySqlDataInsertErrorCode+err.Error(), nil))
			return
		}
		
		
		c.JSON(http.StatusOK, dto.BuildMessageDto(response.OKCode, response.OKMsg, message))
		return
	}	



}
