package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennyclark/web-pos-api/internal/handlers"
)

func Customers(api fiber.Router) {
	customers := api.Group("/customers")
	customers.Get("/", handlers.GetAllCustomers)
	customers.Get("/:id", handlers.GetCustomerById)
	customers.Post("/", handlers.CreateCustomer)
	customers.Put("/:id", handlers.UpdateCustomer)
	customers.Delete("/:id", handlers.DeleteCustomer)
}
