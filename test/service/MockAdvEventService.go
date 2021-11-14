package test

import "github.com/stretchr/testify/mock"

type MockAdvEventService struct {
	mock.Mock
}

func (m *MockAdvEventService) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {

}