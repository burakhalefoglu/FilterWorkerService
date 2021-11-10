package abstract

import "FilterWorkerService/internal/model"

type IGameSessionEveryLoginDal interface {
	Add(data *model.GameSessionEveryLoginRespondModel) error
	GetGameSessionEveryLoginByCustomerId(CustomerId string) (*model.GameSessionEveryLoginRespondModel, error)
	UpdateGameSessionEveryLoginByCustomerId(CustomerId string, data *model.GameSessionEveryLoginRespondModel) error
}
