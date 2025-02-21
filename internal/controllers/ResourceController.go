package controllers

import (
	"API-Creation/internal/models"
	"API-Creation/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Récupérer toutes les ressources
func GetAllResources(c *gin.Context) {
	resources, err := repositories.GetAllResources()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resources)
}

// Créer une ressource
func CreateResource(c *gin.Context) {
	var resource models.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateResource(&resource); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resource)
}
