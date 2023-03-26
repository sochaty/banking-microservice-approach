package service

import "banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaltCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaltCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaltCustomerService {
	return DefaltCustomerService{repo: repository}
}
