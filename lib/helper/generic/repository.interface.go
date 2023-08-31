package generic

import "context"

type GenericRepository interface {
	FindOne(ctx context.Context, id interface{}) (interface{}, error)
	FindAll(ctx context.Context, params []interface{}) ([]interface{}, error)
	BulkUpsert(ctx context.Context, object []interface{}) ([]interface{}, error)
	Upsert(ctx context.Context, object interface{}) (interface{}, error)
	Insert(ctx context.Context, object interface{}) (interface{}, error)
	Update(ctx context.Context, object interface{}) (interface{}, error)
	Delete(ctx context.Context, object interface{}) error
}
