package auth

import (
	"net/http"

	"github.com/ahmadammarm/courses-management/backend/db"
	"github.com/ahmadammarm/courses-management/backend/dto"
	"github.com/ahmadammarm/courses-management/backend/models"
	"github.com/ahmadammarm/courses-management/backend/utils"
	"github.com/gin-gonic/gin"
)

func InstructorRegisterHandler(context *gin.Context) {

	var request = dto.InstructorCreateRequest{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Error validating request",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	instructor := models.Instructor{
		Name:      request.Name,
		Username:  request.Username,
		Email:     request.Email,
		Password:  utils.PasswordHash(request.Password),
		Expertise: request.Expertise,
	}

	database := db.Database

	if err := database.Create(&instructor).Error; err != nil {
		if utils.IsDuplicateEntryError(err) {
			context.JSON(http.StatusConflict, utils.ErrorResponse{
				Success: false,
				Status:  "error",
				Message: "Instructor already exists",
				Errors:  utils.ErrorMessageTranslate(err),
			})
		} else {
			context.JSON(http.StatusInternalServerError, utils.ErrorResponse{
				Success: false,
				Status:  "error",
				Message: "Failed to create instructor",
			})
		}
		return
	}

	context.JSON(http.StatusCreated, utils.SuccessResponse{
		Success: true,
		Status:  "success",
		Message: "Instructor registered successfully",
		Data: dto.InstructorResponse{
			ID:        instructor.ID,
			Name:      instructor.Name,
			Username:  instructor.Username,
			Email:     instructor.Email,
			Expertise: instructor.Expertise,
		},
	})
}
