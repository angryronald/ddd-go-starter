package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/angryronald/ddd-go-starter/lib/helper/generic"
)

type GenericRepository struct {
	db *mongo.Client
}

func (r *GenericRepository) FindOne(ctx context.Context, id interface{}) (interface{}, error) {
	return nil, nil
}

func (r *GenericRepository) FindAll(ctx context.Context, params []interface{}) ([]interface{}, error) {
	return nil, nil
}

func (r *GenericRepository) BulkUpsert(ctx context.Context, object []interface{}) ([]interface{}, error) {
	return nil, nil
}

func (r *GenericRepository) Upsert(ctx context.Context, object interface{}) (interface{}, error) {
	return nil, nil
}

func (r *GenericRepository) Insert(ctx context.Context, object interface{}) (interface{}, error) {
	return nil, nil
}

func (r *GenericRepository) Update(ctx context.Context, object interface{}) (interface{}, error) {
	return nil, nil
}

func (r *GenericRepository) Delete(ctx context.Context, object interface{}) error {
	return nil
}

func NewGenericRepository(db *mongo.Client) generic.GenericRepository {
	return &GenericRepository{
		db: db,
	}
}
