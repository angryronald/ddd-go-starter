package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository/model"
)

type CustomerRedisRepository struct {
	redisClient *redis.Client
}

func (r *CustomerRedisRepository) FindAll(ctx context.Context, params []interface{}) ([]*model.CustomerRepositoryModel, error) {
	return nil, nil
}

func (r *CustomerRedisRepository) FindBy(ctx context.Context, id interface{}) (*model.CustomerRepositoryModel, error) {
	return nil, nil
}

func (r *CustomerRedisRepository) Insert(ctx context.Context, customer *model.CustomerRepositoryModel) (*model.CustomerRepositoryModel, error) {
	return nil, nil
}

func (r *CustomerRedisRepository) Update(ctx context.Context, customer *model.CustomerRepositoryModel) (*model.CustomerRepositoryModel, error) {
	return nil, nil
}

func (r *CustomerRedisRepository) Delete(ctx context.Context, customer *model.CustomerRepositoryModel) error {
	return nil
}

func NewCustomerRepository(redisClient *redis.Client) repository.CustomerRepository {
	return &CustomerRedisRepository{
		redisClient: redisClient,
	}
}
