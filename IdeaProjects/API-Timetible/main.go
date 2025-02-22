package main

import (
	"API-Timetible/helpers"
	"API-Timetible/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	helpers.InitDB() // Initialisation de la DB

	r := gin.Default()
	r.POST("/events", controllers.CreateEvent)
	r.GET("/events", controllers.GetAllEvents)
	r.Run(":8080")
}
