package router

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Sn0wo2/QuickNote/internal/router/handler"
	"github.com/Sn0wo2/QuickNote/internal/setup"
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/Sn0wo2/QuickNote/pkg/response"
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

	api.Group("health", handler.Health())
	api.Group("note/:id", handler.Note())

	api.Use("*", func(ctx *fiber.Ctx) error {
		log.Instance.Warn("API route not found", zap.String("ctx", ctx.String()))

		return ctx.Status(fiber.StatusNotFound).JSON(response.New(false, "api endpoint not found"))
	})

	const staticRoot = "./static"

	app.Static("/", staticRoot, fiber.Static{
		Compress:  true,
		ByteRange: true,
		// index default: index.html
		CacheDuration: 60 * time.Second,
		MaxAge:        60})

	app.Use("*", func(ctx *fiber.Ctx) error {
		if ctx.Method() != fiber.MethodGet {
			return ctx.Next()
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
		return ctx.Next()
	})

	app.Use("*", func(ctx *fiber.Ctx) error {
		log.Instance.Warn("Router not found", zap.String("ctx", ctx.String()))

		return ctx.Status(fiber.StatusNotFound).JSON(response.New(false, "router not found"))
	})
}
