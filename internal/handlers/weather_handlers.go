package handlers

import (
	"net/http"
	"weather-service/internal/services"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	service *services.WeatherService
}

func NewWeatherHandler(service *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) GetGeo(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city parameter is required"})
		return
	}

	// Получаем данные о погоде
	geo, err := h.service.GetGeo(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"geo": geo})
}

func (h *WeatherHandler) GetWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city parameter is required"})
		return
	}

	// Получаем данные о погоде
	weather, err := h.service.GetWeather(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"weather": weather})
}

func (h *WeatherHandler) GetForecast(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City is required"})
		return
	}
	forecast, err := h.service.GetForecast(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, forecast)
}

func (h *WeatherHandler) GetAQI(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City is required"})
		return
	}
	aqi, err := h.service.GetAQI(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aqi)
}
