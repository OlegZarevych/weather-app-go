package mediator

import (
	az "weather-app-go/internal/Azure"
	wh "weather-app-go/internal/WeatherHandler"
)

type WeatherMediator struct {
	weatherHandler wh.WeatherHandler
	messageService az.MessageService
}

func NewWeatherMediator(handler wh.WeatherHandler) *WeatherMediator {
	return &WeatherMediator{
		weatherHandler: handler,
		messageService: az.NewMessageService(),
	}
}

func (m *WeatherMediator) DoWeatherMagic(city string) (map[string]string, error) {

	// TODO: add logic for storing data on storage account
	// TODO: add logic to send message to Service Bus queue
	// TODO: add logic to send event to Event Grid topic
	var weatherData, err = m.weatherHandler.GetWeatherByCity(city)
	if err != nil {
		return nil, err
	}
	m.messageService.SendWeatherMessage(weatherData)
	return weatherData, nil
}
