package services

import "github.com/secmohammed/restaurant-management/container"

type (
	TableService interface {
		GetTable(id int)
		GetTables()
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

func (u *tableService) GetTable(id int) {
}

func (u *tableService) GetTables() {
}

func (u *tableService) CreateTable() {
}

func (u *tableService) UpdateTable(id int) {
}

func (u *tableService) DeleteTable(id int) {
}
