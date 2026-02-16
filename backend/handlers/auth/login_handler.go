package auth

import (
	"net/http"

	"github.com/ahmadammarm/courses-management/backend/db"
	"github.com/ahmadammarm/courses-management/backend/dto"
	"github.com/ahmadammarm/courses-management/backend/models"
	"github.com/ahmadammarm/courses-management/backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func InstructorLoginHandler(context *gin.Context) {

	var request = dto.InstructorLoginRequest{}
	var instructor = models.Instructor{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Error validating request",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	database := db.Database

	if err := database.Where("email = ?", request.Email).First(&request).Error; err != nil {
		context.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Invalid email or password",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(request.Password), []byte(request.Password)); err != nil {
		context.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Invalid email or password",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	accessToken := utils.JWTGenerate(instructor.Username)

	context.JSON(http.StatusOK, utils.SuccessResponse{
		Success: true,
		Status:  "success",
		Message: "Instructor logged in successfully",
		Data: dto.InstructorResponse{
			ID:        instructor.ID,
			Name:      instructor.Name,
			Username:  instructor.Username,
			Email:     instructor.Email,
			Expertise: instructor.Expertise,
			Token:     &accessToken,
		},
	})
}
