package repository

import (
	"complaint-service/config"
	"complaint-service/internal/model"
)

type DeskRepository interface {
	FindAll() ([]model.Desks, error)
	FindByID(id uint) (model.Desks, error)
	SaveDesk(desk string) error
	FindAllByName(name string) ([]model.Desks, error)
}

type deskRepository struct{}

func NewDeskRepository() DeskRepository {
	return &deskRepository{}
}

func (r *deskRepository) FindAll() ([]model.Desks, error) {
	var desks []model.Desks
	result := config.DB.Find(&desks)
	return desks, result.Error
}

func (r *deskRepository) FindAllByName(name string) ([]model.Desks, error) {
	var desks []model.Desks
	result := config.DB.Where("name = ?", name).Find(&desks)
	if result.Error != nil {
		return nil, result.Error
	}
	return desks, nil
}

func (r *deskRepository) FindByID(id uint) (model.Desks, error) {
	var desk model.Desks
	result := config.DB.First(&desk, id)
	if result.Error != nil {
		return desk, result.Error
	}
	return desk, nil
}

func (r *deskRepository) SaveDesk(desk string) error {
	newDesk := model.Desks{Name: desk}
	result := config.DB.Create(&newDesk)
	return result.Error
}
