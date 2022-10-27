package routers

import (
	"assignment-3/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", controllers.GetStatus)

	return router
}
