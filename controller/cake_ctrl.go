package controller

import (
	"kuromi_cakes/models"
	"kuromi_cakes/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CakeController struct {
	cakeService service.CakeService
}

func NewCakeController(cakeService service.CakeService) *CakeController {
	return &CakeController{cakeService: cakeService}
}

// GetCakes возвращает список всех тортов
// @Summary Получить все торты
// @Description Возвращает список всех тортов из базы данных
// @Tags cakes
// @Produce json
// @Success 200 {array} models.Cake
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/cakes [get]
func (ctrl *CakeController) GetAllCakes(c *gin.Context) {
	cakes, err := ctrl.cakeService.GetAllCakes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cakes)
}

// @Summary Получение торта по id
// @Description Получает торт из бд по id
// @Tags cakes
// @Produce json
// @Param id path int true "ID торта"
// @Success 200 {object} models.Cake
// @Failure 400 {object} map[string]interface{}"Ошибка при получении торта"
// @Failure 404 {object} map[string]interface{}"Торт не найден"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/cakes/{id} [get]
func (ctrl *CakeController) GetCakeById(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(uri.ID)
	cake, err := ctrl.cakeService.GetCakeById(uri.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cake)
}

// @Summary Создание нового торта
// @Description Добавляет новый торт в базу данных
// @Tags cakes
// @Accept  json
// @Produce json
// @Param cake body models.CakePost true "Торт"
// @Success 200 {object} models.Cake
// @Failure 400 {object} map[string]interface{}"Ошибка при создании торта"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/cakes [post]
func (ctrl *CakeController) CreateCake(c *gin.Context) {
	var cake models.Cake
	if err := c.ShouldBindJSON(&cake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCake, err := ctrl.cakeService.CreateCake(cake)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCake)
}

// @Summary Изменение торта
// @Description Изменяет торт в бд
// @Tags cakes
// @Accept  json
// @Produce json
// @Param cake body models.Cake true "Торт"
// @Success 200 {object} models.Cake
// @Failure 400 {object} map[string]interface{}"Ошибка при создании торта"
// @Failure 404 {object} map[string]interface{}"Торт не найден"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/cakes [put]
func (ctrl *CakeController) UpdateCake(c *gin.Context) {
	var cake models.Cake
	if err := c.ShouldBindJSON(&cake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ctrl.cakeService.GetCakeById(cake.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	newCake, err := ctrl.cakeService.UpdateCake(cake)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCake)
}

// @Summary Удаление торта по id
// @Description Удаляет торт из бд по id
// @Tags cakes
// @Produce json
// @Param id path int true "ID торта"
// @Success 200 {object} models.Cake
// @Failure 400 {object} map[string]interface{}"Ошибка при получении торта"
// @Failure 404 {object} map[string]interface{}"Торт не найден"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/cakes/{id} [delete]
func (ctrl *CakeController) DeleteCake(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ctrl.cakeService.GetCakeById(uri.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.cakeService.DeleteCake(uri.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cake deleted successfully"})
}
