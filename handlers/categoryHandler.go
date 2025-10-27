package handlers

import (
	"net/http"
	"ozinshe/models"
	"ozinshe/repositories"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryRepo *repositories.CategoryRepository
}

func NewCategoryHandler(categoryRepo *repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{categoryRepo: categoryRepo}
}

func (handler *CategoryHandler) CreateCategory(c *gin.Context) {
	var request models.CategoryDTO
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("could bind json object"))
		return
	}
	category := models.CategoryDTO{
		FileId:     request.FileId,
		MovieCount: request.MovieCount,
		Name:       request.Name,
		Link:       request.Link,
		Id:         request.Id,
	}
	id, err := handler.categoryRepo.Create(c, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
