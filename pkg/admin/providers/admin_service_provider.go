package providers

import (
	"SparksSport/cmd/app/core/providers"
	"SparksSport/pkg/admin/http/handlers"
	"SparksSport/pkg/admin/http/routes"
	"SparksSport/pkg/admin/repositories"
	"SparksSport/pkg/admin/services"
)

type AdminServiceProvider struct{}

func (a *AdminServiceProvider) Register(container *providers.ServiceContainer) {
	// Register Repository
	adminRepo := repositories.NewAdminRepository(container.DB)
	container.RegisterService("AdminRepository", adminRepo)

	// Register Service
	adminService := services.NewAdminService(adminRepo)
	container.RegisterService("AdminService", adminService)

	// Register Handler
	adminHandler := admin_handlers.NewAdminHandler(adminService, container.Response)
	container.RegisterService("AdminHandler", adminHandler)

	// Register Routes
	routes.RegisterAdminRoutes(container.Router, adminHandler)
}
