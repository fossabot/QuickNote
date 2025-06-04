package handler

import (
	"github.com/Sn0wo2/QuickNote/response"
	"github.com/gofiber/fiber/v2"
)

func Health(_ string) func(ctx *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(response.New(true, "success"))
	}
}
