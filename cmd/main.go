package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangxi123051/ep_data/routes"
)

func main() {

	router := gin.New()
	// 加载路由文件中的路由
	routes.WebRoute(router)

	//http://localhost:8000/user/hotzhang/hotaction
	//默认为监听8080端口
	router.Run(":8000")
}
