package repository

import (
	"FilterWorkerService/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockScreenSwipeDal struct {
	mock.Mock
}

func (m *MockScreenSwipeDal) Add(data *model.ScreenSwipeRespondModel) error{
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockScreenSwipeDal) GetScreenSwipeById(ClientId string) (*model.ScreenSwipeRespondModel, error){
	args := m.Called(ClientId)
	return args.Get(0).(*model.ScreenSwipeRespondModel), args.Error(1)
}

func (m *MockScreenSwipeDal) UpdateScreenSwipeById(ClientId string, data *model.ScreenSwipeRespondModel) error {
	args := m.Called(ClientId, data)
	return args.Error(0)
}