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
func CreateCV(c *fiber.Ctx, id uint,name string) error {
	var data models.CV
	if err := c.BodyParser(&data); err != nil {
		return err
	}
		
			// Parse Year string to time.Time
			yearTime, err := time.Parse("2006-01-02", data.Education.Year)
			if err != nil {
				// Handle error
				return err
			}
			// Convert time.Time back to string in desired format
			yearStr := yearTime.Format("2006-01-02")

		cv := models.CV{
			UserID: 		id,
			Name:           name,
			LastName:       data.LastName,
			Email:          data.Email,
			Phone:          data.Phone,
			AboutMe:        data.AboutMe,
			Color:          data.Color,
			Education: models.Education{
				School:      data.Education.School,
				Degree:      data.Education.Degree,
				Year:        yearStr , // Assign parsed time.Time object
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
	//Extracting ID from Request Parameters
	id := c.Params("id")
	var cv models.CV
	//using Preload("Education") and Preload("WorkExperience") ensures that when the CV object is retrieved from the database,
	// its associated Education and WorkExperience records are also fetched simultaneously.
	//The First method is used to retrieve the first matching record from the database based on the provided conditions (in this case, the ID)
	database.DB.Preload("Education").Preload("WorkExperience").First(&cv, id)
	if cv.ID == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message":"CV not found",
			
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