package command

import (
	"context"

	"github.com/google/uuid"

	"github.com/angryronald/ddd-go-starter/internal/customer/application/model"
	"github.com/angryronald/ddd-go-starter/internal/customer/domain/customer"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/external/geolocation"
)

type CustomerCommand struct {
	customerService    customer.CustomerServiceInterface
	geolocationService geolocation.GeolocationExternalServiceInterface
}

func (c *CustomerCommand) AddAddress(ctx context.Context, customerID uuid.UUID, newAddress *model.AddressApplicationModel) (*model.AddressApplicationModel, error) {
	return nil, nil
}

func (c *CustomerCommand) AddCustomer(ctx context.Context, customer *model.CustomerApplicationModel) (*model.CustomerApplicationModel, error) {
	return nil, nil
}

func NewCustomerCommand(
	customerService customer.CustomerServiceInterface,
	geolocationService geolocation.GeolocationExternalServiceInterface,
) CustomerCommandInterface {
	return &CustomerCommand{
		customerService:    customerService,
		geolocationService: geolocationService,
	}
}
