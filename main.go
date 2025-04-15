package main

import (
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"kuromi_cakes/config"
	_ "kuromi_cakes/docs"
	"kuromi_cakes/routes"

	"kuromi_cakes/controller"
	"kuromi_cakes/repository"
	"kuromi_cakes/service"
)

func main() {
	log.Println("Завожу ...дрдрдрдрдр...")

	config.ConnectDB()
	db := config.DB

	cakeRepo := repository.NewCakeRepository(db)
	cakeService := service.NewCakeService(cakeRepo)
	cakeController := controller.NewCakeController(cakeService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService)

	router := routes.SetupRoutes(cakeController, orderController)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8000")
}
