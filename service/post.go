package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"html/template"
)

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func GetPostByPostDetailId(pId int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pId)
	categoryName := dao.GetCategoryById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	content := []rune(post.Content)
	postMore := models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     models.DateDay(post.CreateAt),
		UpdateAt:     models.DateDay(post.UpdateAt),
	}
	var postRes = &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}
	if err != nil {
		return postRes, err
	}
	return postRes, nil
}

func Writing() (*models.WritingRes, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	wr := &models.WritingRes{
		config.Cfg.Viewer.Title,
		categorys,
		config.Cfg.System.CdnURL,
	}
	return wr, nil
}

func Search(condition string) ([]models.SearchResp, error) {
	posts, err := dao.GetPostSearch(condition)
	if err != nil {
		return nil, err
	}
	searchResps := make([]models.SearchResp, len(posts))
	for i := 0; i < len(posts); i++ {
		searchResps[i].Title = posts[i].Title
		searchResps[i].Pid = posts[i].Pid
	}
	return searchResps, nil
}
