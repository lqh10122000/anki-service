package repository

import (
	"complaint-service/config"
	"complaint-service/internal/model"
)

type CustomerRepository interface {
	FindAll() ([]model.Customer, error)
	FindAllPaginated(customers *[]model.Customer, offset int, limit int, order string) error
	Create(customer *model.Customer) error
	Delete(id uint) error
	Update(id uint, customer *model.Customer) error
	Count() (int64, error)
}

type customerRepository struct{}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}

func (r *customerRepository) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	result := config.DB.Find(&customers)
	return customers, result.Error
}

func (r *customerRepository) Create(customer *model.Customer) error {
	return config.DB.Create(customer).Error
}

func (r *customerRepository) Delete(id uint) error {
	return config.DB.Delete(&model.Customer{}, id).Error
}

func (r *customerRepository) Update(id uint, customer *model.Customer) error {
	return config.DB.Model(&model.Customer{}).Where("id = ?", id).Updates(customer).Error
}

func (r *customerRepository) FindAllPaginated(customers *[]model.Customer, offset int, limit int, order string) error {
	return config.DB.Order(order).Limit(limit).Offset(offset).Find(customers).Error
}

func (r *customerRepository) Count() (int64, error) {
	var count int64
	err := config.DB.Model(&model.Customer{}).Count(&count).Error
	return count, err
}
