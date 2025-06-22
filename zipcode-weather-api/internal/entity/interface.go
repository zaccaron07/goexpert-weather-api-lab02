package entity

import "context"

type ZipcodeRepositoryInterface interface {
	Get(context.Context, string) (Zipcode, error)
}

type WeatherRepositoryInterface interface {
	GetByCityName(context.Context, string) (Weather, error)
}
