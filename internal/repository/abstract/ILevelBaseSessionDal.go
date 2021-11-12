package abstract

import "FilterWorkerService/internal/model"

type ILevelBaseSessionDal interface {
	Add(data *model.LevelBaseSessionRespondModel) error
	GetLevelBaseSessionByCustomerId(CustomerId string) (*model.LevelBaseSessionRespondModel, error)
	UpdateLevelBaseSessionByCustomerId(CustomerId string, data *model.LevelBaseSessionRespondModel) error
}
