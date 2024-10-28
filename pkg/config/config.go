package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey            string
	RedisHost         string
	RedisPort         string
	CacheTTL          int
	WeatherAPIBaseURL string
	GeoApiBaseUrl     string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Fatal("Error loading .env file")
		}
		instance = &Config{
			APIKey:            os.Getenv("API_KEY"),
			RedisHost:         os.Getenv("REDIS_HOST"),
			RedisPort:         os.Getenv("REDIS_PORT"),
			CacheTTL:          600, // значение по умолчанию
			WeatherAPIBaseURL: os.Getenv("WEATHER_API_BASE_URL"),
			GeoApiBaseUrl:     os.Getenv("GEO_API_BASE_URL"),
		}
	})
	return instance
}
