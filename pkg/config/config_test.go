package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	os.Setenv("API_KEY", "test_api_key")
	os.Setenv("REDIS_HOST", "localhost")
	os.Setenv("REDIS_PORT", "6379")
	os.Setenv("WEATHER_API_BASE_URL", "https://api.openweathermap.org/data/2.5")

	config := GetConfig()
	if config.APIKey != "test_api_key" {
		t.Errorf("Expected 'test_api_key', got %s", config.APIKey)
	}
	if config.RedisHost != "localhost" {
		t.Errorf("Expected 'localhost', got %s", config.RedisHost)
	}
	if config.RedisPort != "6379" {
		t.Errorf("Expected '6379', got %s", config.RedisPort)
	}
	if config.WeatherAPIBaseURL != "https://api.openweathermap.org/data/2.5" {
		t.Errorf("Expected 'https://api.openweathermap.org/data/2.5', got %s", config.WeatherAPIBaseURL)
	}
}
