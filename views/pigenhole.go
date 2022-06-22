package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"net/http"
)

func (*HTMLapi) Pigenhole(w http.ResponseWriter, r *http.Request) {
	pigenhole := common.Template.Pigeonhole
	pigenholeRes, err := service.FindPostPigenhole()
	if err != nil {
		pigenhole.WriteError(w, errors.New("归档失败"))
		return
	}
	pigenhole.WriteData(w, pigenholeRes)
}
