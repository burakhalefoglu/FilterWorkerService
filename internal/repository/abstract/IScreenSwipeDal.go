package abstract

import "FilterWorkerService/internal/model"

type IScreenSwipeDal interface {
	Add(data *model.ScreenSwipeResponseModel) error
	GetById(ClientId int64, ProjectId int64) (*model.ScreenSwipeResponseModel, error)
	UpdateById(ClientId int64, ProjectId int64, data *model.ScreenSwipeResponseModel) error
}
