package http

import (
	"context"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/external/geolocation"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/external/geolocation/model"
)

type GeolocationHttpService struct {
}

func (s *GeolocationHttpService) GetCoordinate(ctx context.Context, request model.AddressExternalRequestModel) (*model.GeolocationExternalModel, error) {
	return nil, nil
}

func NewGeolocationService() geolocation.GeolocationExternalServiceInterface {
	return &GeolocationHttpService{}
}
