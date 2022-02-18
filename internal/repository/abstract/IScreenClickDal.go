package abstract

import "FilterWorkerService/internal/model"

type IScreenClickDal interface {
	Add(data *model.ScreenClickResponseModel) error
	GetById(ClientId int64, ProjectId int64) (*model.ScreenClickResponseModel, error)
	UpdateById(ClientId int64, ProjectId int64, data *model.ScreenClickResponseModel) error
}
