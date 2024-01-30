package controllers

import (
	"auth-go/database"
	"auth-go/models"
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)
const SecretKey="secret"
func Register(c *fiber.Ctx) error {
	// var data map[string]string
	var data models.User
	//parse the request body and decode it into the data variable
	if err :=c.BodyParser(&data);err!= nil{
		return err
	}
	password,_ := bcrypt.GenerateFromPassword([]byte(data.Password),14)
	user:=models.User{
		Name:data.Name,
		Email:data.Email,
		Password : password,
	}
	//This line uses the Create method of the database connection (database.DB) to insert the new user into the database
	database.DB.Create(&user)
    return c.JSON(user)
}
func Login (c *fiber.Ctx)error{
	// var data map[string]string
	var data models.User
	//parse the request body and decode it into the data variable
	if err :=c.BodyParser(&data);err!= nil{
		return err
	}
	var user models.User
	database.DB.Where("email = ?",data.Email).First(&user)
	//If no user is found (checked by user.Id == 0)
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
	//This function creates a new JWT token with the specified signing method
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt:time.Now().Add(time.Hour*24).Unix(),//1day
	})
	token,err:= claims.SignedString([]byte(SecretKey))
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message":"could not login",
		})
	}
	// creating a cookie containing the JWT token and sending it to the client
	cookie:= fiber.Cookie{
		//This is the name that will be used to identify this cookie on the client side.
		Name :"jwt",
		Value : token,
		Expires : time.Now().Add(time.Hour*24),
		//Makes the cookie accessible only through HTTP requests and not accessible through JavaScript. 
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message":"success",
	})
}
func User (c *fiber.Ctx)error{
	cookie := c.Cookies("jwt")
	token,err:= jwt.ParseWithClaims(cookie,&jwt.StandardClaims{},func (token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey),nil
	})
	if err != nil{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"Unauthorized",
		})
	}
	claims:= token.Claims.(*jwt.StandardClaims)
	var user models.User
	database.DB.Where("id= ?",claims.Issuer).First(&user)
	return c.JSON(user)
}
func Logout (c * fiber.Ctx)error{
	cookie:= fiber.Cookie{
		Name:"jwt",
		Value:"",
		Expires:time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message":"success",
	})
}