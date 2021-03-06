package repository

import (
	"FilterWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockAdvEventDal struct {
	mock.Mock
}

func (m *MockAdvEventDal) Add(data *model.AdvEventRespondModel) error{
	args:=m.Called(data)
	return args.Error(0)
}

func (m *MockAdvEventDal) GetAdvEventById(CustomerId string)(*model.AdvEventRespondModel, error){
	args:=m.Called(CustomerId)
	return args.Get(0).(*model.AdvEventRespondModel), args.Error(1)
}

func (m *MockAdvEventDal) UpdateAdvEventById(CustomerId string, data *model.AdvEventRespondModel) error{
	args:=m.Called(CustomerId, data)
	return args.Error(0)
}