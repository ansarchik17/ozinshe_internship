package handlers

import (
	"net/http"
	"ozinshe/models"
	"ozinshe/repositories"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UsersHandler struct {
	repo *repositories.UsersRepository
}

func NewUserHandlers(repo *repositories.UsersRepository) *UsersHandler {
	return &UsersHandler{repo: repo}
}

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createProfile struct {
}

func (handler *UsersHandler) SignUp(c *gin.Context) {
	var request createUserRequest
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
	user := models.User{
		Email:        request.Email,
		PasswordHash: string(passwordHash),
	}
	id, err := handler.repo.Create(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("Failed to create user"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *UsersHandler) GetProfile(c *gin.Context) {

}
