package main

import (
	"API-Creation/internal/controllers"
	"API-Creation/internal/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialiser la base de données
	helpers.InitDB()

	// Créer un routeur Gin
	r := gin.Default()

	// Routes pour les ressources
	r.POST("/resources", controllers.CreateResource) // Correction ici
	r.GET("/resources", controllers.GetAllResources)

	// Routes pour les alertes
	r.POST("/alerts", controllers.CreateAlert) // Correction ici
	r.GET("/alerts", controllers.GetAllAlerts) // Correction ici

	// Lancer le serveur sur le port 8080
	r.Run(":8080")
}
