//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/entity"
	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/infra/repo"
	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/infra/web"
)

var setZipcodeRepository = wire.NewSet(
	repo.NewZipcodeRepository,
	wire.Bind(new(entity.ZipcodeRepositoryInterface), new(repo.ZipcodeRepository)),
)

func NewWeatherHandler(apiKey string) *web.WebWeatherHandler {
	wire.Build(web.NewWebWeatherHandler)
	return &web.WebWeatherHandler{}
}
