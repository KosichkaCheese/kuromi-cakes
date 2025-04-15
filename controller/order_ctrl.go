package controller

import (
	"kuromi_cakes/models"
	"kuromi_cakes/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

// GetCakes возвращает список всех заказов
// @Summary Получить все заказы
// @Description Возвращает список всех заказов из базы данных
// @Tags orders
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders [get]
func (ctrl *OrderController) GetAllOrders(c *gin.Context) {
	orders, err := ctrl.orderService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// @Summary Получение заказа по id
// @Description Получает заказ из бд по id
// @Tags orders
// @Produce json
// @Param id path int true "ID заказа"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]interface{}"Ошибка при получении заказа"
// @Failure 404 {object} map[string]interface{}"Заказ не найден"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders/{id} [get]
func (ctrl *OrderController) GetOrderById(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := ctrl.orderService.GetOrderById(uri.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary Создание нового заказа
// @Description Добавляет новый заказ в базу данных
// @Tags orders
// @Accept  json
// @Produce json
// @Param order body models.OrderPost true "Заказ"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]interface{}"Ошибка при создании заказа"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders [post]
func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var order models.OrderPost
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newOrder, err := ctrl.orderService.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newOrder)
}

// @Summary Удаление заказа по id
// @Description Удаляет заказ из бд по id
// @Tags orders
// @Produce json
// @Param id path int true "ID заказа"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]interface{}"Ошибка при получении заказа"
// @Failure 404 {object} map[string]interface{}"Заказ не найден"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders/{id} [delete]
func (ctrl *OrderController) DeleteOrder(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ctrl.orderService.GetOrderById(uri.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.orderService.DeleteOrder(uri.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// @Summary Добавление способа доставки
// @Description Добавляет способ доставки в базу данных
// @Tags orders
// @Accept  json
// @Produce json
// @Param delivery body string true "Название способа доставки"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}"Ошибка при добавлении способа доставки"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders/delivery [post]
func (ctrl *OrderController) AddDelivery(c *gin.Context) {
	var delivery struct {
		DeliveryType string `json:"DeliveryType" binding:"required"`
	}

	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.orderService.AddDelivery(delivery.DeliveryType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delivery added successfully"})

}

// @Summary Получение всех способов доставки
// @Description Возвращает список всех способов доставки из базы данных
// @Tags orders
// @Produce json
// @Success 200 {array} models.DeliveryType
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders/delivery [get]
func (ctrl *OrderController) GetAllDeliveries(c *gin.Context) {
	deliveries, err := ctrl.orderService.GetAllDeliveries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deliveries)
}

// @Summary Удаление способа доставки по id
// @Description Удаляет способ доставки из бд по id
// @Tags orders
// @Produce json
// @Param id path int true "ID способа доставки"
// @Success 200 {object} models.DeliveryType
// @Failure 400 {object} map[string]interface{}"Ошибка при получении способа доставки"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders/delivery/{id} [delete]
func (ctrl *OrderController) DeleteDelivery(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.orderService.DeleteDelivery(uri.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delivery deleted successfully"})
}

// @Summary Добавление способа оплаты
// @Description Добавляет способ оплаты в базу данных
// @Tags orders
// @Accept  json
// @Produce json
// @Param payment body string true "Название способа оплаты"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}"Ошибка при добавлении способа оплаты"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders/payment [post]
func (ctrl *OrderController) AddPayment(c *gin.Context) {
	var payment struct {
		PaymentType string `json:"paymentType" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.orderService.AddPayment(payment.PaymentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment added successfully"})
}

// @Summary Получение всех способов оплаты
// @Description Возвращает список всех способов оплаты из базы данных
// @Tags orders
// @Produce json
// @Success 200 {array} models.PaymentType
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders/payment [get]
func (ctrl *OrderController) GetAllPayments(c *gin.Context) {
	payments, err := ctrl.orderService.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

// @Summary Удаление способа оплаты по id
// @Description Удаляет способ оплаты из бд по id
// @Tags orders
// @Produce json
// @Param id path int true "ID способа оплаты"
// @Success 200 {object} models.PaymentType
// @Failure 400 {object} map[string]interface{}"Ошибка при получении способа оплаты"
// @Failure 500 {object} map[string]interface{}"Ошибка сервера"
// @Router /cake_api/orders/payment/{id} [delete]
func (ctrl *OrderController) DeletePayment(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.orderService.DeletePayment(uri.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}

// @Summary Добавить торты в заказ
// @Tags OrderContent
// @Accept json
// @Produce json
// @Param request body models.OrderContentPost true "Данные заказа"
// @Success 201
// @Failure 400 {object} map[string]interface{} "Ошибка запроса"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /cake_api/order_content [post]
func (ctrl *OrderController) CreateOrderContent(c *gin.Context) {
	var request models.OrderContentPost

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.orderService.CreateOrderContent(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Торты добавлены в заказ"})
}

// @Summary Получение содержимого заказа
// @Tags OrderContent
// @Produce json
// @Param orderID path int true "ID заказа"
// @Success 200 {array} models.OrderContent
// @Failure 400 {object} map[string]interface{} "Ошибка запроса"
// @Failure 404 {object} map[string]interface{} "Заказ не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /cake_api/order_content/{orderID} [get]
func (ctrl *OrderController) GetOrderContent(c *gin.Context) {
	var uri struct {
		OrderID int `uri:"orderID" binding:"required,min=1"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ctrl.orderService.GetOrderById(uri.OrderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	orderContents, err := ctrl.orderService.GetContentByOrderID(uri.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orderContents)
}
