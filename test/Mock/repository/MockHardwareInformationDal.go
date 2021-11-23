package repository

import (
	"FilterWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockHardwareInformationDal struct {
	mock.Mock
}

func (m *MockHardwareInformationDal) Add(data *model.HardwareInformationResponseModel) error{
	args:=m.Called(data)
	return args.Error(0)
}