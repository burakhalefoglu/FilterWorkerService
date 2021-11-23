package repository

import (
	"FilterWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockLocationDal struct {
	mock.Mock
}

func (m *MockLocationDal) Add(data *model.LocationResponseModel) error{
	args:=m.Called(data)
	return args.Error(0)
}

