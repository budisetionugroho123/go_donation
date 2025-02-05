package handlers

import (
	"strconv"
	"time"

	"github.com/budisetionugroho123/go_donation/internal/models"
	"github.com/budisetionugroho123/go_donation/internal/repositories"
	"github.com/budisetionugroho123/go_donation/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	repo repositories.RoleRepository
}

func NewRoleHandler(repo repositories.RoleRepository) *RoleHandler {
	return &RoleHandler{repo: repo}
}

func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed to create role", err.Error())

	}
	role.CreatedAt = time.Now()
	createdRole, err := h.repo.CreateRole(role)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create role", err.Error())

	}
	return utils.SendSuccess(c, createdRole)

}
func (h *RoleHandler) GetAllRole(c *fiber.Ctx) error {
	roles, err := h.repo.GetAllRole()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to get data role", err.Error())

	}
	return utils.SendSuccess(c, roles)
}

func (h *RoleHandler) GetRoleById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid role ID", err.Error())
	}
	uintID := uint(id)

	// Ambil role dari repository
	role, err := h.repo.GetRoleById(uintID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to get role", err.Error())
	}
	return utils.SendSuccess(c, role)

}
