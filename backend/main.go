package main

import (
	"log"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/matoous/go-nanoid/v2"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
	app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			nid, _ := gonanoid.New()
			return nid
		},
	})) 
	


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello chi")
	})

	app.Get("/tasks", tasks)
	app.Post("/task", createTask)
	app.Patch("/task/:taskId", updateTask)
	app.Delete("/task/:taskId", deleteTask)
	log.Fatal(app.Listen(":9000"))
}
