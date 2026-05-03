package routes

import (
	"go-fiber-crud/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterTaskRoutes(app *fiber.App, db *gorm.DB) {
	handler := handlers.NewTaskHandler(db)
	api := app.Group("/api")
	tasks := api.Group("/tasks")

	tasks.Post("/", handler.CreateTask)
	tasks.Get("/", handler.GetTasks)
	tasks.Get("/:id", handler.GetTask)
	tasks.Put("/:id", handler.UpdateTask)
	tasks.Delete("/:id", handler.DeleteTask)
}