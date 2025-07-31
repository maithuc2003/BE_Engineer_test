package services

import (
	"tucode_v2/internal/interfaces"
	"tucode_v2/internal/models"
)

// các bộ phận xe, ví dụ: bánh, động cơ...
type RobotService struct {
	IRobotRepo interfaces.IRobotRepository
}

func NewRobotService(robotRepo interfaces.IRobotRepository) interfaces.IRobotService {
	return &RobotService{
		IRobotRepo: robotRepo,
	}
}

// tức là một hành động thuộc về RobotRepo.
func (r *RobotService) CreateRobot(Robot *models.Robots) (int64, error) {
	return r.IRobotRepo.Create(Robot)
}

func (r *RobotService) GetRobots(params models.RobotParams) ([]models.Robots, error) {
	// Validate page và limit
	// Gán mặc định nếu chưa có
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}
	return r.IRobotRepo.GetList(params)
}
func (r *RobotService) UpdateRobot(RobotID int, Robot *models.Robots) (int64, error) {
	return r.IRobotRepo.Update(RobotID, Robot)
}
func (r *RobotService) DeleteRobot(RobotID int) (int64, error) {
	return r.IRobotRepo.Delete(RobotID)
}
func (r *RobotService) GetRobotByID(RobotID int) (*models.Robots, error) {
	return r.IRobotRepo.GetByID(RobotID)
}
