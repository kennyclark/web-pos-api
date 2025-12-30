package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kennyclark/web-pos-api/internal/db"
	"github.com/kennyclark/web-pos-api/internal/db/sqlc"
)

func GetAllCategories(c *fiber.Ctx) error {
	categories, err := db.Q.GetAllCategories(c.Context())
	if err != nil {
		log.Printf("Unable to get all categories: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to get all categories",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    categories,
		"message": "Categories retrieved successfully",
	})
}

func GetCategoryById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Unable to convert ID to int: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	category, err := db.Q.GetCategoryById(c.Context(), int32(id))
	if err != nil {
		log.Printf("Unable to get category by ID: %v\n", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Category not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    category,
		"message": "Category retrieved successfully",
	})
}

func CreateCategory(c *fiber.Ctx) error {
	var category sqlc.Category
	if err := c.BodyParser(&category); err != nil {
		log.Printf("Unable to parse body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	createdCategory, err := db.Q.CreateCategory(c.Context(), category.Name)
	if err != nil {
		log.Printf("Unable to create category: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create category",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    createdCategory,
		"message": "Category created successfully",
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	var category sqlc.Category
	if err := c.BodyParser(&category); err != nil {
		log.Printf("Unable to parse body: %v\n", err)
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
	if int32(id) != category.ID {
		log.Printf("ID mismatch: %d != %d\n", id, category.ID)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID mismatch",
		})
	}

	updatedCategory, err := db.Q.UpdateCategory(c.Context(), sqlc.UpdateCategoryParams{
		ID:   int32(id),
		Name: category.Name,
	})
	if err != nil {
		log.Printf("Unable to update category: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    updatedCategory,
		"message": "Category updated successfully",
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Unable to convert ID to int: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := db.Q.DeleteCategory(c.Context(), int32(id)); err != nil {
		log.Printf("Unable to delete category: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to delete category",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}
