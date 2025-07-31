package interfaces

import "tucode_v2/internal/models"

type IRobotService interface {
	GetRobots(params models.RobotParams) ([]models.Robots, error)
	CreateRobot(Robot *models.Robots) (int64, error)
	UpdateRobot(RobotID int, Robot *models.Robots) (int64, error)
	DeleteRobot(RobotID int) (int64, error)
	GetRobotByID(RobotID int) (*models.Robots, error)
}
