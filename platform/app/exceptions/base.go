package exceptions

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewError(ctx *gin.Context, statusCode int, msg string) {
	ctx.AbortWithStatusJSON(statusCode, ErrorResponse{msg})
}
