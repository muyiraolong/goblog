package views

import (
	"goblog/common"
	"goblog/config"
	"net/http"
)

func (*HTMLapi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}
