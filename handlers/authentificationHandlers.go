package handlers

import (
	"net/http"
	"ozinshe/config"
	"ozinshe/models"
	"ozinshe/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthentificationHandler struct {
	userRepo *repositories.UsersRepository
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthentificationHandler(userRepo *repositories.UsersRepository) *AuthentificationHandler {
	return &AuthentificationHandler{userRepo: userRepo}
}

func (handler *AuthentificationHandler) SigIn(c *gin.Context) {
	var request SignInRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("could not log in"))
		return
	}
	user, err := handler.userRepo.FindByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError("could not find user"))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.NewApiError("invalid password"))
		return
	}
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(user.Id),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Config.JwtExpiresIn)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JwtSecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("could not sign token"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
