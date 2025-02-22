package controllers

import (
	"API-Creation/internal/helpers"
	"API-Creation/internal/models"
	"API-Creation/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Récupérer toutes les alertes
func GetAllAlerts(c *gin.Context) {
	alerts, err := repositories.GetAllAlerts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alerts)
}

// Créer une alerte
func CreateAlert(c *gin.Context) {
	var alert models.Alert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateAlert(&alert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, alert)
}

// Mise à jour d'une alerte
func UpdateAlert() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var alert models.Alert
		if err := c.ShouldBindJSON(&alert); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update alert in the database
		query := `UPDATE alerts SET email = ?, isAll = ? WHERE id = ?`
		_, err := helpers.DB.Exec(query, alert.Email, alert.IsAll, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, alert)
	}
}

// Suppression d'une alerte
func DeleteAlert() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Delete alert from the database
		query := `DELETE FROM alerts WHERE id = ?`
		_, err := helpers.DB.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Alert deleted"})
	}
}
