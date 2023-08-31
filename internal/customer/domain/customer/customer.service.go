package customer

import (
	"context"

	"github.com/google/uuid"

	"github.com/angryronald/ddd-go-starter/internal/customer/domain/model"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/event/publisher"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository"
)

type CustomerService struct {
	publisher  publisher.Publisher
	repository repository.CustomerRepository
}

func (s *CustomerService) GetCustomer(ctx context.Context, id uuid.UUID) (*model.CustomerDomainModel, error) {
	return nil, nil
}

func (s *CustomerService) GetCustomers(ctx context.Context, params []interface{}) ([]*model.CustomerDomainModel, error) {
	return nil, nil
}

func (s *CustomerService) AddCustomer(ctx context.Context, customer *model.CustomerDomainModel) (*model.CustomerDomainModel, error) {
	return nil, nil
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, customer *model.CustomerDomainModel) (*model.CustomerDomainModel, error) {
	return nil, nil
}

func (s *CustomerService) RemoveCustomer(ctx context.Context, customer *model.CustomerDomainModel) (*model.CustomerDomainModel, error) {
	return nil, nil
}

func (s *CustomerService) MultipleAddOrUpdateCustomer(ctx context.Context, customers []*model.CustomerDomainModel) ([]*model.CustomerDomainModel, error) {
	return nil, nil
}

func NewCustomerService(
	publisher publisher.Publisher,
	repository repository.CustomerRepository,
) CustomerServiceInterface {
	return &CustomerService{
		publisher:  publisher,
		repository: repository,
	}
}
