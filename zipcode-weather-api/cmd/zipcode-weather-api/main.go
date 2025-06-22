package main

import (
	"log"
	"os"

	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/infra/web"
	webserver "github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/infra/web/webserver"
)

func main() {
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
	server := webserver.NewWebServer(":8081")
	handler := web.NewWebWeatherHandler(apiKey)
	server.AddHandler("/zipcode/{zipcode}/weather", handler.Fetch)

	server.Start()
}
