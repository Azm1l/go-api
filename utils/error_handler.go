package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(c *gin.Context, err error) {
	var validationErrors []string

	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErr {
			switch fieldErr.Tag() {
			case "required":
				validationErrors = append(validationErrors, fmt.Sprintf("%s is required", fieldErr.Field()))
			case "email":
				validationErrors = append(validationErrors, fmt.Sprintf("%s must be a valid email", fieldErr.Field()))
			case "min":
				validationErrors = append(validationErrors, fmt.Sprintf("%s must be at least %s characters", fieldErr.Field(), fieldErr.Param()))
			default:
				validationErrors = append(validationErrors, fmt.Sprintf("%s is invalid", fieldErr.Field()))
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": validationErrors,
		})
		return
	}
	validationErrors = append(validationErrors, "Invalid JSON format")
	c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": validationErrors})
}
