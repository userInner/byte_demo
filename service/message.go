package service

// import (
// 	// "gorm"
// 	// "log"
// 	"titok_v1/dao"
// 	"titok_v1/models"
// )

type MessageService struct {
	//Token       string `form:"token"`       // 必传：是
	AuthorID    string `form:"author_id"`   // 必传：是 发送人id
	ReceiverID	string `form:"receiver_id"`	// 必传：是 接收人id
	MessageText string `form:"content"` 	// 必传：是 消息内容
	SendTime   	string `form:"create_time`  // 必传：是 发送时间
}


// //发送消息，没写完
// func SendMessage() {
	
// }

// //查看聊天记录，没写完
// func GetMessageListByUserID() {

// }