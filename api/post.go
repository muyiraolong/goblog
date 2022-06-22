package api

import (
	"errors"
	"goblog/common"
	"goblog/dao"
	"goblog/models"
	"goblog/service"
	"goblog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//获取Uid
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
	}
	uid := claim.Uid
	//	post save
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		CIdStr := params["categoryId"].(string)
		CId, err := strconv.Atoi(CIdStr)
		if err != nil {
			panic(err)
		}
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		var postTypes float64
		if params["type"] == nil {
			postTypes = float64(1)
		} else {
			postTypes = params["type"].(float64)
		}
		pType := int(postTypes)
		if err != nil {
			panic(err)
		}
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: CId,
			UserId:     uid,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		//	更新
		params := common.GetRequestJsonParam(r)
		pidFloat := params["pid"].(float64)
		pid := int(pidFloat)
		CIdStr := params["categoryId"].(float64)
		CId := int(CIdStr)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		var postTypes float64
		if params["type"] == nil {
			postTypes = float64(1)
		} else {
			postTypes = params["type"].(float64)
		}
		pType := int(postTypes)
		if err != nil {
			panic(err)
		}
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: CId,
			UserId:     uid,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}
}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(path)
	if err != nil {
		common.Error(w, errors.New("无法识别此路径"))
		return
	}
	post, err := dao.GetPostById(pid)
	if err != nil {
		common.Error(w, errors.New("未找到该文章"))
		return
	}
	common.Success(w, post)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		common.Error(w, errors.New("搜索解析错误"))
	}
	condition := r.Form.Get("val")
	searchResp, err := service.Search(condition)
	if err != nil {
		common.Error(w, errors.New("搜索过程错误"))
	}
	common.Success(w, searchResp)

}
