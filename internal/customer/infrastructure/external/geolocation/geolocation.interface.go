package geolocation

import (
	"context"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/external/geolocation/model"
)

type GeolocationExternalServiceInterface interface {
	GetCoordinate(ctx context.Context, request model.AddressExternalRequestModel) (*model.GeolocationExternalModel, error)
}
