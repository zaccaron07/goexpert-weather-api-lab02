package entity

type ZipcodeRepositoryInterface interface {
	Get(string) (Zipcode, error)
}

type WeatherRepositoryInterface interface {
	GetByCityName(string) (Weather, error)
}
