package services

import (
	"context"
	"errors"
	"time"

	"github.com/secmohammed/restaurant-management/utils"

	"github.com/secmohammed/restaurant-management/container"
	"github.com/secmohammed/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	MenuService interface {
		GetMenu(id int) (*models.Menu, error)
		GetMenus() ([]bson.M, error)
		CreateMenu(menu models.Menu) (*mongo.InsertOneResult, error)
		UpdateMenu(id int, menu models.Menu) (*mongo.UpdateResult, error)
		DeleteMenu(id int)
	}
	menuService struct {
		app *container.App
	}
)

func NewMenuService(app *container.App) MenuService {
	return &menuService{app}
}

func (u *menuService) GetMenu(id int) (*models.Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	menu := &models.Menu{}
	err := u.app.Database.OpenCollection("menu").FindOne(ctx, bson.M{"menu_id": id}).Decode(menu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (m *menuService) GetMenus() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := m.app.Database.OpenCollection("menu").Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var menus []bson.M
	if err := result.All(ctx, &menus); err != nil {
		return nil, err
	}
	return menus, nil
}

func (u *menuService) CreateMenu(menu models.Menu) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	menu.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	menu.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	menu.ID = primitive.NewObjectID()
	menu.MenuID = menu.ID.Hex()
	result, err := u.app.Database.OpenCollection("menu").InsertOne(ctx, menu)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *menuService) UpdateMenu(id int, menu models.Menu) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"menu_id": id}
	var updateObj primitive.D
	if menu.StartDate != nil && menu.EndDate != nil {
		if !utils.InTimeSpan(*menu.StartDate, *menu.EndDate, time.Now()) {
			return nil, errors.New("Start date must be before end date")
		}
		updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.StartDate})
		updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.EndDate})
		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}
		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
		}
		menu.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.UpdatedAt})
		upsert := true
		opt := options.UpdateOptions{Upsert: &upsert}
		result, err := u.app.Database.OpenCollection("menu").UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
		if err != nil {
			return nil, errors.New("Start date must be before end date")
		}
		return result, nil
	}
	return nil, errors.New("Body cannot be empty")
}

func (u *menuService) DeleteMenu(id int) {
}
