package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository/model"
)

type AddressRedisRepository struct {
	redisClient *redis.Client
}

func (r *AddressRedisRepository) FindAll(ctx context.Context, params []interface{}) ([]*model.AddressRepositoryModel, error) {
	return nil, nil
}

func (r *AddressRedisRepository) FindBy(ctx context.Context, id interface{}) (*model.AddressRepositoryModel, error) {
	return nil, nil
}

func (r *AddressRedisRepository) Insert(ctx context.Context, address *model.AddressRepositoryModel) (*model.AddressRepositoryModel, error) {
	return nil, nil
}

func (r *AddressRedisRepository) Update(ctx context.Context, address *model.AddressRepositoryModel) (*model.AddressRepositoryModel, error) {
	return nil, nil
}

func (r *AddressRedisRepository) Delete(ctx context.Context, address *model.AddressRepositoryModel) error {
	return nil
}

func NewAddressRepository(redisClient *redis.Client) repository.AddressRepository {
	return &AddressRedisRepository{
		redisClient: redisClient,
	}
}
