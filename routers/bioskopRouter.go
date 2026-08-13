package routers

import (
	"data-bioskop/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskops", controllers.CreateBioskop)
	router.GET("/bioskops/:bioskopID", controllers.GetBioskop)

	return router
}
	