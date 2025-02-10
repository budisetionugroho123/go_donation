package handlers

import (
	"strconv"

	"github.com/budisetionugroho123/go_donation/internal/dto"
	"github.com/budisetionugroho123/go_donation/internal/models"
	"github.com/budisetionugroho123/go_donation/internal/services"
	"github.com/budisetionugroho123/go_donation/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed to create user", err.Error())
	}

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}
	return utils.SendSuccess(c, "Success create user", createdUser)
}

func (h *UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Query("email")

	// Validasi email kosong
	if email == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "Email is required", "")
	}
	user, err := h.service.GetUserByEmail(email)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to get user", err.Error())
	}
	return utils.SendSuccess(c, "Success get user by email", user)

}
func (h *UserHandler) GetAllUser(c *fiber.Ctx) error {
	users, err := h.service.GetAllUser()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to get user", err.Error())
	}
	return utils.SendSuccess(c, "Success get users", users)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid role ID", err.Error())
	}
	uintID := uint(id)

	if err := h.service.DeleteUser(uintID); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to delete user", err.Error())
	}
	return utils.SendSuccess(c, "Success delete user", nil)

}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed update user", err.Error())
	}
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}
	uintID := uint(id)
	updateUser, err := h.service.UpdateUser(uintID, user)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())

	}
	return utils.SendSuccess(c, "Success update user", updateUser)

}

func (h *UserHandler) GetUserByRole(c *fiber.Ctx) error {
	paramId := c.Params("roleId")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid role ID", err.Error())
	}

	users, err := h.service.GetUserByRole(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed get data", err.Error())
	}
	if users == nil {
		users = []dto.UserResponse{}
	}

	return utils.SendSuccess(c, "Success get user", users)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var request dto.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed request login", err.Error())

	}

	user, err := h.service.GetUserByEmail(request.Email)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed request login", err.Error())
	}
	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return utils.SendError(c, fiber.StatusUnauthorized, "Invalid credentials", "Wrong password")
	}
	token, err := h.service.GenerateToken(user)
	if err != nil {
		return utils.SendError(c, fiber.StatusUnauthorized, "Invalid credentials", err.Error())

	}
	return utils.SendSuccess(c, "Login successful", fiber.Map{"token": token})

}
