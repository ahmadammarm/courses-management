package courses

import (
	"net/http"

	"github.com/ahmadammarm/courses-management/backend/db"
	"github.com/ahmadammarm/courses-management/backend/models"
	"github.com/ahmadammarm/courses-management/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetAllCoursesHandler(context *gin.Context) {
    var courses []models.Course

    database := db.Database

    database.Where("instructor_id = ?", context.GetString("instructor_id")).Find(&courses)

    context.JSON(http.StatusOK, utils.SuccessResponse{
        Success: true,
        Status:  "success",
        Message: "Courses retrieved successfully",
        Data:    courses,
    });
}