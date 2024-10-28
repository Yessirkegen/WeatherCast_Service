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

func (s *WeatherService) GetGeo(city string) (string, error) {
	cachedData, err := s.cache.Get(city)
	if err == nil && cachedData != "" {
		return cachedData, nil
	}
	GeoData, err := s.client.GetGeo(city)
	if err != nil {
		return "", err
	}
	GeoJSON, err := json.Marshal(GeoData)
	if err != nil {
		return "", err
	}

	s.cache.Set(city, string(GeoJSON), config.GetConfig().CacheTTL)
	return string(GeoJSON), nil
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

func (s *WeatherService) GetForecast(city string) (string, error) {
	cachedData, err := s.cache.Get(city)
	if err == nil && cachedData != "" {
		return cachedData, nil
	}
	forecastData, err := s.client.GetForecast(city)
	if err != nil {
		return "", err
	}
	forecastJSON, err := json.Marshal(forecastData)
	if err != nil {
		return "", err
	}
	s.cache.Set(city+"_forecast", string(forecastJSON), config.GetConfig().CacheTTL)
	return string(forecastJSON), nil
}

func (s *WeatherService) GetAQI(city string) (string, error) {
	cachedData, err := s.cache.Get(city)
	if err == nil && cachedData != "" {
		return cachedData, nil
	}
	aqiData, err := s.client.GetAQI(city)
	if err != nil {
		return "", err
	}
	aqiJSON, err := json.Marshal(aqiData)
	if err != nil {
		return "", err
	}
	s.cache.Set(city+"_aqi", string(aqiJSON), config.GetConfig().CacheTTL)
	return string(aqiJSON), nil
}
