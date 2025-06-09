package errorhandler

import (
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/Sn0wo2/QuickNote/pkg/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Error(ctx *fiber.Ctx, err error) error {
	log.Instance.Error(err.Error(),
		zap.Error(err),
		zap.String("ctx", ctx.String()),
	)

	return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(false, "Oops, something went wrong"))
}
