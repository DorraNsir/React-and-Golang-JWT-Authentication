package controllers

import (
	"auth-go/database"
	"auth-go/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	// var data map[string]string
	var data models.User
	
	if err :=c.BodyParser(&data);err!= nil{
		return err
	}
	password,_ := bcrypt.GenerateFromPassword([]byte(data.Password),14)
	user:=models.User{
		Name:data.Name,
		Email:data.Email,
		Password : password,
	}
	database.DB.Create(&user)
    return c.JSON(user)
}
func Login (c *fiber.Ctx)error{
	// var data map[string]string
	var data models.User
	
	if err :=c.BodyParser(&data);err!= nil{
		return err
	}
	var user models.User
	database.DB.Where("email = ?",data.Email).First(&user)
	if user.Id == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message":"user not found",
		})
	}
	if err:= bcrypt.CompareHashAndPassword(user.Password,[]byte(data.Password));err!=nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message":"incorrect password",
		})
	
	}
	return c.JSON(user)
}