package main

import (
	"fmt"

	"github.com/budisetionugroho123/go_donation/internal/config"
	"github.com/budisetionugroho123/go_donation/internal/handlers"
	"github.com/budisetionugroho123/go_donation/internal/router"
	"github.com/budisetionugroho123/go_donation/internal/services"
	"github.com/budisetionugroho123/go_donation/migrations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("first run")
	db := config.InitDB()
	migrations.RunMigration()
	app := fiber.New()

	roleService := services.NewRoleService(db)
	roleHandler := handlers.NewRoleHandler(roleService)
	router.RoleRoute(app, roleHandler)

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)
	router.UserRoutes(app, userHandler)

	organizationService := services.NewOrganizationService(db)
	organizationHandler := handlers.NewOrganizationHandler(organizationService)
	router.OrganizationRoute(app, organizationHandler)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	fmt.Println(db)
	app.Listen(":5000")
}
