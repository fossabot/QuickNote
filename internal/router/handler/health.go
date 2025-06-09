package handler

import (
	"github.com/Sn0wo2/QuickNote/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func Health() func(ctx *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(response.New(true, "success"))
	}
}
