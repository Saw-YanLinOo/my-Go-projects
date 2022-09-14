package books

import (
	"github.com/YanYanUcstt/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBook(c *fiber.Ctx) error {

	id := c.Params("id")

	var books []models.Book

	if result := h.DB.Find(&books, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&books)
}
