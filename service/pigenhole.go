package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
)

func FindPostPigenhole() (*models.PigenholeRes, error) {
	//查询所有的文章，进行月份整理
	//	查询所有分类
	posts, err := dao.GetAllPost()
	if err != nil {
		return nil, err
	}
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	pigenholeMap := make(map[string][]models.Post)
	for i := 0; i < len(posts); i++ {
		month := posts[i].CreateAt.Format("2006-01")
		pigenholeMap[month] = append(pigenholeMap[month], posts[i])
	}
	var pigenholeRes = &models.PigenholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigenholeMap,
	}
	return pigenholeRes, nil
}
