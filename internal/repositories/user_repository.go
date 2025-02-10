package repositories

import (
	"fmt"

	"github.com/budisetionugroho123/go_donation/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, user models.User) (models.User, error)
	GetUserById(id uint) (models.User, error)
	GetUserByRole(roleId int) ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// GetUserByRole implements UserRepository.
func (r *userRepository) GetUserByRole(roleId int) ([]models.User, error) {
	var users []models.User
	result := r.db.Preload("Role").Where("role_id = ?", roleId).Find(&users)
	return users, result.Error
}

// GetUserById implements UserRepository.
func (r *userRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	result := r.db.Preload("Role").Where("id", id).First(&user)
	return user, result.Error
}

// UpdateUser implements UserRepository.
func (r *userRepository) UpdateUser(id uint, user models.User) (models.User, error) {
	err := r.db.Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// DeleteUser implements UserRepository.
func (r *userRepository) DeleteUser(id uint) error {
	var user models.User
	return r.db.Where("id = ?", id).Delete(&user).Error
}

// GetUserByEmail implements UserRepository.
func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := r.db.Where("email", email).First(&user)
	fmt.Println(result)
	return user, result.Error
}

// CreateUser implements UserRepository.
func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	err = r.db.Preload("Role").First(&user, user.ID).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// GetAllUser implements UserRepository.
func (r *userRepository) GetAllUser() ([]models.User, error) {
	var users []models.User
	result := r.db.Preload("Role").Find(&users)
	return users, result.Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
