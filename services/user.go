package services

import (
	"context"
	"errors"
	"time"

	"github.com/secmohammed/restaurant-management/helpers"
	"golang.org/x/crypto/bcrypt"

	"github.com/secmohammed/restaurant-management/models"

	"github.com/secmohammed/restaurant-management/container"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	UserService interface {
		GetUser(id int) (*models.User, error)
		GetUsers(limit, page int) (bson.M, error)
		Signup(user models.User) (*mongo.InsertOneResult, error)
		Login(user models.User) (*models.User, error)
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

func (u *userService) GetUsers(limit, page int) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
	projectStage := bson.D{
		{
			Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "total_count", Value: 1},
				{Key: "user_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$user_items", page, limit}}}},
			},
		},
	}
	result, err := u.app.Database.OpenCollection("user").Aggregate(ctx, mongo.Pipeline{matchStage, projectStage})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var users []bson.M
	if err := result.All(ctx, &users); err != nil {
		return nil, err
	}
	return users[0], nil
}

func (u *userService) Signup(user models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := u.app.Database.OpenCollection("user").CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("user email already exists")
	}
	count, err = u.app.Database.OpenCollection("user").CountDocuments(ctx, bson.M{"phone": user.Phone})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("user phone already exists")
	}
	password, err := HashPassowrd(*user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = &password
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()
	token, refreshToken, _ := helpers.GenerateAuthTokensForUser(user)
	user.Token = &token
	user.RefreshToken = &refreshToken
	result, err := u.app.Database.OpenCollection("user").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) Login(user models.User) (*models.User, error) {
	foundUser := &models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := u.app.Database.OpenCollection("user").FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		return nil, err
	}
	ok, err := VerifyPassword(*user.Password, *foundUser.Password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("invalid password")
	}
	token, refreshToken, _ := helpers.GenerateAuthTokensForUser(user)
	helpers.UpdateAuthTokensForUser(token, refreshToken, *foundUser)
	return foundUser, nil
}

func VerifyPassword(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func HashPassowrd(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (u *userService) UpdateUser(id int) {
}

func (u *userService) DeleteUser(id int) {
}
