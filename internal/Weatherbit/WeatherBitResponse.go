package Weatherbit

type WeatherBitResponse struct {
	Count int              `json:"count"`
	Data  []WeatherBitData `json:"data"`
}
