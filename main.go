package main

import (
	"goblog/common"
	"goblog/sever"
)

func init() {
	common.LoadTemplate()
}

func main() {
	sever.App.Start("127.0.0.1", "8080")
}
