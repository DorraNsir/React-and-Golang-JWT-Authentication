package controllers
import (
	"auth-go/database"
	"auth-go/models"
	// "strconv"
	"time"
	"github.com/gofiber/fiber/v2"
	// "golang.org/x/crypto/bcrypt"
	// "github.com/dgrijalva/jwt-go"
)
func CreateCV(c *fiber.Ctx) error {
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
			Name:           data.Name,
			LastName:       data.LastName,
			Email:          data.Email,
			Phone:          data.Phone,
			Skills:         data.Skills,
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
		}
	database.DB.Create(&cv)
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
func UpdateCV(c *fiber.Ctx) error {
	id := c.Params("id")
	var cv models.CV
	if err := database.DB.First(&cv, id).Error; err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message":"CV not found",
			
		})
	}
	newCV := new(models.CV)
	if err := c.BodyParser(newCV); err != nil {
		return err
	}
	database.DB.Model(&cv).Updates(newCV)
	return c.JSON(cv)
}