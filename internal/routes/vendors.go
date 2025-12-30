package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennyclark/web-pos-api/internal/handlers"
)

func Vendors(api fiber.Router) {
	vendors := api.Group("/vendors")
	vendors.Get("/", handlers.GetAllVendors)
	vendors.Get("/:id", handlers.GetVendorById)
	vendors.Post("/", handlers.CreateVendor)
	vendors.Put("/:id", handlers.UpdateVendor)
	vendors.Delete("/:id", handlers.DeleteVendor)
}
