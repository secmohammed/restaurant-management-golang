package services

import (
	"context"
	"time"

	"github.com/secmohammed/restaurant-management/models"

	"github.com/secmohammed/restaurant-management/container"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	UserService interface {
		GetUser(id int) (*models.User, error)
		GetUsers() ([]bson.M, error)
		CreateUser()
		UpdateUser(id int)
		DeleteUser(id int)
	}
	userService struct {
		app *container.App
	}
)

func NewUserService(app *container.App) UserService {
	return &userService{app}
}

func (u *userService) GetUser(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	user := &models.User{}
	err := u.app.Database.OpenCollection("user").FindOne(ctx, bson.M{"user_id": id}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) GetUsers() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := u.app.Database.OpenCollection("user").Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var users []bson.M
	if err := result.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userService) CreateUser() {
}

func (u *userService) UpdateUser(id int) {
}

func (u *userService) DeleteUser(id int) {
}
