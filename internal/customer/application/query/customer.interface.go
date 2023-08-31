package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/angryronald/ddd-go-starter/internal/customer/application/model"
)

type CustomerQueryInterface interface {
	GetCustomer(ctx context.Context, customerID uuid.UUID) (*model.CustomerApplicationModel, error)
	GetCustomers(ctx context.Context) ([]*model.CustomerApplicationModel, error)
}
