package models

import "time"

type User struct {
	UId      int       `json:"uid"`
	UserName string    `json:"userName"`
	Passwd   string    `json:"passwd"`
	Avatar   string    `json:"avatar"`
	CreatAt  time.Time `json:"creat_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UserInfo struct {
	UId      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
