package main


import (
    "log"

    "github.com/gofiber/fiber/v2"
    "go-chat/src/database"
)

const (
    PORT = ":8000"
)
func main() {
	database.Connect()
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(PORT))
}