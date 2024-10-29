package services

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"weather-service/internal/weatherapi"
	"weather-service/pkg/cache"
	"weather-service/pkg/config"
)

type WeatherService struct {
	cache  *cache.Cache
	client weatherapi.WeatherClient
}

type GeoLocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func NewWeatherService(cache *cache.Cache, client weatherapi.WeatherClient) *WeatherService {
	return &WeatherService{
		cache:  cache,
		client: client,
	}
}

func (s *WeatherService) GetCompleteWeatherData(city string) (map[string]interface{}, error) {
	geoString, err := s.GetGeo(city)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения геоданных: %v", err)
	}

	var geo GeoLocation
	if err := json.Unmarshal([]byte(geoString), &geo); err != nil {
		return nil, fmt.Errorf("ошибка распарсивания геоданных: %v", err)
	}

	weatherCh := make(chan interface{}, 1)
	forecastCh := make(chan interface{}, 1)
	aqiCh := make(chan interface{}, 1)
	errCh := make(chan error, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	// Запускаем горутины
	go func() {
		defer wg.Done()
		weatherData, err := s.GetWeather(city)
		if err != nil {
			errCh <- fmt.Errorf("ошибка получения данных о погоде: %v", err)
			return
		}
		weatherCh <- weatherData
	}()

	go func() {
		defer wg.Done()
		forecastData, err := s.GetForecast(city)
		if err != nil {
			errCh <- fmt.Errorf("ошибка получения данных о прогнозе: %v", err)
			return
		}
		forecastCh <- forecastData
	}()

	go func() {
		defer wg.Done()
		aqiData, err := s.GetAQI(city, geo.Lat, geo.Lon)
		if err != nil {
			errCh <- fmt.Errorf("ошибка получения данных о качестве воздуха: %v", err)
			return
		}
		aqiCh <- aqiData
	}()

	// Закрываем каналы после завершения всех горутин
	go func() {
		wg.Wait()
		close(errCh)
		close(weatherCh)
		close(forecastCh)
		close(aqiCh)
	}()

	// Устанавливаем тайм-аут и ожидаем завершения всех горутин
	select {
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("превышено время ожидания для получения данных от API")
	case err := <-errCh:
		return nil, err
	default:
		// Получаем данные из каналов после завершения всех горутин
		responseData := map[string]interface{}{
			"geo":      geo,
			"weather":  <-weatherCh,
			"forecast": <-forecastCh,
			"aqi":      <-aqiCh,
		}
		return responseData, nil
	}
}

func (s *WeatherService) GetGeo(city string) (string, error) {
	cachedData, err := s.cache.Get(city + "_geo")
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

	s.cache.Set(city+"_geo", string(GeoJSON), config.GetConfig().CacheTTL)
	return string(GeoJSON), nil
}

func (s *WeatherService) GetWeather(city string) (string, error) {
	cachedData, err := s.cache.Get(city + "_weather")
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

	s.cache.Set(city+"_weather", string(weatherJSON), config.GetConfig().CacheTTL)
	return string(weatherJSON), nil
}

func (s *WeatherService) GetForecast(city string) (string, error) {
	cachedData, err := s.cache.Get(city + "_forecast")

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

func (s *WeatherService) GetAQI(city string, geolat float64, geolon float64) (string, error) {
	cachedData, err := s.cache.Get(city + "_aqi")
	if err == nil && cachedData != "" {
		return cachedData, nil
	}
	aqiData, err := s.client.GetAQI(geolat, geolon)
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
