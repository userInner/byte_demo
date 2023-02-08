package dao

import (
	"log"
	"titok_v1/common"
	"titok_v1/models"
	"errors"
)

//获取聊天记录
func GetMessageList(a *models.User,r *models.User) ([]models.Message, error) {
	content := []models.Message{}
	err := common.GetDB().
		Model(&models.Message{}).
		Where("author_id=? and receiver_id=?", a.ID,r.ID).
		Find(&content).Error
	if err != nil {
		log.Println("数据库错误", err.Error)
		return nil,errors.New("get message failed" + err.Error())
	}
	return content,nil
}

//发送聊天,初步写完，未测试
func SendMessage(newMessage *models.Message) error {
	db := common.GetDB()
	if err := db.Model(&models.Message{}).Create(newMessage).Error; err != nil {
		return err
	}
	return nil
}
