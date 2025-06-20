package notfound

import (
	"strings"

	"github.com/Sn0wo2/QuickNote/pkg/common"
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Setup(msg string, router fiber.Router) {
	if msg = strings.ToLower(msg); msg == "" {
		msg = "page not found"
	}

	router.Use("*", func(ctx *fiber.Ctx) error {
		log.Instance.Warn(common.TitleCase(msg),
			zap.String("ctx", common.FiberContextString(ctx)))

		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": msg,
		})
	})
}
