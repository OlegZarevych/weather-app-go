package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	az "weather-app-go/internal/Azure"
	mediator "weather-app-go/internal/Mediator"
	genericWeatherHandler "weather-app-go/internal/WeatherHandler"
	weatherBitHandler "weather-app-go/internal/Weatherbit"
)

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	message := "This HTTP triggered function executed successfully. Pass a city in the query string for a personalized response.\n"

	city := r.URL.Query().Get("city")
	if city != "" {
		message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", city)
	}
	fmt.Fprint(w, message)
	log.Printf("Received request for city: %s", city)

	var weatherService genericWeatherHandler.WeatherHandler = weatherBitHandler.WeatherbitHandler{}
	var messageService az.MessageService = az.NewMessageService()
	var mediator mediator.Mediator = mediator.NewWeatherMediator(weatherService, messageService)
	weatherMap, err := mediator.DoWeatherMagic(city)

	if err != nil {
		log.Printf("Error fetching weather data: %v", err)
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Weather data for %s:\n", city)
	for k, v := range weatherMap {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/Weather", weatherHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
