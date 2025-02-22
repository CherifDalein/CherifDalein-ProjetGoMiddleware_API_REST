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
	r.PUT("/resources/:id", controllers.UpdateResource())    // Route pour mettre à jour une ressource
	r.DELETE("/resources/:id", controllers.DeleteResource()) // Route pour supprimer une ressource

	// Routes pour les alertes
	r.POST("/alerts", controllers.CreateAlert)         // Correction ici
	r.GET("/alerts", controllers.GetAllAlerts)         // Correction ici
	r.PUT("/alerts/:id", controllers.UpdateAlert())    // Route pour mettre à jour une alerte
	r.DELETE("/alerts/:id", controllers.DeleteAlert()) // Route pour supprimer une alerte

	// Lancer le serveur sur le port 8080
	r.Run(":8080")
}
