package postgre

import (
	"context"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository/model"
	"github.com/angryronald/ddd-go-starter/lib/helper/generic"
)

type CustomerPostgreRepository struct {
	repository generic.GenericRepository
}

func (r *CustomerPostgreRepository) FindAll(ctx context.Context, params []interface{}) ([]*model.CustomerRepositoryModel, error) {
	return nil, nil
}

func (r *CustomerPostgreRepository) FindBy(ctx context.Context, id interface{}) (*model.CustomerRepositoryModel, error) {
	return nil, nil
}

func (r *CustomerPostgreRepository) Insert(ctx context.Context, customer *model.CustomerRepositoryModel) (*model.CustomerRepositoryModel, error) {
	return nil, nil
}

func (r *CustomerPostgreRepository) Update(ctx context.Context, customer *model.CustomerRepositoryModel) (*model.CustomerRepositoryModel, error) {
	return nil, nil
}

func (r *CustomerPostgreRepository) Delete(ctx context.Context, customer *model.CustomerRepositoryModel) error {
	return nil
}

func NewCustomerRepository(repository generic.GenericRepository) repository.CustomerRepository {
	return &CustomerPostgreRepository{
		repository: repository,
	}
}
