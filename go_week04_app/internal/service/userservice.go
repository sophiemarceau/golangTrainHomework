package service

import (
	v1 "go_week04_app/api/go_week04_app/v1"
	"go_week04_app/internal/biz"
	"log"
)

type UserServcie struct {
	uBiz *biz.UserBiz
	log *log.Logger
}

func initUserService(uBiz *biz.UserBiz, log *log.Logger) *UserServcie {
	return &UserServcie{uBiz,log}
}

func (u *UserServcie) ShowUserInfo(id string) *v1.UserInfoDTO {
	userDTO := u.uBiz.Get(id)
	return &v1.UserInfoDTO{
		Name: userDTO.Name,
		UserName: userDTO.UserName,
	}
}
