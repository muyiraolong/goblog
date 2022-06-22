package views

import (
	"errors"
	"goblog/common"
	"goblog/models"
	"goblog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLapi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	err := r.ParseForm()
	if err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示数量
	pageSize := 10
	// 页面上涉及的数据需要定义
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	var hr = new(models.HomeRerponse)
	if len(slug) != 0 {
		hr, err = service.GetAllIndexInfoBySlug(slug, page, pageSize)
	} else {
		hr, err = service.GetAllIndexInfo(page, pageSize)
	}
	if err != nil {
		log.Println("index获取数据出错", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	index.WriteData(w, hr)
}
