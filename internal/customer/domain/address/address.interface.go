package address

import (
	"context"

	"github.com/google/uuid"

	"github.com/angryronald/ddd-go-starter/internal/customer/domain/model"
)

type AddressServiceInterface interface {
	GetAddress(ctx context.Context, ID uuid.UUID) (*model.AddressDomainModel, error)
	GetAddressByCustomerID(ctx context.Context, customerID uuid.UUID) (*model.AddressDomainModel, error)
	AddAddress(ctx context.Context, address *model.AddressDomainModel) (*model.AddressDomainModel, error)
	UpdateAddress(ctx context.Context, address *model.AddressDomainModel) (*model.AddressDomainModel, error)
	RemoveAddress(ctx context.Context, address *model.AddressDomainModel) (*model.AddressDomainModel, error)
	MultipleAddOrUpdateAddress(ctx context.Context, address []*model.AddressDomainModel) ([]*model.AddressDomainModel, error)
}
