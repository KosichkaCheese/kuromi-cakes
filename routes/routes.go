package routes

import (
	"kuromi_cakes/controller"

	"github.com/gin-gonic/gin"
)

// @Summary тест сервера
// @Description Тестирование сервера
// @Tags test
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/ping [get]
func SetupRoutes(cakeController *controller.CakeController) *gin.Engine {
	router := gin.Default()
	api := router.Group("/cake_api")
	{

		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		api.GET("/cakes", cakeController.GetAllCakes)
		api.GET("/cakes/:id", cakeController.GetCakeById)
		api.POST("/cakes", cakeController.CreateCake)
	}
	return router
}
