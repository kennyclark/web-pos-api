package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kennyclark/web-pos-api/internal/db"
	"github.com/kennyclark/web-pos-api/internal/db/sqlc"
)

func GetAllVendors(c *fiber.Ctx) error {
	vendors, err := db.Q.GetAllVendors(c.Context())
	if err != nil {
		log.Printf("Unable to get all vendors: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to get all vendors",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    vendors,
		"message": "Vendors retrieved successfully",
	})
}

func GetVendorById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Unable to convert ID to int: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	vendor, err := db.Q.GetVendorById(c.Context(), int32(id))
	if err != nil {
		log.Printf("Unable to get vendor by ID: %v\n", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Vendor not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    vendor,
		"message": "Vendor retrieved successfully",
	})
}

func CreateVendor(c *fiber.Ctx) error {
	var vendor sqlc.CreateVendorParams
	if err := c.BodyParser(&vendor); err != nil {
		log.Printf("Unable to parse request body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdVendor, err := db.Q.CreateVendor(c.Context(), vendor)
	if err != nil {
		log.Printf("Unable to create vendor: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create vendor",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    createdVendor,
		"message": "Vendor created successfully",
	})
}

func UpdateVendor(c *fiber.Ctx) error {
	var vendor sqlc.UpdateVendorParams
	if err := c.BodyParser(&vendor); err != nil {
		log.Printf("Unable to parse request body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Unable to convert ID to int: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if int32(id) != vendor.ID {
		log.Printf("ID mismatch: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID mismatch",
		})
	}

	updatedVendor, err := db.Q.UpdateVendor(c.Context(), vendor)
	if err != nil {
		log.Printf("Unable to update vendor: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update vendor",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    updatedVendor,
		"message": "Vendor updated successfully",
	})
}

func DeleteVendor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Unable to convert ID to int: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := db.Q.DeleteVendor(c.Context(), int32(id)); err != nil {
		log.Printf("Unable to delete vendor: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to delete vendor",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Vendor deleted successfully",
	})
}
