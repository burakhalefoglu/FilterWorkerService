package repository

import (
	"FilterWorkerService/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockScreenClickDal struct {
	mock.Mock
}

func (m *MockScreenClickDal) Add(data *model.ScreenClickRespondModel) error{
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockScreenClickDal) GetScreenClickById(ClientId string) (*model.ScreenClickRespondModel, error){
	args := m.Called(ClientId)
	return args.Get(0).(*model.ScreenClickRespondModel), args.Error(1)
}

func (m *MockScreenClickDal) UpdateScreenClickById(ClientId string, data *model.ScreenClickRespondModel) error{
	args := m.Called(ClientId, data)
	return args.Error(0)
}