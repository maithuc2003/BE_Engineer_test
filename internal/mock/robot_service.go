package mock

import "tucode_v2/internal/models"

type MockRobotRepo struct {
	GetListFunc func(params models.RobotParams) ([]models.Robots, error)
	CreateFunc  func(Robot *models.Robots) (int64, error)
	UpdateFunc  func(RobotID int, Robot *models.Robots) (int64, error)
	DeleteFunc  func(RobotID int) (int64, error)
	GetByIDFunc func(RobotID int) (*models.Robots, error)
}

func (m *MockRobotRepo) GetList(params models.RobotParams) ([]models.Robots, error) {
	return m.GetListFunc(params)
}

func (m *MockRobotRepo) Create(robot *models.Robots) (int64, error) {
	return m.CreateFunc(robot)
}

func (m *MockRobotRepo) Update(id int, robot *models.Robots) (int64, error) {
	return m.UpdateFunc(id, robot)
}

func (m *MockRobotRepo) Delete(id int) (int64, error) {
	return m.DeleteFunc(id)
}

func (m *MockRobotRepo) GetByID(id int) (*models.Robots, error) {
	return m.GetByIDFunc(id)
}
