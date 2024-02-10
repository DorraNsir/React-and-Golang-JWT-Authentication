package database

import (
	"auth-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect( ){
	connection, err:= gorm.Open(mysql.Open("root:@/Auth-go"),&gorm.Config{})
    if err != nil{
        panic("could not connect to the database" + err.Error())
    }
    DB=connection
    
    connection.AutoMigrate(&models.User{})
    connection.AutoMigrate(&models.CV{}, &models.Education{}, &models.WorkExperience{})
}