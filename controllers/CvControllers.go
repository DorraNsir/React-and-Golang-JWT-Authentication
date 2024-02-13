
package controllers
import (
	"auth-go/database"
	"auth-go/models"
	"strconv"
	// "time"
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
		
		if len(data.Skills) == 0{
			newSkill := models.Skills{
				SkillName: "test",
			}
			// database.DB.Create(&newSkill) // Save the skill to the database
			data.Skills = append(data.Skills, newSkill) // Append the skill to data.Skills
		
		}
		if len(data.WorkExperience) == 0 {
			newProject := models.WorkExperience{
				ProjectName: "test",
				Description: "test",
			}
			// database.DB.Create(&newProject) // Save the skill to the database
			data.WorkExperience = append(data.WorkExperience, newProject) // Append the skill to data.Skills
		
		}
		cv := models.CV{
			UserID:        id,
			Name:          " ",
			LastName:      " ",
			Email:         email,
			Phone:         " ",
			AboutMe:       " ",
			Color:         " ",
			Skills:        data.Skills,         // Assign skills from data
			WorkExperience: data.WorkExperience, // Assign work experiences from data
			Education: models.Education{
				School:      " ",
				Degree:      " ",
				Year:        " ", // Assign parsed time.Time object
				Description: " ",
			},
		}
	database.DB.Create(&cv)
	fmt.Println(len(data.Skills))
	return c.JSON(cv)
}

func GetCV(c *fiber.Ctx) error {
	// Extracting UserID from Request Parameters
	userID := c.Params("userID")

	var cv models.CV

	// Using Where method to filter records based on UserID
	database.DB.Preload("Education").Preload("WorkExperience").Preload("Skills").Where("user_id = ?", userID).First(&cv)

	// Check if no CV found for the provided UserID
	if cv.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "CV not found",
		})
	}

	return c.JSON(cv)
}

func UpdateCV(c *fiber.Ctx) error {
    // Extract the user ID from the request parameters
    userID := c.Params("userID")

    // Retrieve the CV associated with the given user ID
    var cv models.CV
    if err := database.DB.Preload("Education").Preload("WorkExperience").Preload("Skills").Where("user_id = ?", userID).First(&cv).Error; err != nil {
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

    // Update the top-level fields of the retrieved CV with the new data
    if err := database.DB.Model(&cv).Updates(updatedCV).Error; err != nil {
        return err
    }
	//     // Update nested fields (Education, WorkExperience, Skills) individually
    if err := database.DB.Model(&cv.Education).Updates(updatedCV.Education).Error; err != nil {
        return err
    }

    // Update nested fields (Skills) individually
    if updatedCV.Skills != nil {
        for i := range updatedCV.Skills {
            if i < len(cv.Skills) {
                fmt.Println("Updating skill:", updatedCV.Skills[i])
                if err := database.DB.Model(&cv.Skills[i]).Updates(&updatedCV.Skills[i]).Error; err != nil {
                    return err
                }
            } else {
                // If the index is out of range, append the new skill
                fmt.Println("Appending new skill:", updatedCV.Skills[i])
                cv.Skills = append(cv.Skills, updatedCV.Skills[i])
            }
        }
    }
	    // Update nested fields (WorkExperience) individually
		if updatedCV.WorkExperience != nil {
			for i := range updatedCV.WorkExperience {
				if i < len(cv.WorkExperience) {
					fmt.Println("Updating work:", updatedCV.WorkExperience[i])
					if err := database.DB.Model(&cv.WorkExperience[i]).Updates(&updatedCV.WorkExperience[i]).Error; err != nil {
						return err
					}
				} else {
					// If the index is out of range, append the new skill
					fmt.Println("Appending new work:", updatedCV.WorkExperience[i])
					cv.WorkExperience= append(cv.WorkExperience, updatedCV.WorkExperience[i])
				}
			}
		}

    // Save the updated CV back to the database
    if err := database.DB.Save(&cv).Error; err != nil {
        return err
    }

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
func DeleteSkillHandler(c *fiber.Ctx) error {
    // Parse CV ID and skill ID from request parameters
    cvID := c.Params("cvID")
    skillID := c.Params("skillID")

    // Retrieve the CV from the database
    var cv models.CV
    if err := database.DB.Preload("Skills").First(&cv, cvID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).SendString("CV not found")
    }

    // Find the index of the skill in the CV's skills array
    var skillIndex int = -1
    for i, skill := range cv.Skills {
        if strconv.Itoa(int(skill.ID)) == skillID {
            skillIndex = i
            break
        }
    }

    // Check if the skill was found
    if skillIndex == -1 {
        return c.Status(fiber.StatusNotFound).SendString("Skill not found")
    }

    // Delete the skill from the database
    if err := database.DB.Delete(&cv.Skills[skillIndex]).Error; err != nil {
        return err
    }

    // Remove the skill from the CV's skills slice
    cv.Skills = append(cv.Skills[:skillIndex], cv.Skills[skillIndex+1:]...)

    // Optionally, respond with a success message
    return c.SendStatus(fiber.StatusOK)
}

