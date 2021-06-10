package handler

import (
	"github.com/charlesonunze/busha-test/services"
	"github.com/gofiber/fiber/v2"
)

func GetMovies(c *fiber.Ctx) error {
	movies, err := services.GetMovies()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Movies list",
		"data":    movies,
	})
}

func AddCommentToMovie(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Comment created successfully.",
		"data":    "comment",
	})
}

func ListCommentsForMovie(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Movies found",
		"data":    "comments",
	})
}

func ListCharactersForMovie(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Characters list",
		"data":    "characters",
	})
}
