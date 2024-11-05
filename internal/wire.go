//go:build wireinject

package internal

import (
	"WareFlow/internal/config"
	"WareFlow/internal/repository"
	"WareFlow/internal/repository_mongo"
	"WareFlow/internal/usecase"
	"WareFlow/pkg/mongo"
	"github.com/google/wire"
)

// Провайдеры конфигурации и MongoDB клиента
func provideConfig() *config.Config {
	return config.NewConfig()
}

func provideMongoClient(cfg *config.Config) (*mongo.Client, error) {
	return config.NewMongoClient(cfg)
}

// Общая функция инициализации приложения
func InitializeApp() (*usecase.OptimizationService, error) {
	wire.Build(
		// Конфигурация
		provideConfig,
		provideMongoClient,

		// Репозитории
		repository.NewRepository,        // основной репозиторий интерфейсов
		repository_mongo.NewCargoDB,     // провайдер MongoDB для Cargo
		repository_mongo.NewTransportDB, // провайдер MongoDB для Transport

		// Сервисы
		usecase.NewOptimizationService, // сервис, использующий репозитории

		// Добавьте любые другие зависимости, которые вам нужны
	)

	// возвращаем основной сервис или структуру, которая представляет приложение
	return &usecase.OptimizationService{}, nil
}
