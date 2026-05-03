package handlers

import (
	"strings"

	"go-fiber-crud/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TaskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{db: db}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var input models.TaskRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}

	title := strings.TrimSpace(input.Title)
	description := strings.TrimSpace(input.Description)
	if title == "" || description == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "title and description are required"})
	}

	task := models.Task{Title: title, Description: description}
	if input.Completed != nil {
		task.Completed = *input.Completed
	}

	if err := h.db.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to create task"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "task created", "data": task})
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	if err := h.db.Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to fetch tasks"})
	}

	return c.JSON(fiber.Map{"message": "success", "data": tasks})
}

func (h *TaskHandler) GetTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid task id"})
	}

	var task models.Task
	if err := h.db.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "task not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to fetch task"})
	}

	return c.JSON(fiber.Map{"message": "success", "data": task})
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid task id"})
	}

	var task models.Task
	if err := h.db.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "task not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to fetch task"})
	}

	var input models.TaskRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}

	title := strings.TrimSpace(input.Title)
	description := strings.TrimSpace(input.Description)
	if title == "" || description == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "title and description are required"})
	}

	task.Title = title
	task.Description = description
	if input.Completed != nil {
		task.Completed = *input.Completed
	}

	if err := h.db.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to update task"})
	}

	return c.JSON(fiber.Map{"message": "task updated", "data": task})
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid task id"})
	}

	var task models.Task
	if err := h.db.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "task not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to fetch task"})
	}

	if err := h.db.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to delete task"})
	}

	return c.JSON(fiber.Map{"message": "task deleted"})
}