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
	// 判断用户是否存在
	friend := &models.Friends{}
	if IsUserFriend(u, f){
		common.GetDB().Model(friend).Create(map[string]interface{}{
			"User_id":   u.ID,
			"Friend_id": f.ID,
			"Create_Time": time.Now(),
			"Update_Time": time.Now(),
		})
	}	
}


//如果相互关注是好友，就返回好友的信息
func FriendsTB(u *models.User, f *models.User) ([]models.User) {
	var friends []models.User
	if IsUserFriend(u, f) {
		common.GetDB().First(&friends,f.ID)
	}
	return friends
}

// 直接查询该用户的好友信息
func GetFriendsList(u *models.User) ([]models.User, error) {
	// 先从friend表中查出该用户对应的信息
	var user []models.Friends
	err := common.GetDB().
		Where("user_id=?", u.ID).
		Find(&user).Error
	if err != nil {
		log.Println("数据库错误", err.Error)
		return nil,err
	}
	// 根据查询出来的用户信息中的好友id去user表中获得对应的信息
	var friends []models.User
	for k,_ :=range user{
		var friend []models.User
		// if IsUserFriend(&models.User{ID: u.ID}, &models.User{ID: user[k].ID}) {
			common.GetDB().First(&friends,user[k].Friend_id)
			friends.append(friends,friend)
			// common.GetDB().Select("user_id=?",user[k].Friend_id).Find(&friends)
		// }
	}	
	return friends,nil
}

