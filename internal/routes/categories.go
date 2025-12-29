package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennyclark/web-pos-api/internal/handlers"
)

func Categories(api fiber.Router) {
	categories := api.Group("/categories")
	categories.Get("/", handlers.GetAllCategories)
	categories.Get("/:id", handlers.GetCategoryById)
	categories.Post("/", handlers.CreateCategory)
	categories.Put("/:id", handlers.UpdateCategory)
	categories.Delete("/:id", handlers.DeleteCategory)
}
