
package controllers
import (
	"auth-go/database"
	"auth-go/models"
	// "strconv"
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
			Name:          name,
			LastName:      "test",
			Email:         email,
			Phone:         "test",
			AboutMe:       "test",
			Color:         "test",
			Skills:        data.Skills,         // Assign skills from data
			WorkExperience: data.WorkExperience, // Assign work experiences from data
			Education: models.Education{
				School:      "test",
				Degree:      "test",
				Year:        "test", // Assign parsed time.Time object
				Description: "test",
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
	fmt.Println("HELOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
    // Append new skills and work experiences to the existing CV
    for _, newSkill := range updatedCV.Skills {
        // Check if the skill already exists
        found := false
        for _, existingSkill := range cv.Skills {
            if existingSkill.ID == newSkill.ID {
                existingSkill.SkillName = newSkill.SkillName
                found = true
                break
            }

        }
        if !found {
            // If the skill doesn't exist, append it
            cv.Skills = append(cv.Skills, newSkill)
			fmt.Println(cv.Skills)
        }
    }

    // for _, newWorkExp := range updatedCV.WorkExperience {
    //     cv.WorkExperience = append(cv.WorkExperience, newWorkExp)
    // }
    // Update the top-level fields of the retrieved CV with the new data
    database.DB.Model(&cv).Updates(updatedCV)

    // Update nested fields (Education) individually
    if err := database.DB.Model(&cv.Education).Updates(updatedCV.Education).Error; err != nil {
        return err
    }

    // Save the updated CV back to the database
    if err := database.DB.Save(&cv).Error; err != nil {
        return err
    }

    return c.JSON(cv)
}

//don't remove this 
// func UpdateCV(c *fiber.Ctx) error {
//     // Extract the user ID from the request parameters
//     userID := c.Params("userID")

//     // Retrieve the CV associated with the given user ID
//     var cv models.CV
//     if err := database.DB.Preload("Education").Preload("WorkExperience").Preload("Skills").Where("user_id = ?", userID).First(&cv).Error; err != nil {
//         c.Status(fiber.StatusNotFound)
//         return c.JSON(fiber.Map{
//             "message": "CV not found",
//         })
//     }

//     // Parse the request body to get the updated CV data
//     updatedCV := new(models.CV)
//     if err := c.BodyParser(updatedCV); err != nil {
//         return err
//     }

//     // Update the top-level fields of the retrieved CV with the new data
//     database.DB.Model(&cv).Updates(updatedCV)

//     // Update nested fields (Education, WorkExperience, Skills) individually
//     if err := database.DB.Model(&cv.Education).Updates(updatedCV.Education).Error; err != nil {
//         return err
//     }
//     if err := database.DB.Model(&cv.WorkExperience).Updates(updatedCV.WorkExperience).Error; err != nil {
//         return err
//     }
//     if err := database.DB.Model(&cv.Skills).Updates(updatedCV.Skills).Error; err != nil {
//         return err
//     }

//     return c.JSON(cv)
// }


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

