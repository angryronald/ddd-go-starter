package kafka

import (
	"context"

	"gocloud.dev/pubsub"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/event/publisher"
)

type KafkaPublisher struct {
	pubsubTopics map[string]*pubsub.Topic
}

func (p *KafkaPublisher) Publish(ctx context.Context, topic string, idempotentId string, event []byte) error {
	return p.pubsubTopics[topic].Send(ctx, &pubsub.Message{
		Body: event,
		Metadata: map[string]string{
			"id": idempotentId,
		},
	})
}

func NewPublisher(pubsubTopics map[string]*pubsub.Topic) publisher.Publisher {
	return &KafkaPublisher{
		pubsubTopics: pubsubTopics,
	}
}
