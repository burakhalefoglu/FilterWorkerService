package repository

import (
	"FilterWorkerService/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockBuyingEventDal struct {
	mock.Mock
}

func (m *MockBuyingEventDal) Add(data *model.BuyingEventRespondModel) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockBuyingEventDal) GetBuyingEventById(ClientId string) (*model.BuyingEventRespondModel, error) {
	args := m.Called(ClientId)
	return args.Get(0).(*model.BuyingEventRespondModel), args.Error(1)
}

func (m *MockBuyingEventDal) UpdateBuyingEventById(ClientId string, data *model.BuyingEventRespondModel) error {
	args := m.Called(ClientId, data)
	return args.Error(0)
}
