package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennyclark/web-pos-api/internal/handlers"
)

func Products(api fiber.Router) {
	products := api.Group("/products")
	products.Get("/", handlers.GetAllProducts)
	products.Get("/:id", handlers.GetProductById)
	products.Post("/", handlers.CreateProduct)
	products.Put("/:id", handlers.UpdateProduct)
	products.Delete("/:id", handlers.DeleteProduct)
}
