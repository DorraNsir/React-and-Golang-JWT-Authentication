package controllers
import (
	"auth-go/database"
	"auth-go/models"
	// "strconv"
	"time"
	"fmt"
	// "gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
	// "golang.org/x/crypto/bcrypt"
	// "github.com/dgrijalva/jwt-go"
)
func CreateCV(c *fiber.Ctx, id uint,name string,email string) error {
	var data models.CV
	if err := c.BodyParser(&data); err != nil {
		return err
	}
		
	// Check if Education Year is provided, if not, set it to an empty string
	educationYear := ""
	if data.Education.Year != "" {
		// If Education Year is provided, parse it
		yearTime, err := time.Parse("2006-01-02", data.Education.Year)
		if err != nil {
			return err
		}
		educationYear = yearTime.Format("2006-01-02")
	}

		cv := models.CV{
			UserID: 		id,
			Name:           name,
			LastName:       data.LastName,
			Email:          email,
			Phone:          data.Phone,
			AboutMe:        data.AboutMe,
			Color:          data.Color,
			Education: models.Education{
				School:      data.Education.School,
				Degree:      data.Education.Degree,
				Year:        educationYear , // Assign parsed time.Time object
				Description: data.Education.Description,
			},
			WorkExperience:models.WorkExperience{
				ProjectName: data.WorkExperience.ProjectName,
				Description: data.WorkExperience.Description,
			},
			Skills:models.Skills{
				SkillName:data.Skills.SkillName,
			},
		}
	database.DB.Create(&cv)
	fmt.Println("hi")
	return c.JSON(cv)
}
func GetCV(c *fiber.Ctx) error {
	// Extracting UserID from Request Parameters
	userID := c.Params("userID")

	var cv models.CV

	// Using Where method to filter records based on UserID
	database.DB.Preload("Education").Preload("WorkExperience").Where("user_id = ?", userID).First(&cv)

	// Check if no CV found for the provided UserID
	if cv.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "CV not found",
		})
	}

	return c.JSON(cv)
}
// Handler to update a CV by ID
// func UpdateCV(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var cv models.CV
// 	if err := database.DB.First(&cv, id).Error; err != nil {
// 		c.Status(fiber.StatusNotFound)
// 		return c.JSON(fiber.Map{
// 			"message":"CV not found",
// 		})
// 	}
// 	newCV := new(models.CV)
// 	if err := c.BodyParser(newCV); err != nil {
// 		return err
// 	}
// 	database.DB.Model(&cv).Updates(newCV)
// 	return c.JSON(cv)
// }
// Handler to update a CV by user ID
func UpdateCV(c *fiber.Ctx) error {
	// Extract the user ID from the request parameters
	userID := c.Params("userID")

	// Retrieve the CV associated with the given user ID
	var cv models.CV
	if err := database.DB.Where("user_id = ?", userID).First(&cv).Error; err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "CV not found",
		})
	}

	// Parse the request body to get the updated CV data
	updatedCV := new(models.CV)
	if err := c.BodyParser(updatedCV); err != nil {
		return err
	}

	// Update the retrieved CV with the new data
	database.DB.Model(&cv).Updates(updatedCV)

	return c.JSON(cv)
}

func DeleteCV(c *fiber.Ctx) error {
	id := c.Params("id")
    // Delete associated Skills
    if err := database.DB.Where("cv_id = ?", id).Delete(&models.Skills{}).Error; err != nil {
        return err
    }

    // Delete associated Education
    if err := database.DB.Where("cv_id = ?", id).Delete(&models.Education{}).Error; err != nil {
        return err
    }

    // Delete associated WorkExperience
    if err := database.DB.Where("cv_id = ?", id).Delete(&models.WorkExperience{}).Error; err != nil {
        return err
    }

    // Finally, delete the CV
    if err := database.DB.Delete(&models.CV{}, id).Error; err != nil {
        return err
    }

    return c.SendString("CV deleted")
}
// func CreateCV(c *fiber.Ctx,id uint,name string,email string) error {
// 	var data models.CV
// 	if err := c.BodyParser(&data); err != nil {
// 		return err
// 	}

// 	// Check if Education Year is provided, if not, set it to an empty string
// 	educationYear := ""
// 	if data.Education.Year != "" {
// 		// If Education Year is provided, parse it
// 		yearTime, err := time.Parse("2006-01-02", data.Education.Year)
// 		if err != nil {
// 			return err
// 		}
// 		educationYear = yearTime.Format("2006-01-02")
// 	}
// 	skills:=models.Skills{
// 		SkillName: data.Skills.SkillName,
// 		CVID: data.ID,
// 	}
// 	education:=models.Education{
// 		School:      data.Education.School,
// 		Degree:      data.Education.Degree,
// 		Year:        educationYear, // Assign parsed time.Time object or empty string
// 		Description: data.Education.Description,
// 		CVID: data.ID,
// 	}
// 	work:=models.WorkExperience{
// 		ProjectName: data.WorkExperience.ProjectName,
// 		Description: data.WorkExperience.Description,
// 		CVID: data.ID,
// 	}

// 	cv := models.CV{
// 		UserID:         id,
// 		Name:           name,
// 		LastName:       data.LastName,
// 		Email:          email,
// 		Phone:          data.Phone,
// 		AboutMe:        data.AboutMe,
// 		Color:          data.Color,
// 		Education: 		education,
// 		WorkExperience: work,
// 		Skills: 		skills,
// 	}
// 	if err := database.DB.Create(&cv).Error; err != nil {
// 		return err
// 	}
// 	if err := database.DB.Create(&work).Error; err != nil {
// 		return err
// 	}
// 	if err := database.DB.Create(&skills).Error; err != nil {
// 		return err
// 	}
// 	if err := database.DB.Create(&education).Error; err != nil {
// 		return err
// 	}

// 	return c.JSON(cv)
// }