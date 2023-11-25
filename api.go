package main

import (
	"github.com/gofiber/fiber/v2"
)

func apiAgent(c *fiber.Ctx) error {
	return c.JSON("not implemented")
}

func apiPing(c *fiber.Ctx) error {
	return c.JSON(CONFIG)
}
