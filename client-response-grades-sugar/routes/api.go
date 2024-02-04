package routes

import (
	"github.com/gofiber/fiber/v2"
	userApps "client-response-grades-sugar/apps/http/controller"
)

func UserRoutes(app *fiber.App) {
    app.Route("/api", func(api fiber.Router) {
        v1 := api.Group("/v1")
        user := v1.Group("/user")
        user.Get("/", userApps.GetAllUsers)
        user.Get("/role/:role", userApps.GetUserByRole)
        user.Get("/name/:name", userApps.GetUserByName)
        user.Get("/id/:id", userApps.GetUserByID)

        auth := v1.Group("/auth")
        auth.Post("/register", userApps.Register)
        auth.Post("/login", userApps.Login)
        // api.Put("/:id", UpdateUser)
        // api.Delete("/:id", DeleteUser)
    })
}

func DashboardRoutes(app *fiber.App){
    app.Route("/dashboard", func(api fiber.Router){
        api.Get("/superuser", userApps.SuperUserDashboard)
    })
}
