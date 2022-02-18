package abstract

import "FilterWorkerService/internal/model"

type IHardwareInformationDal interface {
	Add(data *model.HardwareResponseModel) error
	GetById(ClientId int64, ProjectId int64) (*model.HardwareResponseModel, error)
	UpdateById(ClientId int64, ProjectId int64, data *model.HardwareResponseModel) error
}