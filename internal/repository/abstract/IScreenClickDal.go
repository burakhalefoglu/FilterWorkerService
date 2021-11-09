package abstract

import "FilterWorkerService/internal/model"

type IScreenClickDal interface {
	Add(data *model.ScreenClickRespondModel) error
	GetScreenClickByCustomerId(CustomerId string) (*model.ScreenClickRespondModel, error)
	UpdateScreenClickByCustomerId(CustomerId string, data *model.ScreenClickRespondModel) error
}
