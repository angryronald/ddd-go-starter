package postgre

import (
	"context"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository/model"
	"github.com/angryronald/ddd-go-starter/lib/helper/generic"
)

type AddressPostgreRepository struct {
	repository generic.GenericRepository
}

func (r *AddressPostgreRepository) FindAll(ctx context.Context, params []interface{}) ([]*model.AddressRepositoryModel, error) {
	return nil, nil
}

func (r *AddressPostgreRepository) FindBy(ctx context.Context, id interface{}) (*model.AddressRepositoryModel, error) {
	return nil, nil
}

func (r *AddressPostgreRepository) Insert(ctx context.Context, address *model.AddressRepositoryModel) (*model.AddressRepositoryModel, error) {
	return nil, nil
}

func (r *AddressPostgreRepository) Update(ctx context.Context, address *model.AddressRepositoryModel) (*model.AddressRepositoryModel, error) {
	return nil, nil
}

func (r *AddressPostgreRepository) Delete(ctx context.Context, address *model.AddressRepositoryModel) error {
	return nil
}

func NewAddressRepository(repo generic.GenericRepository) repository.AddressRepository {
	return &AddressPostgreRepository{
		repository: repo,
	}
}
