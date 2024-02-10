package models
 
type User struct{
	Id uint         `json:"id"`
	Name string     `json:"name"`
	LastName string `json:"lastname"`
	Email string    `gorm:"unique" json:"email"`
	Password []byte `json:"-"`
}