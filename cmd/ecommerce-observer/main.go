package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/angryronald/ddd-go-starter/cmd/ecommerce/di"
)

func main() {
	di.CollectDependencies()

	// register all the subscribers
	go di.AllDependencies.CustomerSubscriber.Run(context.Background())
	go di.AllDependencies.AddressSubscriber.Run(context.Background())

	log.Printf("E-Commerce Subscribers are up")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("E-Commerce Subscribers down")
}
