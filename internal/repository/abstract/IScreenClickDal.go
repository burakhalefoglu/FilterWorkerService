package abstract

import "FilterWorkerService/internal/model"

type IScreenClickDal interface {
	Add(data *model.ScreenClickRespondModel) error
	GetScreenClickById(ClientId string) (*model.ScreenClickRespondModel, error)
	UpdateScreenClickById(ClientId string, data *model.ScreenClickRespondModel) error
}
