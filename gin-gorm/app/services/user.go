package services

import (
	"errors"
	"gin-gorm/app/common/request"
	"gin-gorm/app/models"
	"gin-gorm/global"
	"gin-gorm/utils"
)

type userService struct {
}

var UserService = new(userService)

// AddUser 创建用户
func (userService *userService) AddUser(req request.User) (err error, user models.User) {
	var result = global.DB.Where("mobile = ?", req.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: req.Name, Mobile: req.Mobile, Password: utils.BcryptMake([]byte(req.Password))}
	err = global.DB.Create(&user).Error
	return
}
