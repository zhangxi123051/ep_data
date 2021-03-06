package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangxi123051/ep_data/controller"
	"net/http"
	"strings"
)

func WebRoute(router *gin.Engine) {
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})

	router.GET("/parseJson", controller.ParseJson)
	router.GET("/userLogin", controller.UserLogin)

	router.GET("/testJson", controller.TestJson)

	router.POST("/produce", controller.Produce)
	router.POST("/consume", controller.Consume)

	router.POST("/saramaProduce", controller.SaramaProduce)
	router.POST("/saramaConsume", controller.SaramaConsume)

	//etcd
	router.POST("/putGet", controller.PutGet)
	router.POST("/watch", controller.Watch)
}
