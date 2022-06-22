package api

import (
	"goblog/common"
	"goblog/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//	接收用户名与密码
	params := common.GetRequestJsonParam(r)
	userName := params["username"]
	passwd := params["passwd"]
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, loginRes)
}
