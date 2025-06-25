package entity

import (
	"context"
)

type WeatherRepositoryInterface interface {
	GetByZipcode(context.Context, string) (Weather, error)
}
