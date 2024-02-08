package models

import (
	"time"
)

type Education struct {
	School    string `json:school`
	Degree    string `json:degree`
	Year uint 
}
type WorkExperience struct{
	jobTitle string
	company string
	startDate time.Time
	endDate time.Time
}

type CV struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `gorm:"unique" json:"email"`
	phone    string `json:"phone"`
	skills string 
	eduction Education
	workExperience WorkExperience

}