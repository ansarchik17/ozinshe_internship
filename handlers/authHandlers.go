package handlers

import (
	"net/http"
	"ozinshe/config"
	"ozinshe/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (handler *AuthHandler) SignIn(c *gin.Context) {
	var request SignInRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could bind json object"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(request.Password), []byte(request.Email))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.NewApiError("Invalid email or password"))
		return
	}
	claims := jwt.RegisteredClaims{
		Subject:   request.Email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Config.JwtExpiresIn)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JwtSecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("Error signing token"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"token:": tokenString})
}
