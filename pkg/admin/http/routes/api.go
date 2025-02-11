package routes

import (
	"SparksSport/pkg/admin/http/handlers"
	"github.com/gorilla/mux"
)

func RegisterAdminRoutes(router *mux.Router, adminHandler *admin_handlers.AdminHandler) {
	router.HandleFunc("/admins", adminHandler.GetAdmins).Methods("GET")
	router.HandleFunc("/admins/{id}", adminHandler.GetAdmin).Methods("GET")
}
