package views

import (
	"goblog/common"
	"goblog/service"
	"net/http"
)

func (*HTMLapi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr, err := service.Writing()
	if err != nil {
		writing.WriteError(w, err)
		return
	}
	writing.WriteData(w, wr)
}
