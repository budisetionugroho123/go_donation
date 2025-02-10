package handlers

import (
	"strconv"
	"time"

	"github.com/budisetionugroho123/go_donation/internal/models"
	"github.com/budisetionugroho123/go_donation/internal/services"
	"github.com/budisetionugroho123/go_donation/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	service services.RoleService
}

func NewRoleHandler(service services.RoleService) *RoleHandler {
	return &RoleHandler{
		service: service,
	}
}

func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed to create role", err.Error())

	}
	role.CreatedAt = time.Now()
	createdRole, err := h.service.CreateRole(role)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create role", err.Error())

	}
	return utils.SendSuccess(c, "Success create role", createdRole)

}
func (h *RoleHandler) UpdateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed update role", err.Error())
	}
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid role ID", err.Error())
	}
	uintID := uint(id)

	role.ID = uintID
	updateRole, err := h.service.UpdateRole(uintID, role)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to update role", err.Error())

	}
	return utils.SendSuccess(c, "Success update role", updateRole)

}

func (h *RoleHandler) GetAllRole(c *fiber.Ctx) error {
	roles, err := h.service.GetAllRole()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to get data role", err.Error())

	}
	return utils.SendSuccess(c, "Success get all role", roles)
}

func (h *RoleHandler) GetRoleById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid role ID", err.Error())
	}
	uintID := uint(id)

	// Ambil role dari repository
	role, err := h.service.GetRoleById(uintID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to get role", err.Error())
	}
	return utils.SendSuccess(c, "Success get role", role)

}
