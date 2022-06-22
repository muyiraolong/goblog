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

func (*HTMLapi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	path := r.URL.Path
	log.Println(path)
	pIdStr := strings.TrimPrefix(path, "/p/")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("路径错误"))
		return
	}
	postRes, err := service.GetPostByPostDetailId(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询文章出错"))
	}
	detail.WriteData(w, *postRes)
}
