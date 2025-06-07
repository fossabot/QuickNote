package router

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Sn0wo2/QuickNote/log"
	"github.com/Sn0wo2/QuickNote/response"
	"github.com/Sn0wo2/QuickNote/router/handler"
	"github.com/Sn0wo2/QuickNote/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func Setup(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(cors.New())

	api := app.Group(setup.APIVersion)
	const healthPath = "health"
	const notePath = "note"

	api.Group(healthPath, handler.Health(healthPath))
	api.Group(notePath, handler.Note(notePath))

	api.Use("*", func(ctx *fiber.Ctx) error {
		log.Instance.Warn("API route not found", zap.String("ctx", ctx.String()))
		return ctx.Status(fiber.StatusNotFound).JSON(response.New(false, "API endpoint not found"))
	})

	const staticRoot = "./static"
	app.Static("/", staticRoot, fiber.Static{
		Compress:  true,
		ByteRange: true,
		Index:     "index.html",
	})

	app.Use("*", func(ctx *fiber.Ctx) error {
		if ctx.Method() != fiber.MethodGet {
			log.Instance.Warn("Global route not found", zap.String("ctx", ctx.String()))
			return ctx.Status(fiber.StatusNotFound).JSON(response.New(false, "Resource not found"))
		}

		path := strings.TrimPrefix(ctx.Path(), "/")
		base := filepath.Join(staticRoot, path)

		for {
			possibleIndex := filepath.Join(base, "index.html")
			if _, err := os.Stat(possibleIndex); err == nil {
				return ctx.SendFile(possibleIndex)
			}

			if base == staticRoot {
				break
			}

			base = filepath.Dir(base)
		}

		defaultIndex := filepath.Join(staticRoot, "index.html")
		if _, err := os.Stat(defaultIndex); err == nil {
			return ctx.SendFile(defaultIndex)
		}

		log.Instance.Warn("Global route not found", zap.String("ctx", ctx.String()))
		return ctx.Status(fiber.StatusNotFound).JSON(response.New(false, "Resource not found"))
	})
}
