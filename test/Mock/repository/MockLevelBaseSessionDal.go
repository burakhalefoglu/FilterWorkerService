package repository

import (
	"FilterWorkerService/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockLevelBaseSessionDal struct {
	mock.Mock
}

func (m *MockLevelBaseSessionDal) Add(data *model.LevelBaseSessionRespondModel) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockLevelBaseSessionDal) GetLevelBaseSessionById(ClientId string) (*model.LevelBaseSessionRespondModel, error) {
	args := m.Called(ClientId)
	return args.Get(0).(*model.LevelBaseSessionRespondModel), args.Error(1)
}

func (m *MockLevelBaseSessionDal) UpdateLevelBaseSessionById(ClientId string, data *model.LevelBaseSessionRespondModel) error {
	args := m.Called(ClientId, data)
	return args.Error(0)
}