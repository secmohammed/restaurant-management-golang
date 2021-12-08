package services

import (
	"context"
	"errors"
	"time"

	"github.com/secmohammed/restaurant-management/container"
	"github.com/secmohammed/restaurant-management/models"
	"github.com/secmohammed/restaurant-management/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	OrderItemService interface {
		GetOrderItem(id int) (*models.OrderItem, error)
		CreateOrderItem(item models.OrderItemPack, orderID *string) (*mongo.InsertManyResult, error)
		UpdateOrderItem(id int, item models.OrderItem) (*mongo.UpdateResult, error)
		DeleteOrderItem(id int)
		ItemsByOrder(id int) ([]bson.M, error)
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

func (o *orderItemService) ItemsByOrder(id int) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	matchStage := bson.D{{"$match", bson.D{{"order_id", id}}}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "food"}, {"localField", "food_id"}, {"foreignField", "food_id"}, {"as", "food"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$food"}, {"preserveNullAndEmptyArrays", true}}}}
	lookupOrderStage := bson.D{{"$lookup", bson.D{{"from", "order"}, {"localField", "order_id"}, {"foreignField", "order_id"}, {"as", "order"}}}}
	unwindOrderStage := bson.D{{"$unwind", bson.D{{"path", "$order"}, {"preserveNullAndEmptyArrays", true}}}}
	lookupTableStage := bson.D{{"$lookup", bson.D{{"from", "table"}, {"localField", "order.table_id"}, {"foreignField", "table_id"}, {"as", "table"}}}}
	unwindTableStage := bson.D{{"$unwind", bson.D{{"path", "$table"}, {"preserveNullAndEmptyArrays", true}}}}
	projectStage := bson.D{
		{
			"$project", bson.D{
				{Key: "id", Value: 0},
				{Key: "amount", Value: "$food.price"},
				{Key: "total_count", Value: 1},
				{Key: "food_name", Value: "$food.name"},
				{Key: "food_image", Value: "$food.image"},
				{Key: "table_number", Value: "$table.number"},
				{Key: "table_id", Value: "$table.table_id"},
				{Key: "order_id", Value: "$order.order_id"},
				{Key: "price", Value: "$food.price"},
				{Key: "quantity", Value: 1},
			},
		},
	}
	groupStage := bson.D{
		{"$group", bson.D{{"_id", bson.D{{"order_id", "$order_id"}, {"table_id", "$table_id"}, {"table_number", "$table_number"}}}, {"payment_due", bson.D{{"$sum", "$amount"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"order_items", bson.D{{Key: "$push", Value: "$$ROOT"}}}}},
	}

	projectStage2 := bson.D{
		{
			"$project", bson.D{
				{"id", 0},
				{Key: "payment_due", Value: 1},
				{Key: "total_count", Value: 1},
				{Key: "table_number", Value: "$_id.table_number"},
				{Key: "order_items", Value: 1},
			},
		},
	}
	result, err := o.app.Database.OpenCollection("order_item").Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage, lookupOrderStage, unwindOrderStage, lookupTableStage, unwindTableStage, projectStage, groupStage, projectStage2})
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

func (o *orderItemService) CreateOrderItem(pack models.OrderItemPack, orderID *string) (*mongo.InsertManyResult, error) {
	orderItems := []interface{}{}
	for _, item := range pack.Items {
		item.OrderID = orderID
		err := o.app.Validator.Struct(item)
		if err != nil {
			return nil, err
		}
		item.ID = primitive.NewObjectID()
		item.OrderItemID = item.ID.Hex()
		item.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		item.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		num := utils.ToFixed(*item.Price, 2)
		item.Price = &num
		orderItems = append(orderItems, item)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	items, err := o.app.Database.OpenCollection("order_item").InsertMany(ctx, orderItems)
	defer cancel()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (o *orderItemService) UpdateOrderItem(id int, order models.OrderItem) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.M{"order_item_id": id}

	var updateObj primitive.D
	if order.Price != nil {
		updateObj = append(updateObj, primitive.E{Key: "$set", Value: bson.M{"price": order.Price}})
	}
	if order.Size != nil {
		updateObj = append(updateObj, primitive.E{Key: "$set", Value: bson.M{"size": order.Size}})
	}
	if order.FoodID != nil {
		updateObj = append(updateObj, primitive.E{Key: "$set", Value: bson.M{"food_id": order.FoodID}})
	}
	order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: order.UpdatedAt})

	upsert := true
	opt := options.UpdateOptions{Upsert: &upsert}
	result, err := o.app.Database.OpenCollection("invoice").UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
	if err != nil {
		return nil, errors.New("failed to order item invoice")
	}
	return result, nil
}

func (u *orderItemService) DeleteOrderItem(id int) {
}
