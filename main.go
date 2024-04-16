package main

import (
	"authenticaiton-authorization/api"
	"authenticaiton-authorization/utils"
	"github.com/gofiber/fiber/v2"
	"runtime"
	"time"
)

var app *fiber.App

func middlewareForPanic(c *fiber.Ctx) error {
	defer func() {
		c.Accepts("json", "text")
		c.Accepts("application/json")
		if r := recover(); r != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": r,
			})
		}
	}()
	return c.Next()
}
func init() {
	app = fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		IdleTimeout:   120 * time.Second,
		ReadTimeout:   120 * time.Second,
		WriteTimeout:  120 * time.Second,
		ServerHeader:  "AuthenticationAndAuthorization",
		AppName:       "authentication-server",
		BodyLimit:     10 << 20,
		Concurrency:   512 * 1024,
	})
	app.Use(middlewareForPanic)
	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return utils.ErrorHandler(c, err)
		}
		return nil
	})
	authenticationApi := app.Group("/api")
	api.Version1(authenticationApi)
}

func main() {
	maxCores := runtime.NumCPU() / 4
	runtime.GOMAXPROCS(maxCores)

	for {
		err := app.Listen(":8080")
		if err != nil {

		}
	}

}
