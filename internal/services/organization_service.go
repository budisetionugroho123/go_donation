package services

import (
	"fmt"

	"github.com/budisetionugroho123/go_donation/internal/dto"
	"github.com/budisetionugroho123/go_donation/internal/models"
	"github.com/budisetionugroho123/go_donation/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type OrganizationService interface {
	CreateOrganization(user models.User, organization models.Organization) (dto.OrganizationResponse, error)
}

type organizationService struct {
	db               *gorm.DB
	organizationRepo repositories.OrganizationRepository
	userRepo         repositories.UserRepository
}

// CreateOrganization implements OrganizationService.
func (s *organizationService) CreateOrganization(user models.User, organization models.Organization) (dto.OrganizationResponse, error) {
	tx := s.db.Begin()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.OrganizationResponse{}, err
	}
	user.Password = string(hashedPassword)
	result, err := s.userRepo.CreateUser(user)
	if err != nil {
		tx.Rollback()
		return dto.OrganizationResponse{}, err
	}
	fmt.Println(result)
	organization.UserID = result.ID
	resultOrganization, err := s.organizationRepo.CreateOrganization(organization)
	if err != nil {
		tx.Rollback()
		return dto.OrganizationResponse{}, err
	}
	organizationResponse := dto.OrganizationResponse{
		ID:               result.ID,
		Name:             result.Name,
		Email:            result.Email,
		Password:         result.Password,
		RoleID:           int(result.RoleID),
		Phone:            result.Phone,
		Address:          result.Address,
		OrganizationName: resultOrganization.Name,
		OrganizationID:   resultOrganization.ID,
		LogoURL:          resultOrganization.LogoURL,
		Description:      resultOrganization.Description,
		ContactInfo:      resultOrganization.ContactInfo,
	}
	return organizationResponse, nil
}

func NewOrganizationService(db *gorm.DB) OrganizationService {
	return &organizationService{db: db, organizationRepo: repositories.NewOrganizationRepository(db),
		userRepo: repositories.NewUserRepository(db)}
}
