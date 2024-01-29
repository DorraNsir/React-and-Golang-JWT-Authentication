package main

import (

	"auth-go/database"
	"auth-go/routes"
	"github.com/gofiber/fiber/v2"

)

func main() {
    database.Connect()
    app:=fiber.New()
    routes.Setup(app)
    //: This line instructs the web server to start listening on port 8000.
    app.Listen(":8000")
	
}