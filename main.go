package main

import (

	"auth-go/database"
	"auth-go/routes"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"

)

func main() {
    database.Connect()
    app:=fiber.New()
    
    app.Use(cors.New(cors.Config{
        AllowCredentials: true,
    }))
    routes.Setup(app)
    //: This line instructs the web server to start listening on port 8000.
    app.Listen(":8000")
	
}