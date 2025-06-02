package handler

import (
	"strings"

	"github.com/Sn0wo2/QuickNote/log"
	"github.com/Sn0wo2/QuickNote/note"
	"github.com/Sn0wo2/QuickNote/response"
	"github.com/Sn0wo2/QuickNote/setup"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Note(path string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		path := strings.TrimPrefix(strings.TrimPrefix(c.Path(), setup.APIVersion), path)
		if path == "" {
			return c.Status(fiber.StatusNotFound).JSON(response.New(false, "Note not found"))
		}
		n := note.Note{ID: path}
		switch c.Method() {
		case fiber.MethodPost:
			body := c.Query("title")
			if body == "" {
				return c.Status(fiber.StatusBadRequest).JSON(response.New(false, "title is required"))
			}
			n.Title = []byte(body)

			body = c.Query("content")
			if body == "" {
				return c.Status(fiber.StatusBadRequest).JSON(response.New(false, "content is required"))
			}
			n.Content = []byte(body)

			if err := n.Write(); err != nil {
				return err
			}
			log.Instance.Info("Note created",
				zap.String("path", path),
				zap.String("ip", c.IP()),
			)
			return c.Status(fiber.StatusOK).JSON(response.New(true, "success"))
		case fiber.MethodDelete:
			if err := n.Delete(); err != nil {
				// log.Instance.Warn("Note not found", zap.Error(err))
				return c.Status(fiber.StatusNotFound).JSON(response.New(false, "Note not found"))
			}
			log.Instance.Info("Note deleted",
				zap.String("path", path),
				zap.String("ip", c.IP()),
			)
			return c.Status(fiber.StatusOK).JSON(response.New(true, "success"))
		default:
			if err := n.Read(); err != nil {
				// log.Instance.Warn("Note not found", zap.Error(err))
				return c.Status(fiber.StatusNotFound).JSON(response.New(false, "Note not found"))
			}
			/*log.Instance.Info("Note found",
				zap.String("path", path),
				zap.String("ip", c.IP()),
			)*/
			return c.Status(fiber.StatusOK).JSON(response.New(true, "success", note.DisplayNote{
				Title:   string(n.Title),
				Content: string(n.Content),
			}))
		}
	}
}
