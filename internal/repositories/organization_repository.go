package repositories

import (
	"github.com/budisetionugroho123/go_donation/internal/models"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	CreateOrganization(organization models.Organization) (models.Organization, error)
}

type organizationRepository struct {
	db *gorm.DB
}

// CreateOrganization implements OrganizationRepository.
func (r *organizationRepository) CreateOrganization(organization models.Organization) (models.Organization, error) {
	err := r.db.Create(&organization).Error
	if err != nil {
		return models.Organization{}, err
	}
	return organization, nil
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db}
}
