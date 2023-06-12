package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginRequest struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

const (
	InvalidReqDataMsg   = "invalid request data "
	InternalServerError = "internal server error "
	UnsuccessfulLogin   = "incorrect username or password"
)

func (h *Handler) Login(ctx *gin.Context) {
	var body LoginRequest
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": InvalidReqDataMsg + err.Error()})
	}

	if err := login(body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": InternalServerError + err.Error()})
	}
}

func login(body LoginRequest) error {
	psswd, err := hash(body.Password)
	if err != nil {
		return err
	}

	if err := checkUser(LoginRequest{
		Username: body.Username,
		Password: psswd,
	}); err != nil {
		return err
	}

	return nil
}

func checkUser(l LoginRequest) error {
	rows, err := db.Query(getUser, l.Username, l.Password)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func hash(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
