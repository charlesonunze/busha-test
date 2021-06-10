package handler

import (
	"github.com/charlesonunze/busha-test/services"
	"github.com/charlesonunze/busha-test/validators"
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
	payload := struct {
		Body string `json:"body"`
	}{}
	err := c.BodyParser(&payload)
	if err != nil {
		return err
	}

	err = validators.ValidateComment(payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	movieId := c.Params("movieId")
	mID, err := services.FindMovie(movieId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	comment, err := services.CreateComment(payload.Body, mID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Comment created successfully.",
		"data":    comment,
	})
}

func ListCommentsForMovie(c *fiber.Ctx) error {
	movieId := c.Params("movieId")
	comments, err := services.GetComments(movieId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Comments found",
		"data":    comments,
	})
}

func ListCharactersForMovie(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Characters list",
		"data":    "characters",
	})
}
