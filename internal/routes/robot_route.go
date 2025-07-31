package routes

import (
	"tucode_v2/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetRobotRoute(router *gin.Engine, robotHandler *handlers.RobotHandler) {
	router.RouterGroup.GET("/robots", robotHandler.GetRobots)
	router.RouterGroup.POST("/robot/add", robotHandler.CreateRobot)
	router.RouterGroup.DELETE("/robot/:id", robotHandler.DeleteRobot)
	router.RouterGroup.PUT("/robot/:id", robotHandler.UpdateRobot)
	router.RouterGroup.GET("/robot/:id", robotHandler.GetRobotByID)
}
