package repository

import (
	"complaint-service/config"
	"complaint-service/internal/model"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(auth *model.Auth) error
	FindByUsername(username string) ([]model.Auth, error)
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (a *authRepository) Register(auth *model.Auth) error {
	return config.DB.Create(auth).Error
}

func (a *authRepository) FindByUsername(username string) ([]model.Auth, error) {
	var auths []model.Auth
	err := config.DB.Model(&model.Auth{}).Where("username = ?", username).Find(&auths).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrCustomerNotFound
	}
	return auths, err
}
