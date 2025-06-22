package entity

import (
	"context"

	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/infra/repo"
)

type WeatherRepositoryInterface interface {
	GetByZipcode(context.Context, string) (repo.WeatherResponse, error)
}
