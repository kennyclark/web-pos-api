package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kennyclark/web-pos-api/internal/db"
	"github.com/kennyclark/web-pos-api/internal/db/sqlc"
)

func GetAllCustomers(c *fiber.Ctx) error {
	customers, err := db.Q.GetAllCustomers(c.Context())
	if err != nil {
		log.Printf("Unable to get all customers: %v\n", err)

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    customers,
		"message": "Customers retrieved successfully",
	})
}

func GetCustomerById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Unable to convert ID to int: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	customer, err := db.Q.GetCustomerById(c.Context(), int32(id))
	if err != nil {
		log.Printf("Unable to get vendor by ID: %v\n", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    customer,
		"message": "Customer successfully retrieved",
	})
}

func CreateCustomer(c *fiber.Ctx) error {
	var customer sqlc.CreateCustomerParams
	if err := c.BodyParser(&customer); err != nil {
		log.Printf("Unable to parse request body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdCustomer, err := db.Q.CreateCustomer(c.Context(), customer)
	if err != nil {
		log.Printf("Unable to create customer: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create customer",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    createdCustomer,
		"message": "Customer created successfully",
	})
}

func UpdateCustomer(c *fiber.Ctx) error {
	var customer sqlc.UpdateCustomerParams
	if err := c.BodyParser(&customer); err != nil {
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

	if int32(id) != customer.ID {
		log.Printf("ID mismatch: %d != %d\n", id, customer.ID)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID mismatch",
		})
	}

	updatedCustomer, err := db.Q.UpdateCustomer(c.Context(), customer)
	if err != nil {
		log.Printf("Unable to update customer: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update customer",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    updatedCustomer,
		"message": "Customer updated successfully",
	})
}

func DeleteCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Unable to convert ID to int: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := db.Q.DeleteCustomer(c.Context(), int32(id)); err != nil {
		log.Printf("Unable to delete customer: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to delete customer",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Customer deleted successfully",
	})
}
