package abstract

import "FilterWorkerService/internal/model"

type ILevelBaseSessionDal interface {
	Add(data *model.LevelBaseSessionResponseModel) error
	GetById(ClientId int64, ProjectId int64) (*model.LevelBaseSessionResponseModel, error)
	UpdateById(ClientId int64, ProjectId int64, data *model.LevelBaseSessionResponseModel) error
}
