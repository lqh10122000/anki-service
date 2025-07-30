package service

import (
	"complaint-service/internal/model"
	"complaint-service/internal/repository"
)

type CustomerService interface {
	GetAll() ([]model.Customer, error)
	GetAllPaginated(offset, limit int, order string) ([]model.Customer, int64, error)
	Create(customer *model.Customer) error
	Delete(id uint) error
	Update(id uint, customer *model.Customer) error
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo}
}

func (s *customerService) GetAll() ([]model.Customer, error) {
	return s.repo.FindAll()
}

func (s *customerService) GetAllPaginated(offset, limit int, order string) ([]model.Customer, int64, error) {
	var customers []model.Customer
	err := s.repo.FindAllPaginated(&customers, offset, limit, order)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	total, err = s.repo.Count()
	if err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

func (s *customerService) Create(customer *model.Customer) error {
	return s.repo.Create(customer)
}

func (s *customerService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *customerService) Update(id uint, customer *model.Customer) error {
	return s.repo.Update(id, customer)
}
