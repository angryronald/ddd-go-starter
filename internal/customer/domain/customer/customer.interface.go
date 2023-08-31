package customer

import (
	"context"

	"github.com/google/uuid"

	"github.com/angryronald/ddd-go-starter/internal/customer/domain/model"
)

type CustomerServiceInterface interface {
	GetCustomer(ctx context.Context, id uuid.UUID) (*model.CustomerDomainModel, error)
	GetCustomers(ctx context.Context, params []interface{}) ([]*model.CustomerDomainModel, error)
	AddCustomer(ctx context.Context, customer *model.CustomerDomainModel) (*model.CustomerDomainModel, error)
	UpdateCustomer(ctx context.Context, customer *model.CustomerDomainModel) (*model.CustomerDomainModel, error)
	RemoveCustomer(ctx context.Context, customer *model.CustomerDomainModel) (*model.CustomerDomainModel, error)
	MultipleAddOrUpdateCustomer(ctx context.Context, customers []*model.CustomerDomainModel) ([]*model.CustomerDomainModel, error)
}
