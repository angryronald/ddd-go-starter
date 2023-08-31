package main

import (
	"fmt"
	"log"
	netHttp "net/http"

	"github.com/go-chi/chi"

	"github.com/angryronald/ddd-go-starter/cmd/ecommerce/di"
	"github.com/angryronald/ddd-go-starter/cmd/ecommerce/http"
	"github.com/angryronald/ddd-go-starter/config"
)

func main() {
	di.CollectDependencies()

	runHTTP()
	runHTTPProfiler()

	log.Println("shutdown ...")
}

func runHTTP() {
	port := config.GetValue(config.HTTP_PORT)

	if len(port) < 1 {
		panic(fmt.Sprintf("Environment Missing!\n*%s* is required", port))
	}

	var router *chi.Mux
	router = chi.NewRouter()

	router.Mount("/api", http.CompileRoute(router))

	server := &netHttp.Server{
		Addr:    port,
		Handler: router,
	}

	log.Println("HTTP transport run at ", port)
	server.ListenAndServe()
}

func runHTTPProfiler() {
	profilerPort := config.GetValue(config.HTTP_PROFILER_PORT)

	if len(profilerPort) < 1 {
		panic(fmt.Sprintf("Environment Missing!\n*%s* is required", profilerPort))
	}

	var router *chi.Mux
	router = chi.NewRouter()

	router.Mount("/profiler", http.CompileProfilingRoute(router))

	server := &netHttp.Server{
		Addr:    profilerPort,
		Handler: router,
	}

	log.Println("HTTP Profiler transport run at ", profilerPort)
	server.ListenAndServe()
}
