package publisher

import "context"

type Publisher interface {
	Publish(ctx context.Context, topic string, idempotentId string, event []byte) error
}
