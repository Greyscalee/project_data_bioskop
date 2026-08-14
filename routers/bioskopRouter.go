package routers

import (
	"data-bioskop/controllers"

	"github.com/gin-gonic/gin"
)

const bioskopByIDPath = "/bioskops/:bioskopID"

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskops", controllers.CreateBioskop)
	router.GET("/bioskops", controllers.GetBioskops)
	router.GET(bioskopByIDPath, controllers.GetBioskop)
	router.PUT(bioskopByIDPath, controllers.UpdateBioskop)
	router.DELETE(bioskopByIDPath, controllers.DeleteBioskop)

	return router
}
