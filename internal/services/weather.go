package services

import (
	"encoding/json"
	"weather-service/pkg/cache"
	"weather-service/pkg/config"
	"weather-service/pkg/weatherapi"
)

type WeatherService struct {
	cache  *cache.Cache
	client weatherapi.WeatherClient
}

func NewWeatherService(cache *cache.Cache, client weatherapi.WeatherClient) *WeatherService {
	return &WeatherService{
		cache:  cache,
		client: client,
	}
}

func (s *WeatherService) GetWeather(city string) (string, error) {
	cachedData, err := s.cache.Get(city)
	if err == nil && cachedData != "" {
		return cachedData, nil
	}
	weatherData, err := s.client.GetWeather(city)
	if err != nil {
		return "", err
	}
	weatherJSON, err := json.Marshal(weatherData)
	if err != nil {
		return "", err
	}

	s.cache.Set(city, string(weatherJSON), config.GetConfig().CacheTTL)
	return string(weatherJSON), nil
}
