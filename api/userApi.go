package api

import (
	"authenticaiton-authorization/service"
	"github.com/gofiber/fiber/v2"
)

var (
	userService = new(service.UserService)
)

func userApi(api fiber.Router) {
	userGroup := api.Group("/user")
	userGroup.Post("/register", userService.Create)
	userGroup.Post("/login", userService.Login)
	userGroup.Post("/register", userService.Create)
	userGroup.Post("/verify", userService.Verify)
}
