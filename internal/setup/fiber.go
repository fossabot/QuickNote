package setup

import (
	"time"

	"github.com/Sn0wo2/QuickNote/internal/router/errorhandler"
	"github.com/gofiber/fiber/v2"
)

const APIVersion = "/v1/"

func Fiber() *fiber.App {
	return fiber.New(fiber.Config{
		AppName:               "QuickNote",
		CaseSensitive:         true,
		DisableStartupMessage: false,
		ErrorHandler:          errorhandler.Error,
		IdleTimeout:           5 * time.Second,
		Prefork:               true,
		ReadTimeout:           10 * time.Second,
		ReduceMemoryUsage:     true,
		StrictRouting:         true,
		WriteTimeout:          10 * time.Second,
	})
}
