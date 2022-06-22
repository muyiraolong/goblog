package sever

import (
	"goblog/router"
	"log"
	"net/http"
)

type Sever struct {
}

var App = &Sever{}

func (*Sever) Start(ip, port string) {
	// 服务器地址
	Serve := http.Server{Addr: ip + ":" + port}
	// 处理请求 //不加（）时，代表直接传入了整个函数
	router.Router()
	if err := Serve.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
