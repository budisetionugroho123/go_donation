package router

import (
	"github.com/budisetionugroho123/go_donation/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, roleHandler *handlers.RoleHandler) {
	api := app.Group("/api")
	api.Post("/role", roleHandler.CreateRole)
	api.Get("/role", roleHandler.GetAllRole)
	api.Get("/role/:id", roleHandler.GetRoleById)
}
