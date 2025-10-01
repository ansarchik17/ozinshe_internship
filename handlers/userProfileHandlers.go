package handlers

import (
	"net/http"
	"ozinshe/models"
	"ozinshe/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func (handler *UserProfileHandler) UpdateUserProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid user id"))
		return
	}
	_, err = handler.userRepo.FindByIdProfile(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("User not found"))
		return
	}
	var request models.UserProfile
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind JSON"))
		return
	}
	profile := models.UserProfile{
		Name:        request.Name,
		Birthdate:   request.Birthdate,
		Language:    request.Language,
		PhoneNumber: request.PhoneNumber,
		User:        request.User,
	}
	err = handler.userRepo.UpdateProfile(c, id, profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("User could not be updated"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func (handler *UserProfileHandler) ChangePassword(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var req models.SignUpUser
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user, err := handler.userRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("User not found"))
		return
	}

	user.PasswordHash = string(newHash)

	err = handler.userRepo.ChangePassword(c, id, string(newHash))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
