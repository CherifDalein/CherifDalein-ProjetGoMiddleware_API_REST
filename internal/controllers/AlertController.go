package controllers

import (
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
