package main

import (
	"fmt"
	"log"

	_ "account/docs"
	"account/internal/config"
	"account/internal/logger"
	"account/internal/repository"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Account Service
// @version 1.0
// @description Account Service
// @host localhost:9000
// @BasePath /

func main() {
	// Загружаем конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализируем общий логгер с названием сервиса
	logger := logger.New(cfg)

	// Подключаемся к базе данных
	db, err := gorm.Open(postgres.Open(cfg.DbDsn), &gorm.Config{})
	if err != nil {
		logger.Error().Msgf("Failed to connect to database: %v", err)
		return
	}
	logger.Info().Msg("database connected")

	// Инициализируем репозиторий
	repo := repository.NewRepository(db, &logger)
	_ = repo

	router := gin.Default()
	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		logger.Error().Msgf("Failed to set trusted proxies: %v", err)
		return
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", PingExample)
	// Пример использования общего логгера
	logger.Info().Msg("service starting up")
	router.Run(fmt.Sprintf(":%d", cfg.Port))

}

// PingExample godoc
// @Summary		Проверка доступности сервиса
// @Description Возвращает pong
// @Tags		health
// @Success		200 {string}	string "pong"
// @Router		/ping [get]
func PingExample(c *gin.Context) {
	c.String(200, "pong")
}
