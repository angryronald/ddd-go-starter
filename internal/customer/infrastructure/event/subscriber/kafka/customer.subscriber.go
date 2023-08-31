package kafka

import (
	"context"

	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/event/subscriber"
)

type CustomerKafkaSubscriber struct {
	// check address subscriber
}

func (s *CustomerKafkaSubscriber) Run(ctx context.Context) {
	// check address subscriber
}

func NewCustomerSubscriber() subscriber.Subscriber {
	return &CustomerKafkaSubscriber{}
}
