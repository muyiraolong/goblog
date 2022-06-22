package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLapi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	path := r.URL.Path
	log.Println(path)
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("路径错误"))
		return
	}
	err = r.ParseForm()
	if err != nil {
		log.Println("表单获取失败：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	var page int
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	} else {
		page = 1
	}
	//每页显示数量
	pageSize := 10
	categoryResponse, err := service.GetPostByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
