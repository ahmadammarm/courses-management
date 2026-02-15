package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorMessageTranslate(err error) map[string]string {
	errorMaps := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				errorMaps[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorMaps[field] = fmt.Sprintf("%s must be a valid email", field)
			case "min":
				errorMaps[field] = fmt.Sprintf("%s must be at least %s characters long", field, fieldError.Param())
			case "max":
				errorMaps[field] = fmt.Sprintf("%s must be at most %s characters long", field, fieldError.Param())
			case "unique":
				errorMaps[field] = fmt.Sprintf("%s must be unique", field)
			case "numeric":
				errorMaps[field] = fmt.Sprintf("%s must be a number", field)
			default:
				errorMaps[field] = fmt.Sprintf("%s is not valid", field)
			}
		}
	}

    // gorm errors handling
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errorMaps["Username"] = "Username already exists"
			}
			if strings.Contains(err.Error(), "email") {
				errorMaps["Email"] = "Email already exists"
			}
		} else if err == gorm.ErrRecordNotFound {
			errorMaps["Error"] = "Record not found"
		}
	}

	return errorMaps

}
