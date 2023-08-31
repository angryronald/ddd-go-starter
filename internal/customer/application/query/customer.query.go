package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/angryronald/ddd-go-starter/internal/customer/application/model"
	"github.com/angryronald/ddd-go-starter/internal/customer/domain/customer"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/external/geolocation"
)

type CustomerQuery struct {
	customerService    customer.CustomerServiceInterface
	geolocationService geolocation.GeolocationExternalServiceInterface
}

func (q *CustomerQuery) GetCustomer(ctx context.Context, customerID uuid.UUID) (*model.CustomerApplicationModel, error) {
	return nil, nil
}

func (q *CustomerQuery) GetCustomers(ctx context.Context) ([]*model.CustomerApplicationModel, error) {
	return nil, nil
}

func NewCustomerQuery(
	customerService customer.CustomerServiceInterface,
	geolocationService geolocation.GeolocationExternalServiceInterface,
) CustomerQueryInterface {
	return &CustomerQuery{
		customerService:    customerService,
		geolocationService: geolocationService,
	}
}
