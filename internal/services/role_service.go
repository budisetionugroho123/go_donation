package services

import (
	"github.com/budisetionugroho123/go_donation/internal/models"
	"github.com/budisetionugroho123/go_donation/internal/repositories"
	"gorm.io/gorm"
)

type roleService struct {
	roleRepo repositories.RoleRepository
}

// CreateRole implements RoleService.
func (s *roleService) CreateRole(role models.Role) (models.Role, error) {
	return s.roleRepo.CreateRole(role)
}

// GetAllRole implements RoleService.
func (s *roleService) GetAllRole() ([]models.Role, error) {
	return s.roleRepo.GetAllRole()
}

// GetRoleById implements RoleService.
func (s *roleService) GetRoleById(id uint) (models.Role, error) {
	return s.roleRepo.GetRoleById(id)
}

// UpdateRole implements RoleService.
func (s *roleService) UpdateRole(id uint, role models.Role) (models.Role, error) {
	_, err := s.roleRepo.GetRoleById(id)
	if err != nil {
		return models.Role{}, err
	}
	return s.roleRepo.UpdateRole(id, role)
}

type RoleService interface {
	GetAllRole() ([]models.Role, error)
	CreateRole(role models.Role) (models.Role, error)
	GetRoleById(id uint) (models.Role, error)
	UpdateRole(id uint, role models.Role) (models.Role, error)
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleService{roleRepo: repositories.NewRoleRepository(db)}
}
