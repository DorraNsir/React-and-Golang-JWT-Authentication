package main

import (
	"auth-go/database"
	"auth-go/"
)

func main() {
    database.Connect()


    //: This line instructs the web server to start listening on port 8000.
    app.Listen(":8000")
	
}