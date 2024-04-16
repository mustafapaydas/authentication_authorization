package api

import "github.com/gofiber/fiber/v2"

func Version1(api fiber.Router) {
	version1Group := api.Group("/v1")
	userApi(version1Group)
}
