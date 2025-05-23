package core

import (
	"context"
	"fmt"
	"sync"

	"service/internal/domain"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var (
	appConfigs *domain.AppConfig
	onceEnvs   sync.Once
)

func GetEnvironments() *domain.AppConfig {
	onceEnvs.Do(func() {
		appConfigs = initConfigs()
	})

	return appConfigs
}

func initConfigs() *domain.AppConfig {
	_ = godotenv.Load()
	var appConfig domain.AppConfig
	err := envconfig.Process(context.Background(), &appConfig)
	if err != nil {
		panic(fmt.Sprintf("error on load configs: %s", err))
	}
	return &appConfig
}
