package services

import (
	"SparksSport/pkg/admin/models"
	"SparksSport/pkg/admin/repositories"
)

type AdminService interface {
	GetAllAdmins() ([]models.Admin, error)
	GetAdminByID(id uint) (*models.Admin, error)
	CreateAdmin(admin *models.Admin) error
	UpdateAdmin(admin *models.Admin) error
	DeleteAdmin(id uint) error
}

type adminService struct {
	repo repositories.AdminRepository
}

func NewAdminService(repo repositories.AdminRepository) AdminService {
	return &adminService{repo: repo}
}

func (a adminService) GetAllAdmins() ([]models.Admin, error) {
	return a.repo.GetAll()
}

func (a adminService) GetAdminByID(id uint) (*models.Admin, error) {
	return a.repo.GetByID(id)
}

func (a adminService) CreateAdmin(admin *models.Admin) error {
	return a.repo.Create(admin)
}

func (a adminService) UpdateAdmin(admin *models.Admin) error {
	return a.repo.Update(admin)
}

func (a adminService) DeleteAdmin(id uint) error {
	return a.repo.Delete(id)
}
