package admin_handlers

import (
	"SparksSport/cmd/app/core/providers"
	"SparksSport/pkg/admin/http/requests"
	"SparksSport/pkg/admin/services"
	"net/http"
)

type AdminHandler struct {
	adminService services.AdminService
	response     providers.Response
}

func NewAdminHandler(adminService services.AdminService, response providers.Response) *AdminHandler {
	return &AdminHandler{adminService: adminService, response: response}
}

func (h *AdminHandler) GetAdmins(w http.ResponseWriter, r *http.Request) {
	admins, err := h.adminService.GetAllAdmins()
	if err != nil {
		h.response.Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	h.response.Success(w, r, admins, "", http.StatusOK)
}

func (h *AdminHandler) GetAdmin(w http.ResponseWriter, r *http.Request) {
	id, err := requests.ValidateAdminId(r)
	if err != nil {
		h.response.Error(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	admin, err := h.adminService.GetAdminByID(id)
	if err != nil {
		h.response.Error(w, r, err.Error(), http.StatusNotFound)
		return
	}

	h.response.Success(w, r, admin, "", http.StatusOK)
}
