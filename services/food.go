package services

import "github.com/secmohammed/restaurant-management/container"

type (
	FoodService interface {
		GetFood(id int)
		GetFoods()
		CreateFood()
		UpdateFood(id int)
		DeleteFood(id int)
	}
	foodService struct {
		app *container.App
	}
)

func NewFoodService(app *container.App) FoodService {
	return &foodService{app}
}

func (u *foodService) GetFood(id int) {
}

func (u *foodService) GetFoods() {
}

func (u *foodService) CreateFood() {
}

func (u *foodService) UpdateFood(id int) {
}

func (u *foodService) DeleteFood(id int) {
}
