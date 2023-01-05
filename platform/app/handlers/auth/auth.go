package auth

import (
	"github.com/Imomali1/platform/platform/app/exceptions"
	models "github.com/Imomali1/platform/platform/app/models/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var req models.SignUpRequest
	if err := ctx.BindJSON(&req); err != nil {
		exceptions.NewError(ctx, http.StatusUnprocessableEntity, err.Error())
	}
}

func (h *Handler) SignIn(ctx *gin.Context) {

}
