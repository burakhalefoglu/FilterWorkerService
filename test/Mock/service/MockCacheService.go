package service

import "github.com/stretchr/testify/mock"

type MockCacheService struct {
	mock.Mock
}

func (mock *MockCacheService)ManageCache (tableName string, key string) (v int64, s bool, m string){

	args := mock.Called(tableName, key)
	return args.Get(0).(int64), args.Bool(1), args.String(2)

}