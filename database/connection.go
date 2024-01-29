package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func Connect( ){
	_ , err:= gorm.Open(mysql.Open("root:@/Auth-go"),&gorm.Config{})
    if err != nil{
        panic("could not connect to the database")
    }
    
}