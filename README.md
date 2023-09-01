# ddd-go-starter
A starter pack for Domain-Driven Design in Go

## Layering
1. **User Interface** is reflected by endpoints which responsible for handling the requests and generate responses
2. **Application Layer** is reflected by application package which contains use cases
3. **Domain Layer** is reflected by domain package which contains the business logic for specific domain
4. **Infrastructure Layer** is reflected by infrastructure package which contains **event** for domain events, **repository** for repository layer, and **external** for external layer

![Layered Architecture - Eric Evans, 2003](https://github.com/angryronald/ddd-go-starter/blob/main/DDD-Layered-Architecture.jpg)

## Tech Stack
- Redis
- Vault
- Kafka
- Postgre
- Mongo

## Library Dependencies
- go-chi
- sqlx (can be replaced with gorm)
- vault
- mongo
- gocloud.dev pubsub
- cors
- prometheus

## Notes
This starter pack is not contains the testing but having unit test, integration test, and e2e test to 70% is the minimum requirement to achieve the confidence level of 88%.
dockertest (github.com/ory/dockertest) could help for the integration test.