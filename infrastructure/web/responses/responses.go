package responses

import (
	"codifin-challenge/domain/utils"
	"github.com/gin-gonic/gin"
)

type SuccessDTO struct {
	Message interface{} `json:"message"`
}

type ErrorDTO struct {
	Message      string `json:"message"`
	ErrorMessage string `json:"errorMessage"`
}

func SendError(c *gin.Context, err *utils.DBError) {
	c.JSON(err.Code, newErrorResponse(err.UserMessage, err.DevelopMessage.Error()))
}

func SendSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

func newErrorResponse(message, errorMessage string) *ErrorDTO {
	return &ErrorDTO{Message: message, ErrorMessage: errorMessage}
}
