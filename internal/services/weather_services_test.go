package services

import (
	"testing"
	"weather-service/pkg/cache"
	"weather-service/pkg/weatherapi"
)

type MockWeatherClient struct{}

func (m *MockWeatherClient) GetWeather(city string) (weatherapi.WeatherResponse, error) {
	return weatherapi.WeatherResponse{Name: "London"}, nil
}

func TestGetWeather(t *testing.T) {
	cache := cache.NewCache()
	client := &MockWeatherClient{}
	service := NewWeatherService(cache, client)

	weather, err := service.GetWeather("London")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if weather != "London" {
		t.Errorf("Expected 'London', got %s", weather)
	}

	// Проверка кеша
	cachedData, _ := cache.Get("London")
	if cachedData != "London" {
		t.Errorf("Expected cached data 'London', got %s", cachedData)
	}
}
