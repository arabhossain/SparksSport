package bootstrap

import (
	"SparksSport/cmd/app/core/config"
	"SparksSport/cmd/app/core/database"
	"SparksSport/cmd/app/core/middlewares"
	"SparksSport/cmd/app/core/migrations"
	"SparksSport/cmd/app/core/providers"
	"SparksSport/cmd/app/core/utils"
	"fmt"
	"github.com/gorilla/mux"

	adminServiceProviders "SparksSport/pkg/admin/providers"
)

func InitializeApp() (*providers.ServiceContainer, error) {

	utils.InitLogger()

	cfg, err := core_config.LoadConfig()
	if err != nil {
		return nil, err
	}

	err = migrations.Migrate(cfg)
	if err != nil {
		return nil, err
	}

	db, err := database.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	container := &providers.ServiceContainer{
		Config: cfg,
		Router: mux.NewRouter(),
		DB:     db,
		Response: providers.Response{
			Success: utils.SendResponse,
			Error:   utils.SendError,
		},
		Services: make(map[string]interface{}),
	}

	// Add middlewares
	container.Router.Use(middlewares.RequestIDMiddleware)

	container.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("Registered Route:", path)
		}
		return nil
	})

	RegisterProviders(container)

	return container, nil
}

func RegisterProviders(container *providers.ServiceContainer) {
	providers := []providers.ServiceProvider{
		&adminServiceProviders.AdminServiceProvider{},
		// Future providers can be added here (e.g., &UserServiceProvider{})

	}

	for _, provider := range providers {
		provider.Register(container)
	}
}
