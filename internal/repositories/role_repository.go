package repositories

import (
	"github.com/budisetionugroho123/go_donation/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAllRole() ([]models.Role, error)
	CreateRole(role models.Role) (models.Role, error)
	GetRoleById(id uint) (models.Role, error)
	UpdateRole(id uint, role models.Role) (models.Role, error)
}
type roleRepository struct {
	db *gorm.DB
}

// Update implements RoleRepository.
func (r *roleRepository) UpdateRole(id uint, role models.Role) (models.Role, error) {

	// var existingRole models.Role

	// if err := r.db.First(&existingRole, id).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return models.Role{}, errors.New("role not found")
	// 	}
	// 	return models.Role{}, err
	// }

	err := r.db.Where("id", id).Updates(&role).Error
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

// CreateRole implements RoleRepository.
func (r *roleRepository) CreateRole(role models.Role) (models.Role, error) {
	err := r.db.Create(&role).Error
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

// GetAllRole implements RoleRepository.
func (r *roleRepository) GetAllRole() ([]models.Role, error) {
	var roles []models.Role
	result := r.db.Order("id ASC").Find(&roles)
	return roles, result.Error
}

// GetRoleById implements RoleRepository.
func (r *roleRepository) GetRoleById(id uint) (models.Role, error) {
	var role models.Role
	result := r.db.Where("id", id).First(&role)
	return role, result.Error
}
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}
