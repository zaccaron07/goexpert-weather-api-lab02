package main

import (
	"context"
	"log"
	"os"

	tracing "github.com/zaccaron07/goexpert-weather-api-lab02/internal/tracing"
	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/infra/web"
	webserver "github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/infra/web/webserver"
)

func main() {
	collectorURL := os.Getenv("OTEL_COLLECTOR_URL")
	if collectorURL == "" {
		collectorURL = "localhost:4318"
	}
	shutdown, err := tracing.InitTracer("zipcode-weather-api", collectorURL)
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer shutdown(context.Background())

	apiKey := getAPIKey()
	startServer(apiKey)
}

func getAPIKey() string {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY environment variable is required")
	}
	return apiKey
}

func startServer(apiKey string) {
	server := webserver.NewWebServer("127.0.0.1:8081")
	handler := web.NewWebWeatherHandler(apiKey)
	server.AddHandler("/zipcode/{zipcode}/weather", handler.Fetch)

	server.Start()
}
