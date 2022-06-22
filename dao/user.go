package dao

import (
	"fmt"
	"goblog/models"
	"log"
)

func GetUserNameById(UserId int) string {
	row := DB.QueryRow("select  user_name from blog_user where uid=?", UserId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(userName, passwd any) *models.User {
	fmt.Println(passwd)
	row := DB.QueryRow("select  * from blog_user where user_name=? and passwd=?", userName, passwd)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	user := new(models.User)
	err := row.Scan(&user.UId, &user.UserName, &user.Passwd, &user.Avatar, &user.CreatAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}
