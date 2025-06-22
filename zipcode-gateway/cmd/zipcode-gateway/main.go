package main

import (
	"os"

	web "github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/infra/web"
	webserver "github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/infra/web/webserver"
)

func main() {
	weatherAPIURL := getWeatherAPIURL()
	startServer(weatherAPIURL)
}

func startServer(weatherAPIURL string) {
	server := webserver.NewWebServer("127.0.0.1:8080")
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
