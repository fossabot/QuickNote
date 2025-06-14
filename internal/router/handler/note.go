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
		nid := ctx.Params("*")
		if nid == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(false, "invalid nid"))
		}

		n := note.Note{
			NID: nid,
		}

	method:
		switch ctx.Method() {
		case fiber.MethodPost:
			// postNote NID will be ignored.
			// Always use path NID
			var postNote note.DisplayNote

			if err := ctx.BodyParser(&postNote); err != nil {
				log.Instance.Error("Failed to parse note", zap.String("nid", nid), zap.Error(err), zap.String("ctx", ctx.String()))

				return ctx.Status(fiber.StatusBadRequest).JSON(response.New("failed to parse note"))
			}

			n.Title = helper.StringToBytes(postNote.Title)
			n.Content = helper.StringToBytes(postNote.Content)

			if len(n.Title) == 0 && len(n.Content) == 0 {
				log.Instance.Warn("Empty note, override method -> DELETE", zap.String("nid", nid), zap.String("ctx", ctx.String()))

				ctx.Method(fiber.MethodDelete)

				goto method
			}

			if err := n.Write(); err != nil {
				log.Instance.Error("Failed to write note", zap.String("nid", nid), zap.Error(err), zap.String("ctx", ctx.String()))

				return ctx.Status(fiber.StatusInternalServerError).JSON(response.New("failed to write note"))
			}

			log.Instance.Info("Note updated", zap.String("nid", nid), zap.String("ctx", ctx.String()))

			return ctx.Status(fiber.StatusOK).JSON(response.New("success"))

		case fiber.MethodDelete:
			// if note not found don't return error
			if err := n.Delete(); err != nil {
				log.Instance.Warn("Failed to delete note", zap.String("nid", nid), zap.Error(err), zap.String("ctx", ctx.String()))

				return ctx.Status(fiber.StatusInternalServerError).JSON(response.New("failed to delete note"))
			}

			log.Instance.Info("Note deleted", zap.String("nid", nid), zap.String("ctx", ctx.String()))

			return ctx.Status(fiber.StatusOK).JSON(response.New("success"))

		default:
			msg := "success"

			if err := n.Read(); err != nil {
				log.Instance.Warn("Note not found", zap.String("nid", nid), zap.Error(err), zap.String("ctx", ctx.String()))

				msg = "note not found"
			}

			log.Instance.Info("Note found", zap.String("nid", nid), zap.String("ctx", ctx.String()))

			return ctx.Status(fiber.StatusOK).JSON(response.New(msg, note.DisplayNote{
				NID:     n.NID,
				Title:   helper.BytesToString(n.Title),
				Content: helper.BytesToString(n.Content),
			}))
		}
	}
}
