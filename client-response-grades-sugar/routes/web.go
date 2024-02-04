package routes

import (
	"github.com/gofiber/fiber/v2"
	userApps "client-response-grades-sugar/apps/http/controller"
)

func PageUser(app *fiber.App){
    app.Route("/", func(api fiber.Router){
        api.Get("/", userApps.PageUser)
    })
}
