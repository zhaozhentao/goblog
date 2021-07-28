package main

import (
	"github.com/zhaozhentao/goblog/app/http/middlewares"
	"github.com/zhaozhentao/goblog/bootstrap"
	"github.com/zhaozhentao/goblog/config"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
