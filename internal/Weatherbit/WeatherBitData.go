package Weatherbit

type WeatherBitData struct {
	AppTemp      float64           `json:"app_temp"`
	Aqi          int               `json:"aqi"`
	CityName     string            `json:"city_name"`
	Clouds       int               `json:"clouds"`
	CountryCode  string            `json:"country_code"`
	Datetime     string            `json:"datetime"`
	Dewpt        float64           `json:"dewpt"`
	Dhi          float64           `json:"dhi"`
	Dni          float64           `json:"dni"`
	ElevAngle    float64           `json:"elev_angle"`
	Ghi          float64           `json:"ghi"`
	Gust         float64           `json:"gust"`
	HAngle       float64           `json:"h_angle"`
	Lat          float64           `json:"lat"`
	Lon          float64           `json:"lon"`
	ObTime       string            `json:"ob_time"`
	Pod          string            `json:"pod"`
	Precip       float64           `json:"precip"`
	Pres         float64           `json:"pres"`
	Rh           int               `json:"rh"`
	Slp          float64           `json:"slp"`
	Snow         float64           `json:"snow"`
	SolarRad     float64           `json:"solar_rad"`
	Sources      []string          `json:"sources"`
	StateCode    string            `json:"state_code"`
	Station      string            `json:"station"`
	Sunrise      string            `json:"sunrise"`
	Sunset       string            `json:"sunset"`
	Temp         float64           `json:"temp"`
	Timezone     string            `json:"timezone"`
	Ts           int64             `json:"ts"`
	Uv           float64           `json:"uv"`
	Vis          float64           `json:"vis"`
	Weather      WeatherBitDetails `json:"weather"`
	WindCdir     string            `json:"wind_cdir"`
	WindCdirFull string            `json:"wind_cdir_full"`
	WindDir      int               `json:"wind_dir"`
	WindSpd      float64           `json:"wind_spd"`
}
