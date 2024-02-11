package models

// Define database models
type Education struct {
	ID          uint      `gorm:"primaryKey"`
	School      string    `json:"school"`
	Degree      string    `json:"degree"`
	Year        string    `json:"year"`
	Description string    `json:"description"`
	CVID        uint      `gorm:"foreignKey"`
	 // Foreign key
}

type WorkExperience struct {
	ID          uint   `gorm:"primaryKey"`
	ProjectName string `json:"projectName"`
	Description string `json:"description"`
	CVID        uint   `gorm:"foreignKey"`
}
type Skills struct{
	CVID        uint   `gorm:"foreignKey"`
	ID          uint   `gorm:"primaryKey"`
	SkillName   string `json:"skill"`

}

type CV struct {
	UserID         uint           `gorm:"foreignKey"`
	ID             uint           `gorm:"primaryKey"`
	Name           string         `json:"name"`
	LastName       string         `json:"lastname"`
	Email          string         ` json:"email"`
	Phone          string         `json:"phone"`
	AboutMe        string         `json:"aboutMe"`
	Color          string         `json:"color"`
	Skills         Skills         `json:"skills"`
	Education      Education      `json:"education"`
	WorkExperience WorkExperience `json:"workExperience"`	
}