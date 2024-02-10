package models

// Define database models
type Education struct {
	ID          uint      `gorm:"primaryKey"`
	School      string    `json:"school"`
	Degree      string    `json:"degree"`
	Year        string    `json:"year"`
	Description string    `json:"description"`
	CVID        uint   // Foreign key
}

type WorkExperience struct {
	ID          uint   `gorm:"primaryKey"`
	ProjectName string `json:"projectName"`
	Description string `json:"description"`
	CVID        uint   // Foreign key
}

type CV struct {
	ID             uint           `gorm:"primaryKey"`
	Name           string         `json:"name"`
	LastName       string         `json:"lastname"`
	Email          string         `gorm:"unique" json:"email"`
	Phone          string         `json:"phone"`
	Skills         string         `json:"skills"`
	AboutMe        string         `json:"aboutMe"`
	Color          string         `json:"color"`
	Education      Education      `json:"education"`
	WorkExperience WorkExperience `json:"workExperience"`
	
}