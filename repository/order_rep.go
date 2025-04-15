package repository

import (
	"kuromi_cakes/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderById(id int) (models.Order, error)
	CreateOrder(order models.OrderPost) (models.Order, error)
	DeleteOrder(id int) error
	AddDelivery(deliveryType string) error
	AddPayment(paymentType string) error
	GetAllDeliveries() ([]models.DeliveryType, error)
	GetAllPayments() ([]models.PaymentType, error)
	DeleteDelivery(id int) error
	DeletePayment(id int) error
	CreateOrderContent(orderContents []models.OrderContent) error
	GetContentByOrderID(orderID int) ([]models.OrderContent, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("DeliveryType").Preload("PaymentType").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetOrderById(id int) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("DeliveryType").Preload("PaymentType").Where("id = ?", id).First(&order).Error
	return order, err
}

func (r *orderRepository) CreateOrder(order models.OrderPost) (models.Order, error) {
	neworder := models.Order{
		DeliveryTypeID: order.DeliveryTypeID,
		PaymentTypeID:  order.PaymentTypeID,
		Name:           order.Name,
		Address:        order.Address,
		Phone:          order.Phone,
		Email:          order.Email,
		TotalPrice:     order.TotalPrice,
	}

	err := r.db.Create(&neworder).Error
	return neworder, err
}

func (r *orderRepository) DeleteOrder(id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Order{}).Error
	return err
}

func (r *orderRepository) AddDelivery(deliveryType string) error {
	err := r.db.Create(&models.DeliveryType{Name: deliveryType}).Error
	return err
}

func (r *orderRepository) AddPayment(paymentType string) error {
	err := r.db.Create(&models.PaymentType{Name: paymentType}).Error
	return err
}

func (r *orderRepository) GetAllDeliveries() ([]models.DeliveryType, error) {
	var deliveries []models.DeliveryType
	err := r.db.Find(&deliveries).Error
	return deliveries, err
}

func (r *orderRepository) GetAllPayments() ([]models.PaymentType, error) {
	var payments []models.PaymentType
	err := r.db.Find(&payments).Error
	return payments, err
}

func (r *orderRepository) DeleteDelivery(id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.DeliveryType{}).Error
	return err
}

func (r *orderRepository) DeletePayment(id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.PaymentType{}).Error
	return err
}

func (r *orderRepository) CreateOrderContent(orderContents []models.OrderContent) error {
	err := r.db.Create(&orderContents).Error
	return err
}

func (r *orderRepository) GetContentByOrderID(orderID int) ([]models.OrderContent, error) {
	var orderContents []models.OrderContent
	err := r.db.Preload("Cake").Where("order_id = ?", orderID).Find(&orderContents).Error
	return orderContents, err
}
