package main

import (
	"github.com/zhaozhentao/goblog/app/http/middlewares"
	"github.com/zhaozhentao/goblog/bootstrap"
	"net/http"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
