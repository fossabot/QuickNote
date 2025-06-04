package router

import (
	"github.com/Sn0wo2/QuickNote/log"
	"github.com/Sn0wo2/QuickNote/response"
	"github.com/Sn0wo2/QuickNote/router/handler"
	"github.com/Sn0wo2/QuickNote/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"go.uber.org/zap"
)

func Setup(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	api := app.Group(setup.APIVersion)

	const healthPath = "health"
	const notePath = "note"

	api.Group(healthPath, handler.Health(healthPath))
	api.Group(notePath, handler.Note(notePath))

	app.Use("*", func(ctx *fiber.Ctx) error {
		log.Instance.Warn("Route not found",
			zap.String("ctx", ctx.String()),
		)

		return ctx.Status(fiber.StatusNotFound).JSON(response.New(false, "Route not found"))
	})
}
