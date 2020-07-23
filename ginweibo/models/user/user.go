package user

import (
	"crypto/md5"
	"encoding/hex"
	"ginweibo/models/database"
	"ginweibo/models"
	"ginweibo/utils/auth"
	"ginweibo/utils/rand"
	"strconv"
	"time"

	"github.com/lexkong/log"
)

type User struct {
	models.BaseModel
	Name            string    `gorm:"column:name;type:varchar(255);not null"`
	Email           string    `gorm:"column:email;type:varchar(255);unique;not null"`
	Avatar          string    `gorm:"column:avatar;type:varchar(255);not null"`
	Password        string    `gorm:"column:password;type:varchar(255);not null"`
	IsAdmin         uint      `gorm:"column:is_admin;type:tinyint(1)"`
	ActivationToken string    `gorm:"column:activation_token;type:varchar(255)"`
	Activated       uint      `gorm:"column:activated;type:tinyint(1);not null"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at"`
	// 记住我标记，存入 cookie 中，下次带上时即可直接登录
	RememberToken string `gorm:"column:remember_token;type:varchar(100)"`
}

func (User) TableName() string {
	return "users"
}

// Encrypt 对密码进行加密
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Compare 验证用户密码
func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *User) Create() (err error) {
	if err = u.Encrypt(); err != nil {
		log.Warnf("用户创建失败: %v", err)
		return err
	}
	if u.RememberToken == "" {
		u.RememberToken = string(rand.RandomCreateBytes(10))
	}
	if u.ActivationToken == "" {
		u.ActivationToken = string(rand.RandomCreateBytes(30))
	}
	if err = database.DB.Create(&u).Error; err != nil {
		log.Warnf("用户创建失败: %v", err)
		return err
	}
	return nil
}

func (u *User) Update(needEncryotPwd bool) (err error) {
	if needEncryotPwd {
		if err = u.Encrypt(); err != nil {
			log.Warnf("用户更新失败: %v", err)
			return err
		}
	}
	if err = database.DB.Save(&u).Error; err != nil {
		log.Warnf("用户更新失败: %v", err)
		return err
	}
	return nil
}

func Delete(id int) (err error) {
	user := &User{}
	user.BaseModel.ID = uint(id)
	// Unscoped: 永久删除 (由于该操作是管理员操作的，所以不使用软删除)
	if err = database.DB.Unscoped().Delete(&user).Error; err != nil {
		log.Warnf("用户删除失败: %v", err)
		return err
	}
	return nil
}

func Get(id int) (*User, error) {
	user := &User{}
	d := database.DB.First(&user, id)
	return user, d.Error
}

func GetByEmail(email string) (*User, error) {
	user := &User{}
	d := database.DB.Where("email = ?", email).First(&user)
	return user, d.Error
}

func GetByActivationToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("activation_token = ?", token).First(&user)
	return user, d.Error
}

func GetByRememberToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("remember_token = ?", token).First(&user)
	return user, d.Error
}

func List(offset, limit int) (users []*User, err error) {
	users = make([]*User, 0)
	if err := database.DB.Offset(offset).Limit(limit).Order("id").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func All() (users []*User, err error) {
	users = make([]*User, 0)
	if err := database.DB.Order("id").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func AllCount() (count int, err error) {
	err = database.DB.Model(&User{}).Count(&count).Error
	return
}

// Gravatar 获取用户头像
func (u *User) Gravatar() string {
	if u.Avatar != "" {
		return u.Avatar
	}
	hash := md5.Sum([]byte(u.Email))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
}

func (u *User) GetIDstring() string {
	return strconv.Itoa(int(u.ID))
}

func (u *User) IsAdminRole() bool {
	return u.IsAdmin == models.TrueTinyint
}

func (u *User) IsActivated() bool {
	return u.Activated == models.TrueTinyint
}
