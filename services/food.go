package services

import (
	"context"
	"errors"
	"time"

	"github.com/secmohammed/restaurant-management/utils"

	"github.com/secmohammed/restaurant-management/models"

	"github.com/secmohammed/restaurant-management/container"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	FoodService interface {
		GetFood(id int) (*models.Food, error)
		GetFoods(page, limit int) ([]bson.M, error)
		CreateFood(food models.Food) (*mongo.InsertOneResult, error)
		UpdateFood(id int, food models.Food) (*mongo.UpdateResult, error)
		DeleteFood(id int)
	}
	foodService struct {
		app *container.App
	}
)

func NewFoodService(app *container.App) FoodService {
	return &foodService{app}
}

func (u *foodService) GetFood(id int) (*models.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	food := &models.Food{}
	err := u.app.Database.OpenCollection("food").FindOne(ctx, bson.M{"food_id": id}).Decode(food)
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (f *foodService) GetFoods(page, limit int) ([]bson.M, error) {
	offset := (page - 1) * limit
	matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
	groupStage := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "null"}, {Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}}, {Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}}}}}
	projectionStage := bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}, {Key: "total_count", Value: 1}, {Key: "foods", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", offset, limit}}}}}}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := f.app.Database.OpenCollection("food").Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectionStage})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var foods []bson.M
	if err := result.All(ctx, &foods); err != nil {
		return nil, err
	}
	return foods, nil
}

func (u *foodService) CreateFood(food models.Food) (*mongo.InsertOneResult, error) {
	menu := models.Menu{}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := u.app.Database.OpenCollection("menu").FindOne(ctx, bson.M{"menu_id": food.MenuID}).Decode(&menu)
	if err != nil {
		return nil, err
	}
	food.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.ID = primitive.NewObjectID()
	food.FoodID = food.ID.Hex()
	num := utils.ToFixed(*food.Price, 2)
	food.Price = &num
	result, err := u.app.Database.OpenCollection("food").InsertOne(ctx, food)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (f *foodService) UpdateFood(id int, food models.Food) (*mongo.UpdateResult, error) {
	var updateObj primitive.D
	menu := models.Menu{}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if food.Name != nil {
		updateObj = append(updateObj, bson.E{Key: "name", Value: *food.Name})
	}
	if food.Price != nil {
		updateObj = append(updateObj, bson.E{Key: "price", Value: *food.Price})
	}
	if food.Image != nil {
		updateObj = append(updateObj, bson.E{Key: "image", Value: *food.Image})
	}
	if food.MenuID != nil {
		err := f.app.Database.OpenCollection("menu").FindOne(ctx, bson.M{"menu_id": *food.MenuID}).Decode(&menu)
		if err != nil {
			return nil, errors.New("failed to find menu with id " + *food.MenuID)
		}
		updateObj = append(updateObj, bson.E{Key: "menu_id", Value: *food.MenuID})
	}
	food.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: food.UpdatedAt})
	upsert := true
	opt := options.UpdateOptions{Upsert: &upsert}
	filter := bson.M{"food_id": id}

	result, err := f.app.Database.OpenCollection("menu").UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
	if err != nil {
		return nil, errors.New("failed to update food item")
	}
	return result, nil
}

func (u *foodService) DeleteFood(id int) {
}
