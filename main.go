package main

import (
	"log"

	"fiber.comexample/controllers"
	"fiber.comexample/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	//app.Use(logger.New())

	micro.Route("/notes", func(router fiber.Router) {
		router.Post("/", controllers.CreateNoteHandler)
		router.Get("", controllers.FindNotes)
	})
	micro.Route("/notes/:noteId", func(router fiber.Router) {
		router.Delete("", controllers.DeleteNote)
		router.Get("", controllers.FindNoteById)
		router.Patch("", controllers.UpdateNote)
	})

	micro.Get("healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, SQlite and Gorm",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
