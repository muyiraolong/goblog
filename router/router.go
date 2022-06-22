package router

import (
	"goblog/api"
	"goblog/views"
	"net/http"
)

func Router() {
	// 1.返回页面 views 2.返回数据（json）3.返回静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/golang", views.HTML.Index)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/pigenhole", views.HTML.Pigenhole)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)

	// 注意：先处理func，再监听，否则会报错
	// 处理静态资源，筛选掉resource，重定向到public/resource
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
