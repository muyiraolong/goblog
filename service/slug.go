package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"html/template"
)

func GetAllIndexInfoBySlug(slug string, page, pageSize int) (*models.HomeRerponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPageBySlug(slug, page, pageSize)
	if err != nil {
		return nil, err
	}
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	total := dao.CountGetAllPostBySlug(slug)
	pageCount := (total-1)/10 + 1
	pages := make([]int, pageCount)
	for i := 1; i <= pageCount; i++ {
		pages[i-1] = i
	}
	var hr = &models.HomeRerponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pageCount,
	}
	return hr, nil
}
