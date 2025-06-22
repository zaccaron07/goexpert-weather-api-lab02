package entity

import (
	"fmt"
	"strconv"
)

type Weather struct {
	TemperatureCelsius    float64
	TemperatureFahrenheit float64
	TemperatureKelvin     float64
}

func NewWeather(temperatureCelcius float64) *Weather {
	temperatureFahrenheit, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", (temperatureCelcius*1.8)+32), 64)
	temperatureKelvin, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", temperatureCelcius+273), 64)

	return &Weather{
		TemperatureCelsius:    temperatureCelcius,
		TemperatureFahrenheit: temperatureFahrenheit,
		TemperatureKelvin:     temperatureKelvin,
	}
}
