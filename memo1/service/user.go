package service

import (
	"jiuxia/memo1/model"
	"jiuxia/memo1/serializer"
	"jiuxia/memo1/utils"

	"github.com/jinzhu/gorm"
)

type UseService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (service *UseService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "该用户已注册！",
		}
	}
	user.UserName = service.UserName
	user.Email = service.Email
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库创建用户失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}
func (service *UseService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在",
			}
		} else {
			return serializer.Response{
				Status: 500,
				Msg:    "数据库错误",
			}
		}

	}
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	token, err := utils.GenerateToken(service.UserName, service.Email, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "登录成功",
		Data: serializer.TokenData{
			User: serializer.User{
				UserName: service.UserName,
				Email:    service.Email,
			},
			Token: token,
		},
	}
}
