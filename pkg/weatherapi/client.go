package weatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-service/pkg/config"
)

type GeoResponse struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state,omitempty"`
}

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

type ForecastResponse struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			Humidity  int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		} `json:"wind"`
		Visibility int     `json:"visibility"`
		Pop        float64 `json:"pop"`
		DtTxt      string  `json:"dt_txt"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country string `json:"country"`
	} `json:"city"`
}

type AQIResponse struct {
	Coord []float64 `json:"coord"`
	List  []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			AQI int `json:"aqi"`
		} `json:"main"`
		Components struct {
			CO    float64 `json:"co"`
			NO    float64 `json:"no"`
			NO2   float64 `json:"no2"`
			O3    float64 `json:"o3"`
			SO2   float64 `json:"so2"`
			PM2_5 float64 `json:"pm2_5"`
			PM10  float64 `json:"pm10"`
			NH3   float64 `json:"nh3"`
		} `json:"components"`
	} `json:"list"`
}

type WeatherClient interface {
	GetGeo(city string) (GeoResponse, error)
	GetWeather(city string) (WeatherResponse, error)
	GetForecast(city string) (ForecastResponse, error) // Новый метод
	GetAQI(city string) (AQIResponse, error)           // Новый метод
}

// weatherClient структура для работы с погодным API
type weatherClient struct {
	apiKey         string
	baseWeatherURL string
	baseGeoURL     string
}

// NewWeatherClient создает новый экземпляр WeatherClient
func NewWeatherClient() WeatherClient {
	cfg := config.GetConfig()
	return &weatherClient{
		apiKey:         cfg.APIKey,
		baseWeatherURL: cfg.WeatherAPIBaseURL,
		baseGeoURL:     cfg.GeoApiBaseUrl,
	}
}

func (c *weatherClient) GetGeo(city string) (GeoResponse, error) {
	// Формируем URL для запроса к геокодинг API
	url := fmt.Sprintf("%s/direct?q=%s&limit=5&appid=%s", c.baseGeoURL, city, c.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return GeoResponse{}, err
	}
	defer resp.Body.Close()

	// Проверка на успешный ответ от API
	if resp.StatusCode != http.StatusOK {
		return GeoResponse{}, fmt.Errorf("не удалось получить геоданные: %s", resp.Status)
	}

	var geoResponses []GeoResponse
	// Декодируем ответ как массив структур GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResponses); err != nil {
		return GeoResponse{}, err
	}

	// Проверяем, что мы получили хотя бы одну запись
	if len(geoResponses) == 0 {
		return GeoResponse{}, fmt.Errorf("геоданные не найдены для города: %s", city)
	}

	// Возвращаем первый элемент из массива
	return geoResponses[0], nil
}

// GetWeather получает данные о погоде для указанного города
func (c *weatherClient) GetWeather(city string) (WeatherResponse, error) {
	// Формируем URL для запроса
	url := fmt.Sprintf("%s/weather?q=%s&appid=%s&units=metric", c.baseWeatherURL, city, c.apiKey) // Исправлено: добавлен '&' перед 'appid'
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

func (c *weatherClient) GetForecast(city string) (ForecastResponse, error) {
	url := fmt.Sprintf("%s/forecast?q=%s&appid=%s", c.baseWeatherURL, city, c.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return ForecastResponse{}, err
	}
	defer resp.Body.Close()
	var forecastResp ForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&forecastResp); err != nil {
		return ForecastResponse{}, err
	}
	return forecastResp, nil
}

func (c *weatherClient) GetAQI(city string) (AQIResponse, error) {
	url := fmt.Sprintf("%s/air_pollution?q=%s&appid=%s", c.baseWeatherURL, city, c.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return AQIResponse{}, err
	}
	defer resp.Body.Close()
	var aqiResp AQIResponse
	if err := json.NewDecoder(resp.Body).Decode(&aqiResp); err != nil {
		return AQIResponse{}, err
	}
	return aqiResp, nil
}
