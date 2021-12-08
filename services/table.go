package services

import (
	"context"
	"errors"
	"time"

	"github.com/secmohammed/restaurant-management/models"

	"github.com/secmohammed/restaurant-management/container"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	TableService interface {
		GetTable(id int) (*models.Table, error)
		GetTables() ([]bson.M, error)
		CreateTable(table models.Table) (*mongo.InsertOneResult, error)
		UpdateTable(id int, table models.Table) (*mongo.UpdateResult, error)
		DeleteTable(id int)
	}
	tableService struct {
		app *container.App
	}
)

func NewTableService(app *container.App) TableService {
	return &tableService{app}
}

func (u *tableService) GetTable(id int) (*models.Table, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	table := &models.Table{}
	err := u.app.Database.OpenCollection("table").FindOne(ctx, bson.M{"table_id": id}).Decode(table)
	if err != nil {
		return nil, err
	}
	return table, nil
}

func (t *tableService) GetTables() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := t.app.Database.OpenCollection("table").Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var tables []bson.M
	if err := result.All(ctx, &tables); err != nil {
		return nil, err
	}
	return tables, nil
}

func (t *tableService) CreateTable(table models.Table) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.app.Database.OpenCollection("table").InsertOne(ctx, table)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *tableService) UpdateTable(id int, table models.Table) (*mongo.UpdateResult, error) {
	var updateObj primitive.D
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if table.Capacity != nil {
		updateObj = append(updateObj, bson.E{Key: "capacity", Value: *table.Capacity})
	}
	if table.Number != nil {
		updateObj = append(updateObj, bson.E{Key: "number", Value: *table.Number})
	}
	table.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: table.UpdatedAt})
	upsert := true
	opt := options.UpdateOptions{Upsert: &upsert}
	filter := bson.M{"table_id": id}

	result, err := t.app.Database.OpenCollection("table").UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
	if err != nil {
		return nil, errors.New("failed to update food item")
	}
	return result, nil
}

func (u *tableService) DeleteTable(id int) {
}
