package handler

import (
	"github.com/Sn0wo2/QuickNote/internal/note"
	"github.com/Sn0wo2/QuickNote/pkg/common"
	"github.com/Sn0wo2/QuickNote/pkg/helper"
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/Sn0wo2/QuickNote/pkg/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Note() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		nid := ctx.Params("*")
		if nid == "" {
			log.Instance.Warn("Invalid nid",
				zap.String("ctx", common.FiberContextString(ctx)))

			return ctx.Status(fiber.StatusBadRequest).JSON(response.New("invalid nid"))
		}

		n := note.Note{
			NID: nid,
		}

	method:
		switch ctx.Method() {
		case fiber.MethodPost:
			if err := ctx.BodyParser(&n); err != nil {
				log.Instance.Error("Failed to parse note",
					zap.String("nid", nid),
					zap.Error(err),
					zap.String("ctx", common.FiberContextString(ctx)))

				return ctx.Status(fiber.StatusBadRequest).JSON(response.New("failed to parse note"))
			}

			n.Title = helper.StringToBytes(n.DisplayTitle)
			n.Content = helper.StringToBytes(n.DisplayContent)

			if len(n.Title) == 0 && len(n.Content) == 0 {
				log.Instance.Warn("Empty note, override method -> DELETE",
					zap.String("nid", nid),
					zap.String("ctx", common.FiberContextString(ctx)))

				ctx.Method(fiber.MethodDelete)

				goto method
			}

			if err := n.Write(); err != nil {
				log.Instance.Error("Failed to write note", zap.String("nid", nid),
					zap.Error(err),
					zap.String("ctx", common.FiberContextString(ctx)))

				return ctx.Status(fiber.StatusInternalServerError).JSON(response.New("failed to write note"))
			}

			log.Instance.Info("Note updated",
				zap.String("nid", nid),
				zap.String("ctx", common.FiberContextString(ctx)))

			return ctx.Status(fiber.StatusOK).JSON(response.New("success"))

		case fiber.MethodDelete:
			// if note not found don't return error
			if err := n.Delete(); err != nil {
				log.Instance.Warn("Failed to delete note",
					zap.String("nid", nid),
					zap.Error(err),
					zap.String("ctx", common.FiberContextString(ctx)))

				return ctx.Status(fiber.StatusInternalServerError).JSON(response.New("failed to delete note"))
			}

			log.Instance.Info("Note deleted",
				zap.String("nid", nid),
				zap.String("ctx", common.FiberContextString(ctx)))

			return ctx.Status(fiber.StatusOK).JSON(response.New("success"))

		case fiber.MethodGet:
			msg := "success"

			if err := n.Read(); err != nil {
				log.Instance.Warn("Note not found", zap.String("nid", nid),
					zap.Error(err),
					zap.String("ctx", common.FiberContextString(ctx)))

				msg = "note not found"
			}

			log.Instance.Info("Note found",
				zap.String("nid", nid),
				zap.String("ctx", common.FiberContextString(ctx)))

			n.DisplayTitle = helper.BytesToString(n.Title)
			n.DisplayContent = helper.BytesToString(n.Content)

			return ctx.Status(fiber.StatusOK).JSON(response.New(msg, n))
		// TODO: Too late for fallback method...
		default:
			log.Instance.Warn("Invalid method",
				zap.String("nid", nid),
				zap.String("ctx", common.FiberContextString(ctx)))

			return ctx.Next()
		}
	}
}
