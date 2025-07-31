package services_test

import (
	"errors"
	"testing"
	"tucode_v2/internal/mock"
	"tucode_v2/internal/models"
	"tucode_v2/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestGetRobots(t *testing.T) {
	tests := []struct {
		name         string
		input        models.RobotParams
		expectedCall models.RobotParams
		mockResult   []models.Robots
		mockError    error
		wantResult   []models.Robots
		wantError    error
	}{
		{
			name: "Valid params",
			input: models.RobotParams{
				Page:  2,
				Limit: 5,
			},
			expectedCall: models.RobotParams{Page: 2, Limit: 5},
			mockResult: []models.Robots{
				{ID: 1, Name: "Robot A"},
				{ID: 2, Name: "Robot B"},
			},
			mockError:  nil,
			wantResult: []models.Robots{{ID: 1, Name: "Robot A"}, {ID: 2, Name: "Robot B"}},
			wantError:  nil,
		},
		{
			name: "Missing page & limit (should default)",
			input: models.RobotParams{
				Page:  0,
				Limit: 0,
			},
			expectedCall: models.RobotParams{Page: 1, Limit: 10},
			mockResult:   []models.Robots{},
			mockError:    nil,
			wantResult:   []models.Robots{},
			wantError:    nil,
		},
		{
			name: "Repo error",
			input: models.RobotParams{
				Page:  1,
				Limit: 10,
			},
			expectedCall: models.RobotParams{Page: 1, Limit: 10},
			mockResult:   nil,
			mockError:    errors.New("database failed"),
			wantResult:   nil,
			wantError:    errors.New("database failed"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mock.MockRobotRepo{
				GetListFunc: func(p models.RobotParams) ([]models.Robots, error) {
					// check defaulting logic works
					assert.Equal(t, tc.expectedCall, p)
					return tc.mockResult, tc.mockError
				},
			}

			service := &services.RobotService{
				IRobotRepo: mockRepo,
			}

			got, err := service.GetRobots(tc.input)

			assert.Equal(t, tc.wantResult, got)
			if tc.wantError != nil {
				assert.EqualError(t, err, tc.wantError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
