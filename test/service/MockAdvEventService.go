package test

import "github.com/stretchr/testify/mock"

type MockAdvEventService struct {
	mock.Mock
}

func (m *MockAdvEventService) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	args := m.Called(data)
	return  args.Bool(0), args.String(1)
}