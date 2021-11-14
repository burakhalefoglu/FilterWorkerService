package abstract

import "FilterWorkerService/internal/model"

type IScreenSwipeDal interface {
	Add(data *model.ScreenSwipeRespondModel) error
	GetScreenSwipeById(ClientId string) (*model.ScreenSwipeRespondModel, error)
	UpdateScreenSwipeById(ClientId string, data *model.ScreenSwipeRespondModel) error
}
