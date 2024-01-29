package routes

import (
	
	"github.com/gofiber/fiber/v3"
	
)

func main() {
    

	// this Go program creates a simple web server using the Fiber framework,
    app := fiber.New()//This line creates a new instance of the Fiber web framework. Fiber is a web framework for Go that is designed to be fast and easy to use.

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
}