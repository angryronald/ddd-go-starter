package command

import (
	"context"

	"github.com/google/uuid"

	"github.com/angryronald/ddd-go-starter/internal/customer/application/model"
)

type CustomerCommandInterface interface {
	AddAddress(ctx context.Context, customerID uuid.UUID, newAddress *model.AddressApplicationModel) (*model.AddressApplicationModel, error)
	AddCustomer(ctx context.Context, customer *model.CustomerApplicationModel) (*model.CustomerApplicationModel, error)
}
