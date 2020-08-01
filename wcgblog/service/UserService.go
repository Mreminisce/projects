package service

import (
	"wcgblog/model"

	"github.com/jinzhu/gorm"
)

// 获取user
func GetUser(id interface{}) (*model.User, error) {
	var user model.User
	err := db.First(&user, id).Error
	return &user, err
}

// 新增user
func UserInsert(user *model.User) error {
	return db.Save(user).Error
}

// 通过用户名获取用户
func GetUserByUsername(useranme string) (*model.User, error) {
	var user model.User
	e := db.First(&user, "email =?", useranme).Error
	return &user, e
}

// 获取所有用户
func ListUsers() ([]*model.User, error) {
	var user []*model.User
	err := db.Find(&user, "is_admin =?", true).Error
	return user, err
}

// 获取用户的锁
func Lock(user *model.User) error {
	return db.Model(user).Updates(map[string]interface{}{
		"lock_state": user.LockState,
	}).Error
}

// 修改昵称
func UpdateProfile(user *model.User, avatarUrl string) error {
	return db.Model(&user).Update(model.User{AvatarUrl: avatarUrl}).Error
}

// 修改邮箱
func UpdateUserEmail(user *model.User, email string) error {
	if len(email) > 0 {
		return db.Model(user).Update("email", email).Error
	} else {
		return db.Model(user).Update("email", gorm.Expr("NULL")).Error
	}
}
