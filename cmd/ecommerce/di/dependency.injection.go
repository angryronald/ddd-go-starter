package di

import (
	"sync"

	"gocloud.dev/pubsub"

	"github.com/redis/go-redis/v9"

	"github.com/angryronald/ddd-go-starter/internal/customer/application/command"
	"github.com/angryronald/ddd-go-starter/internal/customer/application/query"
	"github.com/angryronald/ddd-go-starter/internal/customer/domain/customer"
	"github.com/angryronald/ddd-go-starter/internal/customer/endpoint"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/event/publisher/kafka"
	kafkaPublisher "github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/event/publisher/kafka"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/event/subscriber"
	kafkaSubscriber "github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/event/subscriber/kafka"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/external/geolocation/http"
	redisInternal "github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository/memcached/redis"
	"github.com/angryronald/ddd-go-starter/internal/customer/infrastructure/repository/sql/postgre"
	genericPostgre "github.com/angryronald/ddd-go-starter/lib/helper/generic/postgre"
)

type Dependencies struct {
	CustomerEndpoint   endpoint.CustomerEndpointInterface
	CustomerSubscriber subscriber.Subscriber
	AddressSubscriber  subscriber.Subscriber
}

var syncOnce sync.Once
var AllDependencies Dependencies

func CollectDependencies() {
	syncOnce.Do(func() {
		AllDependencies = Dependencies{
			CustomerEndpoint: endpoint.NewCustomerEndpoint(
				command.NewCustomerCommand(
					customer.NewCustomerService(
						kafkaPublisher.NewPublisher(
							map[string]*pubsub.Topic{},
						),
						postgre.NewCustomerRepository(
							genericPostgre.NewGenericRepository(
								nil,
							),
						),
					),
					http.NewGeolocationService(),
				),
				query.NewCustomerQuery(
					customer.NewCustomerService(
						kafka.NewPublisher(
							map[string]*pubsub.Topic{},
						),
						redisInternal.NewCustomerRepository(
							redis.NewClient(nil),
						),
					),
					http.NewGeolocationService(),
				),
			),
			CustomerSubscriber: kafkaSubscriber.NewCustomerSubscriber(),
			AddressSubscriber: kafkaSubscriber.NewAddressSubscriber(
				nil,
				postgre.NewAddressRepository(
					genericPostgre.NewGenericRepository(
						nil,
					),
				),
			),
		}
	})
}
