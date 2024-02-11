package routes

import (
	"auth-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    
    app.Post("/api/register",controllers.Register)
	app.Post("/api/login",controllers.Login)
	app.Get("/api/user",controllers.User)
	app.Post("/api/logout",controllers.Logout)
	// app.Post("/api/cv", controllers.CreateCV)
	app.Get("/api/cv/:id", controllers.GetCV)
	app.Patch("/api/cv/:UserID", controllers.UpdateCV)
	app.Delete("/api/cv/:id", controllers.DeleteCV)

}