package handlers

import (
	"net/http"
	"ozinshe/models"
	"ozinshe/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserProfileHandler struct {
	userRepo *repositories.UsersRepository
}

func NewUserProfileHandler(userRepo *repositories.UsersRepository) *UserProfileHandler {
	return &UserProfileHandler{userRepo: userRepo}
}

func (handler *UserProfileHandler) CreateUserProfile(c *gin.Context) {
	var input models.UserProfile
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind JSON"))
		return
	}

	existingUser, err := handler.userRepo.FindByEmail(c, input.User.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("User not found"))
		return
	}

	userMvp := models.UserMVP{
		Id:    existingUser.Id,
		Email: existingUser.Email,
	}

	profile := models.UserProfile{
		Name:        input.Name,
		Birthdate:   input.Birthdate,
		Language:    input.Language,
		PhoneNumber: input.PhoneNumber,
		User:        userMvp,
	}

	profileId, err := handler.userRepo.CreateProfile(c, profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("Could not create profile"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": profileId})
}

func (handler *UserProfileHandler) GetUserProfile(c *gin.Context) {
	profileId := c.Param("id")
	id, err := strconv.Atoi(profileId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid user id"))
		return
	}
	userProfile, err := handler.userRepo.FindByIdProfile(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("User not found"))
		return
	}
	c.JSON(http.StatusOK, userProfile)
}
