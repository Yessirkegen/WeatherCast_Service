package configs

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	APIkey            string
	RedisHost         string
	RedisPort         string
	CacheTTl          int
	WeatherApiBaseUrl string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
		instance = &Config{
			APIkey:            os.Getenv("API_KEY"),
			RedisHost:         os.Getenv("REDIS_HOST"),
			RedisPort:         os.Getenv("REDIS_PORT"),
			CacheTTl:          600,
			WeatherApiBaseUrl: os.Getenv("WEATHER_API_BASE_URL"),
		}
	})
	return instance
}
