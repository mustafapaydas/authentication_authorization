package service

import (
	"github.com/gofiber/fiber/v2"
)

type IService interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type AbstractService struct {
	IService
}

func (s *AbstractService) Create(c *fiber.Ctx) error {

	return nil
}
