package handler

import (
	"io"
	"strings"

	"github.com/Sn0wo2/QuickNote/internal/note"
	"github.com/Sn0wo2/QuickNote/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func Import() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		fileHeader, err := ctx.FormFile("import")
		if err != nil {
			return err
		}
		fileName := fileHeader.Filename
		if !strings.HasSuffix(fileName, ".qnote") {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New("invalid file format"))
		}
		file, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		result, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		n := note.Note{NID: strings.TrimSuffix(fileName, ".qnote")}
		if err = n.Decode(result); err != nil {
			return err
		}
		if err = n.Write(); err != nil {
			return err
		}
		return ctx.JSON(response.New("success"))
	}
}
