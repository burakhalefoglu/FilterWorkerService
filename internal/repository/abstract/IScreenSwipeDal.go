package abstract

import "FilterWorkerService/internal/model"

type IScreenSwipeDal interface {
	Add(data *model.ScreenSwipeRespondModel) error
	GetScreenSwipeByCustomerId(CustomerId string) (*model.ScreenSwipeRespondModel, error)
	UpdateScreenSwipeByCustomerId(CustomerId string, data *model.ScreenSwipeRespondModel) error
}
