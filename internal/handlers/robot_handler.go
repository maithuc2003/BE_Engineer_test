package handlers

import (
	"net/http"
	"tucode_v2/internal/interfaces"
	"tucode_v2/internal/models"

	"github.com/gin-gonic/gin"
)

type RobotHandler struct {
	IRobotService interfaces.IRobotService
}

func NewRobotHandler(robotService interfaces.IRobotService) *RobotHandler {
	return &RobotHandler{
		IRobotService: robotService,
	}
}

func (r *RobotHandler) CreateRobot(c *gin.Context) {
	// 1. Khởi tạo cái bình trống
	var robot models.Robots
	// 2. Rót dữ liệu JSON từ HTTP body vào cái bình
	if err := c.BindJSON(&robot); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	id, err := r.IRobotService.CreateRobot(&robot)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"robot_new_id": id, "message": "Robot created successfully"})
}

func (r *RobotHandler) GetRobots(c *gin.Context) {
	var params models.RobotParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query params"})
		return
	}
	robots, err := r.IRobotService.GetRobots(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, robots)
}

func (r *RobotHandler) UpdateRobot(c *gin.Context) {
	var params models.RobotParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query params"})
		return
	}
	var robot models.Robots
	if err := c.BindJSON(&robot); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
	}
	updatedId, err := r.IRobotService.UpdateRobot(params.ID, &robot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updatedID": updatedId, "message": "Robot updated successfully"})
}
func (r *RobotHandler) DeleteRobot(c *gin.Context) {
	var params models.RobotParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query params"})
		return
	}
	deletedID, err := r.IRobotService.DeleteRobot(params.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted_id": deletedID, "message": "Robot deleted successfully"})
}

func (r *RobotHandler) GetRobotByID(c *gin.Context) {
	var params models.RobotParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query params"})
		return
	}
	robot, err := r.IRobotService.GetRobotByID(params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, robot)
}

// http://localhost:8080/robots/filter?filter=color&value=red
// http://localhost:8080/robots/filter?filter=created_at&from=2025-07-30 11:01:32&to2025-07-30 11:01:32
// http://localhost:8080/robots/filter?filter=color&value=red

// SELECT * FROM `robots` WHERE color = "red"
// SELECT * FROM robots
// WHERE created_at BETWEEN '2025-07-30 11:01:32' AND '2025-07-30 11:01:32';
