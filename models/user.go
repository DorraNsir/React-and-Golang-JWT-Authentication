package models
 
type User struct{
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `gorm:"unique" json:"email"`
	Password []byte `json:"-"`
}