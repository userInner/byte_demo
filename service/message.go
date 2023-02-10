package service

// import (
// 	// "gorm"
// 	// "log"
// 	"titok_v1/dao"
// 	"titok_v1/models"
// )

type MessageService struct {
	//Token       string `form:"token"`       // 必传：是
	FromUserID  string `form:"from_user_id" bind:"required"`   // 必传：是 发送人id
	ToUserID	string `form:"to_user_id" bind:"required"`	// 必传：是 接收人id
	Content 	string `form:"content" bind:"required"` 	// 必传：是 消息内容
	CreateTime  string `form:"create_time bind:"required"`  // 必传：是 发送时间
}

type SendMessageResp struct {
	StatusCode int64  `json:"status_code"`// 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"` // 返回状态描述
}

//发送消息，初步写完，未测试
func (message *MessageService) SendMessage(c *gin.Context) *resp.SendMessageResp {
	newMessage := &models.Message{
		From_User_ID:  message.FromUserID
		To_User_ID:    message.ToUserID
		Content:       message.Content
		Create_Time:   message.CreateTime
	}

	err := dao.SendMessage(newMessage)
	if err != nil {
		return &resp.SendMessageResp{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  ErrServerInternal,
		}
	}

// }

// //查看聊天记录，没写完
// func GetMessageListByUserID() {

// }