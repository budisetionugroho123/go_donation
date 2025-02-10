package router

import (
	"github.com/budisetionugroho123/go_donation/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func OrganizationRoute(app *fiber.App, organizationHandler *handlers.OrganizationHandler) {
	api := app.Group("/api")
	api.Post("/organization", organizationHandler.CreateOrganization)
}
