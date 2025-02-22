package controllers

import (
	"API-Creation/internal/helpers"
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

// Mise à jour d'une ressource
func UpdateResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var resource models.Resource
		if err := c.ShouldBindJSON(&resource); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update resource in the database
		query := `UPDATE resources SET name = ?, ucaID = ? WHERE id = ?`
		_, err := helpers.DB.Exec(query, resource.Name, resource.UcaID, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resource)
	}
}

// Suppression d'une ressource
func DeleteResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Delete resource from the database
		query := `DELETE FROM resources WHERE id = ?`
		_, err := helpers.DB.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Resource deleted"})
	}
}
