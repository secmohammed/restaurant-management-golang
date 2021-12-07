package services

import (
	"context"
	"errors"
	"time"

	"github.com/secmohammed/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/secmohammed/restaurant-management/container"
)

type (
	OrderService interface {
		GetOrder(id int) (*models.Order, error)
		GetOrders() ([]bson.M, error)
		CreateOrder(order models.Order) (*mongo.InsertOneResult, error)
		UpdateOrder(id int, order models.Order) (*mongo.UpdateResult, error)
		DeleteOrder(id int)
	}
	orderService struct {
		app *container.App
	}
)

func NewOrderService(app *container.App) OrderService {
	return &orderService{app}
}

func (o *orderService) GetOrder(id int) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	order := &models.Order{}
	err := o.app.Database.OpenCollection("order").FindOne(ctx, bson.M{"order_id": id}).Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *orderService) GetOrders() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := o.app.Database.OpenCollection("order").Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var orders []bson.M
	if err := result.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderService) CreateOrder(order models.Order) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	order.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.OrderID = order.ID.Hex()
	result, err := o.app.Database.OpenCollection("order").InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *orderService) UpdateOrder(id int, order models.Order) (*mongo.UpdateResult, error) {
	table := models.Table{}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.M{"order_id": id}
	var updateObj primitive.D
	if order.TableID != nil {
		err := o.app.Database.OpenCollection("table").FindOne(ctx, bson.M{"table_id": *order.TableID}).Decode(&table)
		if err != nil {
			return nil, errors.New("table not found")
		}
		updateObj = append(updateObj, bson.E{Key: "table_id", Value: order.TableID})
	}
	order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: order.UpdatedAt})
	upsert := true
	opt := options.UpdateOptions{Upsert: &upsert}
	result, err := o.app.Database.OpenCollection("order").UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
	if err != nil {
		return nil, errors.New("failed to update order")
	}
	return result, nil
}

func (u *orderService) DeleteOrder(id int) {
}
