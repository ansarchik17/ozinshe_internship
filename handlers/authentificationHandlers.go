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

//type createUserRequest struct {
//	Email    string `json:"email"`
//	Password string `json:"password"`
//}
//
//type SignInRequest struct {
//	Email    string `json:"email"`
//	Password string `json:"password"`
//}

func NewAuthentificationHandler(userRepo *repositories.UsersRepository) *AuthentificationHandler {
	return &AuthentificationHandler{userRepo: userRepo}
}

func (handler *AuthentificationHandler) SignUp(c *gin.Context) {
	var request models.SignUpUser
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid payload"))
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("Failed to hash password"))
		return
	}
	user := models.SignUpUser{
		Email:        request.Email,
		PasswordHash: string(passwordHash),
	}
	id, err := handler.userRepo.Create(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("Failed to create user"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *AuthentificationHandler) SigIn(c *gin.Context) {
	var request models.SignInRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("could not log in"))
		return
	}
	SignUpUser, err := handler.userRepo.FindByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError("could not find user"))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(SignUpUser.PasswordHash), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.NewApiError("invalid password"))
		return
	}
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(SignUpUser.Id),
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

func (handler *AuthentificationHandler) Logout(c *gin.Context) {
	c.Status(http.StatusOK)
}
