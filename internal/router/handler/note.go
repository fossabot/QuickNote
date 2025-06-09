package handler

import (
	"github.com/Sn0wo2/QuickNote/internal/note"
	"github.com/Sn0wo2/QuickNote/pkg/helper"
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/Sn0wo2/QuickNote/pkg/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Note() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(false, "invalid note id"))
		}

		n := note.Note{
			NID: id,
		}

		switch ctx.Method() {
		case fiber.MethodPost:
			var postNote note.DisplayNote

			if err := ctx.BodyParser(&postNote); err != nil {
				log.Instance.Error("Failed to parse note", zap.Error(err), zap.String("ctx", ctx.String()))

				return ctx.Status(fiber.StatusBadRequest).JSON(response.New(false, "failed to parse note"))
			}

			n.Title = helper.StringToBytes(postNote.Title)
			n.Content = helper.StringToBytes(postNote.Content)

			if err := n.Write(); err != nil {
				log.Instance.Error("Failed to write note", zap.Error(err), zap.String("ctx", ctx.String()))

				return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(false, "failed to write note"))
			}

			log.Instance.Info("Note updated", zap.String("ctx", ctx.String()))

			return ctx.Status(fiber.StatusOK).JSON(response.New(true, "success"))

		case fiber.MethodDelete:
			// error. if not found dont return error
			if err := n.Delete(); err != nil {
				log.Instance.Warn("Failed to delete note", zap.Error(err), zap.String("ctx", ctx.String()))

				return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(false, "Failed to delete note"))
			}

			log.Instance.Info("Note deleted", zap.String("ctx", ctx.String()))

			return ctx.Status(fiber.StatusOK).JSON(response.New(true, "success"))

		default:
			msg := "success"
			if err := n.Read(); err != nil {
				log.Instance.Warn("Note not found", zap.Error(err), zap.String("ctx", ctx.String()))
				msg = "Note not found"
			}

			log.Instance.Info("Note found", zap.String("ctx", ctx.String()))

			return ctx.Status(fiber.StatusOK).JSON(response.New(true, msg, note.DisplayNote{
				NID:     n.NID,
				Title:   helper.BytesToString(n.Title),
				Content: helper.BytesToString(n.Content),
			}))
		}
	}
}
