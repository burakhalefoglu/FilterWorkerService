package abstract

import "FilterWorkerService/internal/model"

type IGameSessionDal interface {
	Add(data *model.GameSessionResponseModel) error
	GetById(ClientId int64, ProjectId int64) (*model.GameSessionResponseModel, error)
	UpdateById(ClientId int64, ProjectId int64, data *model.GameSessionResponseModel) error
}
