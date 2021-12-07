package services

import (
	"context"
	"time"

	"github.com/secmohammed/restaurant-management/container"
	"github.com/secmohammed/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	OrderItemService interface {
		GetOrderItem(id int) (*models.OrderItem, error)
		GetOrderItems() ([]bson.M, error)
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

func (o *orderItemService) GetOrderItem(id int) (*models.OrderItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	order_item := &models.OrderItem{}
	err := o.app.Database.OpenCollection("order_item").FindOne(ctx, bson.M{"order_item_id": id}).Decode(order_item)
	if err != nil {
		return nil, err
	}
	return order_item, nil
}

func (o *orderItemService) GetOrderItems() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := o.app.Database.OpenCollection("order_item").Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var order_items []bson.M
	if err := result.All(ctx, &order_items); err != nil {
		return nil, err
	}
	return order_items, nil
}

func (u *orderItemService) CreateOrderItem() {
}

func (u *orderItemService) UpdateOrderItem(id int) {
}

func (u *orderItemService) DeleteOrderItem(id int) {
}
