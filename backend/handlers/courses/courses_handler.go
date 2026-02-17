package courses

import (
	"net/http"

	"github.com/ahmadammarm/courses-management/backend/db"
	"github.com/ahmadammarm/courses-management/backend/dto"
	"github.com/ahmadammarm/courses-management/backend/models"
	"github.com/ahmadammarm/courses-management/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

func GetAllCoursesHandler(context *gin.Context) {
	var courses []models.Course

	err := db.Database.Select("id", "title", "description", "slug", "content", "status", "instructor_id", "created_at", "updated_at").Find(&courses).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Error retrieving courses",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	context.JSON(http.StatusOK, utils.SuccessResponse{
		Success: true,
		Status:  "success",
		Message: "Courses retrieved successfully",
		Data:    courses,
	})
}

func GetAllCoursesByInstructorHandler(context *gin.Context) {
	instructorID := context.GetString("instructor_id")

	if instructorID == "" {
		context.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	var courses []dto.CourseWithInstructorResponse

	err := db.Database.
		Table("courses").
		Select(`
            courses.id,
            courses.title,
            courses.description,
            courses.slug,
            courses.content,
            courses.status,
            courses.created_at,
            courses.updated_at,
            instructors.id as instructor_id,
            instructors.name as instructor_name,
            instructors.username as instructor_username
        `).
		Joins("JOIN instructors ON instructors.id = courses.instructor_id").
		Where("courses.instructor_id = ?", instructorID).
		Scan(&courses).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Failed to retrieve courses",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	context.JSON(http.StatusOK, utils.SuccessResponse{
		Success: true,
		Status:  "success",
		Message: "Courses retrieved successfully",
		Data:    courses,
	})
}

func CreateCourseHandler(context *gin.Context) {
	var request = dto.CreateCourseRequest{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Error validating request",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	instructorID := context.GetString("instructor_id")
	if instructorID == "" {
		context.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Instructor ID not found in context",
			Errors:  utils.ErrorMessageTranslate(nil),
		})
		return
	}

	slugCourse := utils.SlugGenerate(request.Title)

	var course = models.Course{
		ID:           cuid2.Generate(),
		Title:        request.Title,
		Description:  request.Description,
		Slug:         slugCourse,
		Content:      request.Content,
		Status:       request.Status,
		InstructorID: instructorID,
	}

	database := db.Database

	if err := database.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&course).Error; err != nil {
			return err
		}

		if err := tx.Preload("Instructor", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "username", "email", "expertise", "created_at", "updated_at")
		}).First(&course, "id = ?", course.ID).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		context.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Error creating course",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	context.JSON(http.StatusOK, utils.SuccessResponse{
		Success: true,
		Status:  "success",
		Message: "Course created successfully",
		Data: dto.CourseResponse{
			ID:           course.ID,
			Title:        course.Title,
			Description:  course.Description,
			Slug:         course.Slug,
			Content:      course.Content,
			Status:       course.Status,
			InstructorID: course.InstructorID,
			CreatedAt:    course.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    course.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func GetCourseByIDHandler(context *gin.Context) {
	courseID := context.Param("id")

	var course models.Course

	err := db.Database.Select("id", "title", "description", "slug", "content", "status", "instructor_id", "created_at", "updated_at").First(&course, "id = ?", courseID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, utils.ErrorResponse{
				Success: false,
				Status:  "error",
				Message: "Course not found",
			})
		} else {
			context.JSON(http.StatusInternalServerError, utils.ErrorResponse{
				Success: false,
				Status:  "error",
				Message: "Error retrieving course",
				Errors:  utils.ErrorMessageTranslate(err),
			})
		}
		return
	}

	context.JSON(http.StatusOK, utils.SuccessResponse{
		Success: true,
		Status:  "success",
		Message: "Course retrieved successfully",
		Data:    course,
	})
}

func UpdateCourseHandler(context *gin.Context) {
	courseID := context.Param("id")
	instructorID := context.GetString("instructor_id")

	if instructorID == "" {
		context.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	var request = dto.UpdateCourseRequest{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Error validating request",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	var course models.Course

	err := db.Database.First(&course, "id = ? AND instructor_id = ?", courseID, instructorID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, utils.ErrorResponse{
				Success: false,
				Status:  "error",
				Message: "Course not found",
			})
		} else {
			context.JSON(http.StatusInternalServerError, utils.ErrorResponse{
				Success: false,
				Status:  "error",
				Message: "Error retrieving course",
				Errors:  utils.ErrorMessageTranslate(err),
			})
		}
		return
	}

	course.Title = request.Title
	course.Description = request.Description
	course.Content = request.Content
	course.Status = request.Status

	if err := db.Database.Save(&course).Error; err != nil {
		context.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Error updating course",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	context.JSON(http.StatusOK, utils.SuccessResponse{
		Success: true,
		Status:  "success",
		Message: "Course updated successfully",
		Data:    course,
	})
}

func DeleteCourseHandler(context *gin.Context) {
	courseID := context.Param("id")
	instructorID := context.GetString("instructor_id")

	if instructorID == "" {
		context.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	var course models.Course

	err := db.Database.First(&course, "id = ? AND instructor_id = ?", courseID, instructorID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, utils.ErrorResponse{
				Success: false,
				Status:  "error",
				Message: "Course not found",
			})
		} else {
			context.JSON(http.StatusInternalServerError, utils.ErrorResponse{
				Success: false,
				Status:  "error",
				Message: "Error retrieving course",
				Errors:  utils.ErrorMessageTranslate(err),
			})
		}
		return
	}

	if err := db.Database.Delete(&course).Error; err != nil {
		context.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Success: false,
			Status:  "error",
			Message: "Error deleting course",
			Errors:  utils.ErrorMessageTranslate(err),
		})
		return
	}

	context.JSON(http.StatusOK, utils.SuccessResponse{
		Success: true,
		Status:  "success",
		Message: "Course deleted successfully",
	})
}
