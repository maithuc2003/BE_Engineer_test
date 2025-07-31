package interfaces

import "tucode_v2/internal/models"

type IRobotRepository interface {
	Create(Robot *models.Robots) (int64, error)
	Update(RobotID int, Robot *models.Robots) (int64, error)
	Delete(RobotID int) (int64, error)
	GetByID(RobotID int) (*models.Robots, error)
	GetList(params models.RobotParams) ([]models.Robots, error)
}

// CreateRobot
// ReadRobotPagination
// UpdateRobot
// DeleteRobot
