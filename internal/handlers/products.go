package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kennyclark/web-pos-api/internal/db"
	"github.com/kennyclark/web-pos-api/internal/db/sqlc"
)

func GetAllProducts(c *fiber.Ctx) error {
	products, err := db.Q.GetAllProducts(c.Context())
	if err != nil {
		log.Printf("Unable to get all products: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to get all products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    products,
		"message": "Products retrieved successfully",
	})
}

func GetProductById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Unable to convert ID to int: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	product, err := db.Q.GetProductById(c.Context(), int32(id))
	if err != nil {
		log.Printf("Unable to get product by ID: %v\n", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    product,
		"message": "Product retrieved successfully",
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product sqlc.CreateProductParams
	if err := c.BodyParser(&product); err != nil {
		log.Printf("Unable to parse request body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdProduct, err := db.Q.CreateProduct(c.Context(), product)
	if err != nil {
		log.Printf("Unable to create product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    createdProduct,
		"message": "Product created successfully",
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	var product sqlc.UpdateProductParams
	if err := c.BodyParser(&product); err != nil {
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

	if int32(id) != product.ID {
		log.Printf("ID mismatch: %d != %d\n", id, product.ID)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID mismatch",
		})
	}

	updatedProduct, err := db.Q.UpdateProduct(c.Context(), product)
	if err != nil {
		log.Printf("Unable to update product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    updatedProduct,
		"message": "Product updated successfully",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Unable to convert ID to int: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := db.Q.DeleteProduct(c.Context(), int32(id)); err != nil {
		log.Printf("Unable to delete product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to delete product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
