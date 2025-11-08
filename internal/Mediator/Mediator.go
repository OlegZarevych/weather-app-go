package mediator

type Mediator interface {
	DoWeatherMagic(city string) (map[string]string, error)
}
