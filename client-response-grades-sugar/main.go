package main

import (
	// userApps "client-response-grades-sugar/apps/user"
	routes "client-response-grades-sugar/routes"
	config "client-response-grades-sugar/config"
	middleware "client-response-grades-sugar/apps/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDatabase()

	routes.UserRoutes(app)
	routes.PageUser(app)
	middleware.JWTMiddleware(app)
	routes.DashboardRoutes(app)

	app.Listen(":3000")
}