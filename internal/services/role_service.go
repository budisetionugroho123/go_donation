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
func (*roleService) CreateRole(role models.Role) (models.Role, error) {
	panic("unimplemented")
}

// GetAllRole implements RoleService.
func (*roleService) GetAllRole() (models.Role, error) {
	panic("unimplemented")
}

// GetRoleById implements RoleService.
func (s *roleService) GetRoleById(id uint) (models.Role, error) {
	// if id == 1 {
	// 	return models.Role{}, errors.New("admin cannot be searched")
	// }
	return s.roleRepo.GetRoleById(id)
}

// UpdateRole implements RoleService.
func (*roleService) UpdateRole(id uint, role models.Role) (models.Role, error) {
	panic("unimplemented")
}

type RoleService interface {
	GetAllRole() (models.Role, error)
	CreateRole(role models.Role) (models.Role, error)
	GetRoleById(id uint) (models.Role, error)
	UpdateRole(id uint, role models.Role) (models.Role, error)
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleService{roleRepo: repositories.NewRoleRepository(db)}
}
