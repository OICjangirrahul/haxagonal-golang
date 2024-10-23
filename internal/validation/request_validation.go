package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type APIError struct {
	Status int    `json:"status"`
	Field  string `json:"field"`
	Msg    string `json:"message"`
}

func ValidateRequest(c *gin.Context, validate *validator.Validate, model interface{}) error {
	if err := c.ShouldBindJSON(model); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			var errors []APIError
			for _, fe := range ve {
				errors = append(errors, APIError{
					Status: http.StatusBadRequest,
					Field:  fe.Field(),
					Msg:    msgForTag(fe.Tag()),
				})
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			c.Abort()
			return err
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		c.Abort()
		return err
	}

	if err := validate.Struct(model); err != nil {
		var errors []APIError
		for _, fe := range err.(validator.ValidationErrors) {
			errors = append(errors, APIError{
				Field: fe.Field(),
				Msg:   msgForTag(fe.Tag()),
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "errors": errors})
		c.Abort()
		return err
	}

	return nil
}

// 後で変更
func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required."
	case "email":
		return "Must be a valid email address."
	case "min":
		return "Must be at least the specified length."
	case "max":
		return "Cannot exceed the specified length."
	case "len":
		return "Must be exactly the specified length."
	case "alphanum":
		return "Must be alphanumeric."
	default:
		return "Invalid input."
	}
}
