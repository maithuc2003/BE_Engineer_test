package main

import (
	"tucode_v2/internal/database"
	"tucode_v2/internal/handlers"
	"tucode_v2/internal/repositories"
	"tucode_v2/internal/routes"
	"tucode_v2/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()
	robotRepo := repositories.NewRobotRepo(db)
	robotService := services.NewRobotService(robotRepo)
	robotHandler := handlers.NewRobotHandler(robotService)
	r := gin.Default()
	routes.SetRobotRoute(r, robotHandler)
	r.Run(":8080")
}
