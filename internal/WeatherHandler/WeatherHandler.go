package internal

type WeatherHandler interface {
	GetWeatherByCity(city string) (map[string]string, error)
}
