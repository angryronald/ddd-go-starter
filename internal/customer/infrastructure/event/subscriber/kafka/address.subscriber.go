package kafka

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/pubsub"

	"github.com/angryronald/ddd-go-starter/internal/customer/constant"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/event/subscriber"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository"
)

type AddressKafkaSubscriber struct {
	subscription map[string]*pubsub.Subscription
	repository   repository.AddressRepository
}

func (s *AddressKafkaSubscriber) Run(ctx context.Context) {
	go s.AddAddress(ctx)
	go s.EditAddress(ctx)
	go s.RemoveAddress(ctx)
	go s.MultipleAddOrUpdateAddress(ctx)
}

func (s *AddressKafkaSubscriber) AddAddress(ctx context.Context) {
	for {
		msg, err := s.subscription[constant.ADDRESS_CREATED_EVENT].Receive(ctx)
		if err != nil {
			// Errors from Receive indicate that Receive will no longer succeed.
			log.Printf("Receiving message: %v", err)
			break
		}
		// Do work based on the message, for example:
		fmt.Printf("Got message: %q\n", msg.Body)
		// Messages must always be acknowledged with Ack.
		msg.Ack()
	}
}

func (s *AddressKafkaSubscriber) EditAddress(ctx context.Context) {
	for {
		msg, err := s.subscription[constant.ADDRESS_UPDATED_EVENT].Receive(ctx)
		if err != nil {
			// Errors from Receive indicate that Receive will no longer succeed.
			log.Printf("Receiving message: %v", err)
			break
		}
		// Do work based on the message, for example:
		fmt.Printf("Got message: %q\n", msg.Body)
		// Messages must always be acknowledged with Ack.
		msg.Ack()
	}
}

func (s *AddressKafkaSubscriber) RemoveAddress(ctx context.Context) {
	for {
		msg, err := s.subscription[constant.ADDRESS_DELETED_EVENT].Receive(ctx)
		if err != nil {
			// Errors from Receive indicate that Receive will no longer succeed.
			log.Printf("Receiving message: %v", err)
			break
		}
		// Do work based on the message, for example:
		fmt.Printf("Got message: %q\n", msg.Body)
		// Messages must always be acknowledged with Ack.
		msg.Ack()
	}
}

func (s *AddressKafkaSubscriber) MultipleAddOrUpdateAddress(ctx context.Context) {
	for {
		msg, err := s.subscription[constant.MULTIPLE_ADDRESSES_UPSERT_EVENT].Receive(ctx)
		if err != nil {
			// Errors from Receive indicate that Receive will no longer succeed.
			log.Printf("Receiving message: %v", err)
			break
		}
		// Do work based on the message, for example:
		fmt.Printf("Got message: %q\n", msg.Body)
		// Messages must always be acknowledged with Ack.
		msg.Ack()
	}
}

func NewAddressSubscriber(
	subscription map[string]*pubsub.Subscription,
	repository repository.AddressRepository,
) subscriber.Subscriber {
	return &AddressKafkaSubscriber{
		subscription: subscription,
		repository:   repository,
	}
}
