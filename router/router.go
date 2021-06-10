package router

import (
	"github.com/charlesonunze/busha-test/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", handler.Hello)

	// Movies Endpoints
	movies := api.Group("/movies")
	movies.Get("/", handler.GetMovies)
	movies.Post("/:movieId/comments", handler.AddCommentToMovie)
	movies.Get("/:movieId/comments", handler.ListCommentsForMovie)
	movies.Get("/:movieId/characters", handler.ListCharactersForMovie)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Resource not found.",
			"data":    nil,
		})
	})
}
