package abstract

import "FilterWorkerService/internal/model"

type IBuyingEventDal interface {
	Add(data *model.BuyingEventRespondModel) error
	GetByCustomerId(CustomerId string)(*model.BuyingEventRespondModel, error)
	UpdateByCustomerId(CustomerId string, data *model.BuyingEventRespondModel) error
}