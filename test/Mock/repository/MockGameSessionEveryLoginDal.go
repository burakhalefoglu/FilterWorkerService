package repository

import (
	"FilterWorkerService/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockGameSessionEveryLoginDal struct {
	mock.Mock
}

func (m *MockGameSessionEveryLoginDal) Add(data *model.GameSessionEveryLoginRespondModel) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockGameSessionEveryLoginDal) GetGameSessionEveryLoginById(ClientId string) (*model.GameSessionEveryLoginRespondModel, error) {
	args := m.Called(ClientId)
	return args.Get(0).(*model.GameSessionEveryLoginRespondModel), args.Error(1)
}

func (m *MockGameSessionEveryLoginDal) UpdateGameSessionEveryLoginById(ClientId string, data *model.GameSessionEveryLoginRespondModel) error{
	args := m.Called(ClientId, data)
	return args.Error(0)
}