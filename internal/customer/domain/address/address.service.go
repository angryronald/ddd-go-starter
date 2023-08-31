package address

import (
	"context"

	"github.com/google/uuid"

	"github.com/angryronald/ddd-go-starter/internal/customer/domain/model"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/event/publisher"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository"
)

type AddressService struct {
	repository repository.AddressRepository
	publisher  publisher.Publisher
}

func (s *AddressService) GetAddress(ctx context.Context, ID uuid.UUID) (*model.AddressDomainModel, error) {
	return nil, nil
}

func (s *AddressService) GetAddressByCustomerID(ctx context.Context, customerID uuid.UUID) (*model.AddressDomainModel, error) {
	return nil, nil
}

func (s *AddressService) AddAddress(ctx context.Context, address *model.AddressDomainModel) (*model.AddressDomainModel, error) {
	return nil, nil
}

func (s *AddressService) UpdateAddress(ctx context.Context, address *model.AddressDomainModel) (*model.AddressDomainModel, error) {
	return nil, nil
}

func (s *AddressService) RemoveAddress(ctx context.Context, address *model.AddressDomainModel) (*model.AddressDomainModel, error) {
	return nil, nil
}

func (s *AddressService) MultipleAddOrUpdateAddress(ctx context.Context, address []*model.AddressDomainModel) ([]*model.AddressDomainModel, error) {
	return nil, nil
}

func NewAddressService(
	repository repository.AddressRepository,
	publisher publisher.Publisher,
) AddressServiceInterface {
	return &AddressService{
		repository: repository,
		publisher:  publisher,
	}
}
