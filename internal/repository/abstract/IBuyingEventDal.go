package abstract

import "FilterWorkerService/internal/model"

type IBuyingEventDal interface {
	Add(data *model.BuyingEventResponseModel) error
	GetById(ClientId int64, ProjectId int64) (*model.BuyingEventResponseModel, error)
	UpdateById(ClientId int64, ProjectId int64, data *model.BuyingEventResponseModel) error
}