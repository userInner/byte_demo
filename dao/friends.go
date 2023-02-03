package dao

import (
	"time"
	"log"
	"titok_v1/common"
	"titok_v1/models"
)

//是否相互关注
func IsUserFriend(u *models.User, f *models.User) bool {	
	follow_u := &models.Follow{}
	err := common.GetDB().
		Where("user_id=? and to_user_id=? and is_follow=1", u.ID, f.ID).
		Find(follow_u).Error
	if err != nil {
		log.Println("数据库错误", err.Error)
		return false
	}
	flag_u:=follow_u.IsFollow
	follow_f := &models.Follow{}
	err = common.GetDB().
		Where("user_id=? and to_user_id=? and is_follow=1", f.ID, u.ID).
		Find(follow_f).Error
	if err != nil {
		log.Println("数据库错误", err.Error)
		return false
	}
	flag_f:=follow_f.IsFollow
	if flag_f&&flag_u{
		return true
	}else{
		return false
	}
}

// 如果相互关注，就将两者信息加入到friend表中
func AddFriends(u *models.User, f *models.User){
	friend := &models.Friends{}
	common.GetDB().Model(friend).Create(map[string]interface{}{
		"User_id":   u.ID,
		"Friend_id": f.ID,
		"Create_Time": time.Now(),
		"Update_Time": time.Now(),
	})
}


//如果相互关注是好友，就返回好友的信息
func FriendsTB(u *models.User, f *models.User) ([]models.User) {
	var friends []models.User
	if IsUserFriend(u, f) {
		common.GetDB().First(&friends,f.ID)
	}
	return friends
}

