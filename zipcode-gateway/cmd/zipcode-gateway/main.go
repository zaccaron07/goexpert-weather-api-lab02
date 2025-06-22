package main

import (
	"context"
	"log"
	"os"

	tracing "github.com/zaccaron07/goexpert-weather-api-lab02/internal/tracing"
	web "github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/infra/web"
	webserver "github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/infra/web/webserver"
)

func main() {
	shutdown, err := tracing.InitTracer("zipcode-gateway")
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer shutdown(context.Background())

	weatherAPIURL := getWeatherAPIURL()
	startServer(weatherAPIURL)
}

func startServer(weatherAPIURL string) {
	server := webserver.NewWebServer(":8080")
	handler := web.NewWebWeatherHandler(weatherAPIURL)
	server.AddHandler("/zipcode", handler.ZipcodeHandler)

	server.Start()
}

func getWeatherAPIURL() string {
	weatherAPIURL := os.Getenv("WEATHER_API_URL")
	if weatherAPIURL == "" {
		weatherAPIURL = "http://localhost:8081"
	}
	return weatherAPIURL
}
