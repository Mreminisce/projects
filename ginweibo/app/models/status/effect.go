package status

import (
	"ginweibo/database"

	"github.com/lexkong/log"
)

func (s *Status) Create() (err error) {
	if err = database.DB.Create(&s).Error; err != nil {
		log.Warnf("微博创建失败: %v", err)
		return err
	}
	return nil
}

func Delete(id int) (err error) {
	status := &Status{}
	status.BaseModel.ID = uint(id)
	if err = database.DB.Delete(&status).Error; err != nil {
		log.Warnf("微博删除失败: %v", err)
		return err
	}
	return nil
}
