package routes

import (
	"github.com/gin-gonic/gin"
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
}