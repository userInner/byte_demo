package response

type SendMessageResp struct {
	StatusCode int64  `json:"status_code"`// 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"` // 返回状态描述
}

type MessageListResp struct {
	MessageList []Message `json:"message_list"`// 用户列表
	StatusCode  string    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg   *string   `json:"status_msg"`  // 返回状态描述
}

// Message
type Message struct {
	Content    string `json:"content"`    // 消息内容
	CreateTime string `json:"create_time"`// 消息发送时间 yyyy-MM-dd HH:MM:ss
	ID         int64  `json:"id"`         // 消息id
}