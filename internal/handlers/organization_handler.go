package handlers

import (
	"github.com/budisetionugroho123/go_donation/internal/dto"
	"github.com/budisetionugroho123/go_donation/internal/models"
	"github.com/budisetionugroho123/go_donation/internal/services"
	"github.com/budisetionugroho123/go_donation/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type OrganizationHandler struct {
	service services.OrganizationService
}

func NewOrganizationHandler(service services.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{service: service}
}

func (h *OrganizationHandler) CreateOrganization(c *fiber.Ctx) error {
	var request dto.OrganizationRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed to create user", err.Error())
	}
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		RoleID:   uint(request.RoleID),
		Address:  request.Address,
		Phone:    request.Phone,
	}

	organization := models.Organization{
		Name:        request.OrganizationName,
		Description: request.Description,
		LogoURL:     request.LogoURL,
		ContactInfo: request.ContactInfo,
	}
	result, err := h.service.CreateOrganization(user, organization)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Failed to create user", err.Error())
	}
	return utils.SendSuccess(c, "Success delete user", result)

}
