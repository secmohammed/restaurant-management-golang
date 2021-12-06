package services

import "github.com/secmohammed/restaurant-management/container"

type (
	OrderItemService interface {
		GetOrderItem(id int)
		GetOrderItems()
		CreateOrderItem()
		UpdateOrderItem(id int)
		DeleteOrderItem(id int)
	}
	orderItemService struct {
		app *container.App
	}
)

func NewOrderItemService(app *container.App) OrderItemService {
	return &orderItemService{app}
}

func (u *orderItemService) GetOrderItem(id int) {
}

func (u *orderItemService) GetOrderItems() {
}

func (u *orderItemService) CreateOrderItem() {
}

func (u *orderItemService) UpdateOrderItem(id int) {
}

func (u *orderItemService) DeleteOrderItem(id int) {
}
