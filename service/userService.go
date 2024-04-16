package service

import (
	"authenticaiton-authorization/entity"
	"authenticaiton-authorization/helper"
	logic2 "authenticaiton-authorization/logic"
	"authenticaiton-authorization/utils"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type UserService struct {
	*AbstractService
}

var (
	logic = new(logic2.UserLogic)
)

func (s *UserService) Create(c *fiber.Ctx) error {
	var user entity.User
	body := c.Body()
	if err := sonic.Unmarshal(body, &user); err != nil {
		return utils.UnprocessableEntity(c, err.Error())
	}

	token, err := logic.Create(&user)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(token)
}
func (s *UserService) Verify(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	jwtToken, err := helper.VerifyToken(strings.ReplaceAll(authHeader, "Bearer ", ""))
	if err != nil {
	}
	fmt.Println(jwtToken)
	return c.Status(200).JSON(authHeader)
}

func (s *UserService) Login(c *fiber.Ctx) error {
	var user entity.User
	body := c.Body()
	if err := sonic.Unmarshal(body, &user); err != nil {
		return utils.UnprocessableEntity(c, err.Error())
	}

	token, err := logic.Login(&user)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(token)
}
