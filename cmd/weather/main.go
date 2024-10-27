package main

import (
	"weather-service/internal/handlers"
	"weather-service/internal/services"
	"weather-service/pkg/cache"
	"weather-service/pkg/weatherapi"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cache := cache.NewCache()
	weatherClient := weatherapi.NewWeatherClient()
	weatherService := services.NewWeatherService(cache, weatherClient)
	weatherHandler := handlers.NewWeatherHandler(weatherService)

	r.GET("/weather", weatherHandler.GetWeather)

	r.Run(":8080")
}
