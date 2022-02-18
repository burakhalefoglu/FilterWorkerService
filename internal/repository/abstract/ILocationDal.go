package abstract

import "FilterWorkerService/internal/model"

type ILocationDal interface {
	Add(data *model.LocationResponseModel) error
	GetById(ClientId int64, ProjectId int64) (*model.LocationResponseModel, error)
	UpdateById(ClientId int64, ProjectId int64, data *model.LocationResponseModel) error
}
