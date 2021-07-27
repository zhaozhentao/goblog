package bootstrap

import (
	"github.com/gorilla/mux"
	"github.com/zhaozhentao/goblog/pkg/route"
	"github.com/zhaozhentao/goblog/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}
