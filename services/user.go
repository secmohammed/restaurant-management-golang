package services

import "github.com/secmohammed/restaurant-management/container"

type (
	UserService interface {
		GetUser(id int)
		GetUsers()
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

func (u *userService) GetUser(id int) {
}

func (u *userService) GetUsers() {
}

func (u *userService) CreateUser() {
}

func (u *userService) UpdateUser(id int) {
}

func (u *userService) DeleteUser(id int) {
}
