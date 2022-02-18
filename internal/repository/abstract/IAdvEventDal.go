package abstract

import "FilterWorkerService/internal/model"

type IAdvEventDal interface {
	Add(data *model.AdvEventResponseModel) error
	GetById(ClientId int64, ProjectId int64) (*model.AdvEventResponseModel, error)
	UpdateById(ClientId int64, ProjectId int64, data *model.AdvEventResponseModel) error
}
