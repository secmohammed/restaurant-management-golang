package services

import "github.com/secmohammed/restaurant-management/container"

type (
	OrderService interface {
		GetOrder(id int)
		GetOrders()
		CreateOrder()
		UpdateOrder(id int)
		DeleteOrder(id int)
	}
	orderService struct {
		app *container.App
	}
)

func NewOrderService(app *container.App) OrderService {
	return &orderService{app}
}

func (u *orderService) GetOrder(id int) {
}

func (u *orderService) GetOrders() {
}

func (u *orderService) CreateOrder() {
}

func (u *orderService) UpdateOrder(id int) {
}

func (u *orderService) DeleteOrder(id int) {
}
