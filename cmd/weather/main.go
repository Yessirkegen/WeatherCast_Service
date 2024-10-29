package main

import (
	"weather-service/internal/handlers"
	"weather-service/internal/services"
	"weather-service/internal/weatherapi"
	"weather-service/pkg/cache"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cache := cache.NewCache()
	weatherClient := weatherapi.NewWeatherClient()
	weatherService := services.NewWeatherService(cache, weatherClient)
	weatherHandler := handlers.NewWeatherHandler(weatherService)

	/*r.GET("/weather", weatherHandler.GetWeather)
	r.GET("/forecast", weatherHandler.GetForecast) // Новый маршрут
	r.GET("/aqi", weatherHandler.GetAQI)    */            // Новый маршрут
	r.GET("/weather-data", weatherHandler.GetWeatherData) // Новый маршрут для получения данных о погоде
	r.Run(":8061")
}
