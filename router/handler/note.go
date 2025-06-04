package handler

import (
	"strings"

	"github.com/Sn0wo2/QuickNote/helper"
	"github.com/Sn0wo2/QuickNote/log"
	"github.com/Sn0wo2/QuickNote/note"
	"github.com/Sn0wo2/QuickNote/response"
	"github.com/Sn0wo2/QuickNote/setup"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Note(path string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		notePath := strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(ctx.Path(), setup.APIVersion), path), "/")
		if notePath == "" {
			return ctx.Status(fiber.StatusNotFound).JSON(response.New(false, "Note not found"))
		}

		n := note.Note{
			NID: notePath,
			Key: helper.StringToBytes(ctx.Query("key")),
		}

		switch ctx.Method() {
		case fiber.MethodPost:
			title := ctx.Query("title")
			content := ctx.Query("content")

			if title == "" || content == "" {
				msg := "title is required"
				if content == "" {
					msg = "content is required"
				}
				return ctx.Status(fiber.StatusBadRequest).JSON(response.New(false, msg))
			}

			n.Title = helper.StringToBytes(title)
			n.Content = helper.StringToBytes(content)

			if err := n.Write(); err != nil {
				log.Instance.Error("Failed to write note", zap.Error(err), zap.String("ctx", ctx.String()))
				return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(false, "failed to write note"))
			}

			log.Instance.Info("Note created", zap.String("ctx", ctx.String()))
			return ctx.Status(fiber.StatusOK).JSON(response.New(true, "success"))

		case fiber.MethodDelete:
			if err := n.Delete(); err != nil {
				log.Instance.Warn("Note not found", zap.Error(err), zap.String("ctx", ctx.String()))
				return ctx.Status(fiber.StatusNotFound).JSON(response.New(false, err.Error()))
			}

			log.Instance.Info("Note deleted", zap.String("ctx", ctx.String()))
			return ctx.Status(fiber.StatusOK).JSON(response.New(true, "success"))

		default:
			if err := n.Read(); err != nil {
				log.Instance.Warn("Note not found", zap.Error(err), zap.String("ctx", ctx.String()))
				return ctx.Status(fiber.StatusNotFound).JSON(response.New(false, err.Error()))
			}

			log.Instance.Info("Note found", zap.String("ctx", ctx.String()))
			return ctx.Status(fiber.StatusOK).JSON(response.New(true, "success", note.DisplayNote{
				NID:     n.NID,
				Title:   helper.BytesToString(n.Title),
				Content: helper.BytesToString(n.Content),
			}))
		}
	}
}
