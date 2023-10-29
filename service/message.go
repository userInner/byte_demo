package service

// import (
// 	// "errors"
// 	// "github.com/gin-gonic/gin"
// 	// "log"
// 	// "net/http"
// 	// "time"
// 	// "titok_v1/dao"
// 	// "titok_v1/middleware"
// 	// "titok_v1/models"
// 	// resp "titok_v1/response"
// 	// "titok_v1/utils"
// )

type SendMessageService struct {
	Token       string `form:"token"`       // 必传：是
	
	// FromUserID  int64 `form:"from_user_id" bind:"required"`   // 必传：是 发送人id
	ToUserID	int64 `form:"to_user_id" bind:"required"`	// 必传：是 接收人id
	Content 	string `form:"content" bind:"required"` 	// 必传：是 消息内容
	ActionType  int32  `form:"action_type" bind:"required"`	// 必传：是 1为发送消息
	// CreateTime  time.Time `form:"create_time bind:"required"`  // 必传：是 发送时间
}



//发送消息，初步写完，未测试
// func (message *SendMessageService) SendMessage(c *gin.Context) *resp.SendMessageResp {
// 	from_user_id :=
// 	newMessage := &models.Message{
// 		FromUserID:  message.FromUserID,
// 		ToUserID:    message.ToUserID,
// 		Content:       message.Content,
// 		CreateTime:   message.CreateTime,
// 	}

// 	err := dao.SendMessage(newMessage)
// 	if err != nil {
// 		return &resp.SendMessageResp{
// 			StatusCode: http.StatusInternalServerError,
// 			StatusMsg:  ErrServerInternal,
// 		}
// 	}

// 	return &resp.SendMessageResp{
// 		StatusCode: 0,
// 		StatusMsg:  "success",
// 	}
// }

// //查看聊天记录，没写完
// func GetMessageListByUserID() {

// }