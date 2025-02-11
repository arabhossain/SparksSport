package repositories

import (
	"SparksSport/pkg/admin/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAll() ([]models.Admin, error)
	GetByID(id uint) (*models.Admin, error)
	Create(admin *models.Admin) error
	Update(admin *models.Admin) error
	Delete(id uint) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) GetByID(id uint) (*models.Admin, error) {
	var admin models.Admin
	result := r.db.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (r *adminRepository) Create(admin *models.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) Update(admin *models.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) Delete(id uint) error {
	return r.db.Delete(&models.Admin{}, id).Error
}

func (r *adminRepository) GetAll() ([]models.Admin, error) {
	var admins []models.Admin
	result := r.db.Find(&admins)
	return admins, result.Error
}
