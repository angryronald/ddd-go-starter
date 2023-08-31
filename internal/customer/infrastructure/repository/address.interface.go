package repository

import (
	"context"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository/model"
)

type AddressRepository interface {
	FindAll(ctx context.Context, params []interface{}) ([]*model.AddressRepositoryModel, error)
	FindBy(ctx context.Context, id interface{}) (*model.AddressRepositoryModel, error)
	Insert(ctx context.Context, address *model.AddressRepositoryModel) (*model.AddressRepositoryModel, error)
	Update(ctx context.Context, address *model.AddressRepositoryModel) (*model.AddressRepositoryModel, error)
	Delete(ctx context.Context, address *model.AddressRepositoryModel) error
}
