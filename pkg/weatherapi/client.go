package weatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-service/pkg/config"
)

// WeatherResponse структура для хранения ответа от API
type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
}

// WeatherClient интерфейс для работы с погодным API
type WeatherClient interface {
	GetWeather(city string) (WeatherResponse, error)
}

// weatherClient структура для работы с погодным API
type weatherClient struct {
	apiKey  string
	baseURL string
}

// NewWeatherClient создает новый экземпляр WeatherClient
func NewWeatherClient() WeatherClient {
	cfg := config.GetConfig()
	return &weatherClient{
		apiKey:  cfg.APIKey,
		baseURL: cfg.WeatherAPIBaseURL,
	}
}

// GetWeather получает данные о погоде для указанного города
func (c *weatherClient) GetWeather(city string) (WeatherResponse, error) {
	// Формируем URL для запроса
	url := fmt.Sprintf("%s/weather?q=%s&appid=%s", c.baseURL, city, c.apiKey) // Исправлено: добавлен '&' перед 'appid'
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var weatherResp WeatherResponse
	// Исправлено: используем resp.Body с заглавной буквы
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return WeatherResponse{}, err // Исправлено: убрали лишний 'weatherResp'
	}

	return weatherResp, nil // Исправлено: убрали лишний 'weatherResp'
}
