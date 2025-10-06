package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WeatherResponse struct {
	Count int           `json:"count"`
	Data  []WeatherData `json:"data"`
}

type WeatherData struct {
	AppTemp      float64        `json:"app_temp"`
	Aqi          int            `json:"aqi"`
	CityName     string         `json:"city_name"`
	Clouds       int            `json:"clouds"`
	CountryCode  string         `json:"country_code"`
	Datetime     string         `json:"datetime"`
	Dewpt        float64        `json:"dewpt"`
	Dhi          float64        `json:"dhi"`
	Dni          float64        `json:"dni"`
	ElevAngle    float64        `json:"elev_angle"`
	Ghi          float64        `json:"ghi"`
	Gust         float64        `json:"gust"`
	HAngle       float64        `json:"h_angle"`
	Lat          float64        `json:"lat"`
	Lon          float64        `json:"lon"`
	ObTime       string         `json:"ob_time"`
	Pod          string         `json:"pod"`
	Precip       float64        `json:"precip"`
	Pres         float64        `json:"pres"`
	Rh           int            `json:"rh"`
	Slp          float64        `json:"slp"`
	Snow         float64        `json:"snow"`
	SolarRad     float64        `json:"solar_rad"`
	Sources      []string       `json:"sources"`
	StateCode    string         `json:"state_code"`
	Station      string         `json:"station"`
	Sunrise      string         `json:"sunrise"`
	Sunset       string         `json:"sunset"`
	Temp         float64        `json:"temp"`
	Timezone     string         `json:"timezone"`
	Ts           int64          `json:"ts"`
	Uv           float64        `json:"uv"`
	Vis          float64        `json:"vis"`
	Weather      WeatherDetails `json:"weather"`
	WindCdir     string         `json:"wind_cdir"`
	WindCdirFull string         `json:"wind_cdir_full"`
	WindDir      int            `json:"wind_dir"`
	WindSpd      float64        `json:"wind_spd"`
}

type WeatherDetails struct {
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Code        int    `json:"code"`
}

func GetWeatherByCity(city string) (map[string]string, error) {
	var apiKey string = "fd398472bcdd48508f4790d178dd82eb"
	var baseUrl string = "https://api.weatherbit.io"

	url := fmt.Sprintf("%s/v2.0/current?city=%s&key=%s", baseUrl, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching weather data: %v", err)
		return nil, err
	}
	var weather *WeatherResponse
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

func parseWeatherResponse(resp *http.Response) (*WeatherResponse, error) {
	defer resp.Body.Close()
	var result WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode weather response: %w", err)
	}
	return &result, nil
}

func transformWeatherResponse(weather *WeatherResponse) (map[string]string, error) {
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
