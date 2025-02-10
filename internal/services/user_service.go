package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/budisetionugroho123/go_donation/internal/config"
	"github.com/budisetionugroho123/go_donation/internal/dto"
	"github.com/budisetionugroho123/go_donation/internal/models"
	"github.com/budisetionugroho123/go_donation/internal/repositories"
	"github.com/budisetionugroho123/go_donation/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUser() ([]dto.UserResponse, error)
	CreateUser(user models.User) (models.User, error)
	GenerateToken(user models.User) (string, error)
	GetUserByEmail(email string) (models.User, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, user models.User) (dto.UserResponse, error)
	GetUserByRole(roleId int) ([]dto.UserResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

// GenerateToken implements UserService.
func (*userService) GenerateToken(user models.User) (string, error) {
	secret := config.GetJwtScret()
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role_id": user.RoleID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GetUserByRole implements UserService.
func (s *userService) GetUserByRole(roleId int) ([]dto.UserResponse, error) {
	users, err := s.userRepo.GetUserByRole(roleId)

	if err != nil {
		return nil, err
	}
	var userResponses []dto.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Address:   user.Address,
			Role:      user.Role.Name,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return userResponses, nil
}

// UpdateUser implements UserService.

// DeleteUser implements UserService.
func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}

// GetUserByEmail implements UserService.
func (s *userService) GetUserByEmail(email string) (models.User, error) {

	return s.userRepo.GetUserByEmail(email)
}

// CreateUser implements UserService.
func (s *userService) CreateUser(user models.User) (models.User, error) {
	existingUser, _ := s.userRepo.GetUserByEmail(user.Email)
	fmt.Println(user.Email)
	if existingUser.ID != 0 {
		return models.User{}, errors.New("email already registered")
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.CreateUser(user)
}

// GetAllUser implements UserService.
func (s *userService) GetAllUser() ([]dto.UserResponse, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Address:   user.Address,
			Role:      user.Role.Name,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return userResponses, nil
}
func (s *userService) UpdateUser(id uint, user models.User) (dto.UserResponse, error) {
	_, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	s.userRepo.UpdateUser(id, user)
	updatedUser, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	userResponse := dto.UserResponse{
		ID:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		Phone:     updatedUser.Phone,
		Role:      updatedUser.Role.Name,
		Address:   updatedUser.Address,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return userResponse, nil
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{userRepo: repositories.NewUserRepository(db)}
}
