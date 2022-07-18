package restapi

import (
	"github.com/gin-gonic/gin"
)

func runAPI() {
	router := gin.Default()

	router.GET("/controllers", getControllers)
	router.GET("/cameras", getCameras)

	router.Run("localhost:8080")
}

func getControllers(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, main.GetControllers())
}

func getCameras(c *gin.Context) {

}
