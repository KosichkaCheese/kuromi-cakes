package routes

import (
	"kuromi_cakes/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @Summary тест сервера
// @Description Тестирование сервера
// @Tags test
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/ping [get]
func SetupRoutes(cakeController *controller.CakeController, orderController *controller.OrderController) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
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
		api.PUT("/cakes", cakeController.UpdateCake)
		api.DELETE("/cakes/:id", cakeController.DeleteCake)

		api.GET("/orders", orderController.GetAllOrders)
		api.GET("/orders/:id", orderController.GetOrderById)
		api.POST("/orders", orderController.CreateOrder)
		api.DELETE("/orders/:id", orderController.DeleteOrder)
		api.GET("/orders/delivery", orderController.GetAllDeliveries)
		api.POST("/orders/delivery", orderController.AddDelivery)
		api.DELETE("/orders/delivery/:id", orderController.DeleteDelivery)
		api.GET("/orders/payment", orderController.GetAllPayments)
		api.POST("/orders/payment", orderController.AddPayment)
		api.DELETE("/orders/payment/:id", orderController.DeletePayment)

		api.POST("/order_content", orderController.CreateOrderContent)
		api.GET("/order_content/:orderID", orderController.GetOrderContent)
	}
	return router
}
