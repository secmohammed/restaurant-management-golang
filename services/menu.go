package services

import "github.com/secmohammed/restaurant-management/container"

type (
	MenuService interface {
		GetMenu(id int)
		GetMenus()
		CreateMenu()
		UpdateMenu(id int)
		DeleteMenu(id int)
	}
	menuService struct {
		app *container.App
	}
)

func NewMenuService(app *container.App) MenuService {
	return &menuService{app}
}

func (u *menuService) GetMenu(id int) {
}

func (u *menuService) GetMenus() {
}

func (u *menuService) CreateMenu() {
}

func (u *menuService) UpdateMenu(id int) {
}

func (u *menuService) DeleteMenu(id int) {
}
