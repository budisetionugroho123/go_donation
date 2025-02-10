package router

import (
	"github.com/budisetionugroho123/go_donation/internal/handlers"
	"github.com/budisetionugroho123/go_donation/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	api := app.Group("/api")
	api.Post("/user/create", userHandler.CreateUser)
	api.Post("/user/login", userHandler.Login)
	api.Get("/user/get-by-email", userHandler.GetUserByEmail)
	api.Get("/users/", middleware.AuthMiddeware, userHandler.GetAllUser)
	api.Get("/get-user-by-role/:roleId", userHandler.GetUserByRole)
	api.Delete("/users/:id", userHandler.DeleteUser)
	api.Put("/users/:id", userHandler.UpdateUser)

}
