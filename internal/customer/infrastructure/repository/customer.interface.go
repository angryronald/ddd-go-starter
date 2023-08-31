package repository

import (
	"context"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository/model"
)

type CustomerRepository interface {
	FindAll(ctx context.Context, params []interface{}) ([]*model.CustomerRepositoryModel, error)
	FindBy(ctx context.Context, id interface{}) (*model.CustomerRepositoryModel, error)
	Insert(ctx context.Context, customer *model.CustomerRepositoryModel) (*model.CustomerRepositoryModel, error)
	Update(ctx context.Context, customer *model.CustomerRepositoryModel) (*model.CustomerRepositoryModel, error)
	Delete(ctx context.Context, customer *model.CustomerRepositoryModel) error
}
