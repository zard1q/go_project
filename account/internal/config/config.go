package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// Config содержит конфигурацию приложения
type Config struct {
	ServiceName string `env:"SERVICE_NAME,required" envDefault:"account-service"`
	AppEnv      string `env:"APP_ENV,required" envDefault:"development"`
	Host        string `env:"HTTP_HOST,required" envDefault:"localhost"`
	Port        int    `env:"HTTP_PORT,required" envDefault:"9000"`
	LogLevel    string `env:"LOG_LEVEL,required" envDefault:"info"`
	DbDsn       string `env:"DB_DSN,required"`
}

// Load загружает конфигурацию из переменных окружения
func Load() (*Config, error) {
	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
