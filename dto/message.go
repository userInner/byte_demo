package dto

import (
	"time"
	"titok_v1/models"
)

type MessageDto struct {
	Content    string   `json:"content"`     // 聊天内容
	CreateDate time.Time   `json:"create_date"` // 内容发布日期，格式 mm-dd
	ID         int64    `json:"id"`          // 查看聊天的id
}

type MessageListDto struct {
	MessageList []MessageDto `json:"message_list"` // 聊天列表
	StatusCode  string       `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string       `json:"status_msg"`   // 返回状态描述
}

func BuildMessageDto(code string, msg string,message *models.Message) *MessageDto {
	if message == nil {
		return nil
	}
	return &MessageDto{
		Content:    message.Content,
		CreateDate: message.CreateTime,
		ID:         message.ID,
	}
}

func BuildMessageListDto(code string, msg string, message []models.Message) *MessageListDto {
	if len(message) == 0 ||message == nil {
		return &MessageListDto{
			MessageList: nil,
			StatusCode:  code,
			StatusMsg:   msg,
		}
	}
	resMessageDtos := make([]MessageDto, len(message))
	for k, _ := range message {
		resMessageDtos[k].ID = message[k].ID
		resMessageDtos[k].CreateDate = message[k].CreateTime
		resMessageDtos[k].Content = message[k].Content
	}
	return &MessageListDto{
		MessageList: resMessageDtos,
		StatusCode:  code,
		StatusMsg:   msg,
	}
}
