package abstract

import "FilterWorkerService/internal/model"

type IGameSessionEveryLoginDal interface {
	Add(data *model.GameSessionEveryLoginRespondModel) error
	GetGameSessionEveryLoginById(ClientId string) (*model.GameSessionEveryLoginRespondModel, error)
	UpdateGameSessionEveryLoginById(ClientId string, data *model.GameSessionEveryLoginRespondModel) error
}
