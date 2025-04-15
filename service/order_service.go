package service

import (
	"kuromi_cakes/models"
	"kuromi_cakes/repository"
)

type OrderService interface {
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
	CreateOrderContent(request models.OrderContentPost) error
	GetContentByOrderID(orderID int) ([]models.OrderContent, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
	return s.orderRepository.GetAllOrders()
}

func (s *orderService) GetOrderById(id int) (models.Order, error) {
	return s.orderRepository.GetOrderById(id)
}

func (s *orderService) CreateOrder(order models.OrderPost) (models.Order, error) {
	return s.orderRepository.CreateOrder(order)
}

func (s *orderService) DeleteOrder(id int) error {
	return s.orderRepository.DeleteOrder(id)
}

func (s *orderService) AddDelivery(deliveryType string) error {
	return s.orderRepository.AddDelivery(deliveryType)
}

func (s *orderService) AddPayment(paymentType string) error {
	return s.orderRepository.AddPayment(paymentType)
}

func (s *orderService) GetAllDeliveries() ([]models.DeliveryType, error) {
	return s.orderRepository.GetAllDeliveries()
}

func (s *orderService) GetAllPayments() ([]models.PaymentType, error) {
	return s.orderRepository.GetAllPayments()
}

func (s *orderService) DeleteDelivery(id int) error {
	return s.orderRepository.DeleteDelivery(id)
}

func (s *orderService) DeletePayment(id int) error {
	return s.orderRepository.DeletePayment(id)
}

func (s *orderService) CreateOrderContent(request models.OrderContentPost) error {
	var orderContents []models.OrderContent

	for _, item := range request.Items {
		orderContents = append(orderContents, models.OrderContent{
			OrderID:  request.OrderID,
			CakeID:   item.CakeID,
			Quantity: item.Quantity,
		})
	}

	return s.orderRepository.CreateOrderContent(orderContents)
}

func (s *orderService) GetContentByOrderID(orderID int) ([]models.OrderContent, error) {
	return s.orderRepository.GetContentByOrderID(orderID)
}
