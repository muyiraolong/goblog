package service

import (
	"errors"
	"goblog/dao"
	"goblog/models"
	"goblog/utils"
)

func Login(userName, passwd any) (*models.LoginRes, error) {
	passwdStr := utils.Md5Crypt(passwd.(string), "hxwdsb")
	user := dao.GetUser(userName, passwdStr)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.UId
	//jwt生成token   加密规则.过期时间.A与B的加密算法
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	userInfo := models.UserInfo{
		user.UId,
		user.UserName,
		user.Avatar,
	}
	loginRes := new(models.LoginRes)
	loginRes.Token = token
	loginRes.UserInfo = userInfo
	return loginRes, nil
}
