package services

import (
	"context"
	"time"

	"github.com/secmohammed/restaurant-management/models"

	"github.com/secmohammed/restaurant-management/container"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	TableService interface {
		GetTable(id int) (*models.Table, error)
		GetTables() ([]bson.M, error)
		CreateTable()
		UpdateTable(id int)
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

func (u *tableService) CreateTable() {
}

func (u *tableService) UpdateTable(id int) {
}

func (u *tableService) DeleteTable(id int) {
}
