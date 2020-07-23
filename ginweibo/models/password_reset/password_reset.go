package passwordreset

import (
	"ginweibo/models/database"
	"ginweibo/utils/rand"
	"time"

	"github.com/lexkong/log"
)

// 重置密码模型
type PasswordReset struct {
	Email     string    `gorm:"column:email;type:varchar(255);not null" sql:"index"`
	Token     string    `gorm:"column:token;type:varchar(255);not null" sql:"index"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (PasswordReset) TableName() string {
	return "password_resets"
}

func (p *PasswordReset) Create() (err error) {
	token := string(rand.RandomCreateBytes(30))
	// 如已存在则先删除 (可以判断下，不能创建太频繁)
	if existedPwd, err := GetByEmail(p.Email); err == nil {
		if err = DeleteByEmail(existedPwd.Email); err != nil {
			return err
		}
	}
	p.Token = token
	if err = database.DB.Create(&p).Error; err != nil {
		log.Warnf("%s 创建失败: %v", p.TableName(), err)
		return err
	}
	return nil
}

func DeleteByEmail(email string) (err error) {
	pwd := &PasswordReset{}
	if err = database.DB.Where("email = ?", email).Delete(&pwd).Error; err != nil {
		log.Warnf("%s 删除失败: %v", pwd.TableName(), err)
		return err
	}
	return nil
}

func DeleteByToken(token string) (err error) {
	pwd := &PasswordReset{}
	if err = database.DB.Where("token = ?", token).Delete(&pwd).Error; err != nil {
		log.Warnf("%s 删除失败: %v", pwd.TableName(), err)
		return err
	}
	return nil
}

func GetByEmail(email string) (*PasswordReset, error) {
	p := &PasswordReset{}
	d := database.DB.Where("email = ?", email).First(&p)
	return p, d.Error
}

func GetByToken(token string) (*PasswordReset, error) {
	p := &PasswordReset{}
	d := database.DB.Where("token = ?", token).First(&p)
	return p, d.Error
}
