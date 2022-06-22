package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// 映射配置文件
type tomlConfig struct { //需要大写，因为是提供给main.go来读的
	Viewer Viewer //虽然名字与toml中名称不匹配，但是通过toml工具依然可以读取
	System SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	Cfg = new(tomlConfig)
	_, err := toml.DecodeFile("config/config.toml", &Cfg) //将路径提供给后面的Cfg
	if err != nil {
		log.Fatal(err)
	}
	// 手动赋值toml没有的
	Cfg.System.AppName = "awsl-goblog"
	currentdir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentdir
	Cfg.System.Version = 1.0
}
