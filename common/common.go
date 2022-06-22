package common

import (
	"encoding/json"
	"goblog/config"
	"goblog/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate
var w sync.WaitGroup

func LoadTemplate() {
	w.Add(1)
	var err error
	go func() {
		defer w.Done()
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/") //赋值一个template
		if err != nil {
			panic(err) //不需要继续的地方，就不要打印日志了，直接panic挂掉就好了
			// panic与log.fatal区别：panic会执行defer，结束的是当前函数，然后递归向上层传入panic来终止整个应用程序
			// log.fatal=print+os.exit(1)
		}
	}()
	w.Wait()
}

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = 303
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func Success(w http.ResponseWriter, data any) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

// GetRequestJsonParam 获取json参数
func GetRequestJsonParam(r *http.Request) map[string]any {
	var params map[string]any
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
