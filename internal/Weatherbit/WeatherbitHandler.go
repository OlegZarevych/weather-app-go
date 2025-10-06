package Weatherbit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WeatherbitHandler struct{}

func (h WeatherbitHandler) GetWeatherByCity(city string) (map[string]string, error) {
	var apiKey string = ""
	var baseUrl string = "https://api.weatherbit.io"

	url := fmt.Sprintf("%s/v2.0/current?city=%s&key=%s", baseUrl, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching weather data: %v", err)
		return nil, err
	}
	var weather *WeatherBitResponse
	weather, err = parseWeatherResponse(resp)
	if err != nil {
		log.Printf("Error parsing weather data: %v", err)
		return nil, err
	}
	weatherMap, err := transformWeatherResponse(weather)
	if err != nil {
		log.Printf("Error transforming weather data: %v", err)
		return nil, err
	}
	return weatherMap, nil
}

func parseWeatherResponse(resp *http.Response) (*WeatherBitResponse, error) {
	defer resp.Body.Close()
	var result WeatherBitResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode weather response: %w", err)
	}
	return &result, nil
}

func transformWeatherResponse(weather *WeatherBitResponse) (map[string]string, error) {
	if weather.Count == 1 {
		log.Printf("Transform weather response to map")
		weatherMap := make(map[string]string, 2)
		weatherMap["description"] = weather.Data[0].Weather.Description
		weatherMap["temp"] = fmt.Sprintf("%v", weather.Data[0].Temp)
		return weatherMap, nil
	}
	if weather.Count == 0 {
		log.Printf("No weather data found.")
		return nil, fmt.Errorf("no weather data found")
	} else {
		log.Fatal("There is more than one weather data entry. I don't know what to do.")
		return nil, fmt.Errorf("unexpected number of weather data entries: %d", weather.Count)
	}
}
