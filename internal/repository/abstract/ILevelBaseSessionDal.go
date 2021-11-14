package abstract

import "FilterWorkerService/internal/model"

type ILevelBaseSessionDal interface {
	Add(data *model.LevelBaseSessionRespondModel) error
	GetLevelBaseSessionById(ClientId string) (*model.LevelBaseSessionRespondModel, error)
	UpdateLevelBaseSessionById(ClientId string, data *model.LevelBaseSessionRespondModel) error
}
