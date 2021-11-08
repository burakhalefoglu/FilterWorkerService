package abstract

import "FilterWorkerService/internal/model"

type IBuyingEventDal interface {
	Add(data *model.BuyingEventRespondModel) error
	GetBuyingEventByCustomerId(CustomerId string)(*model.BuyingEventRespondModel, error)
	UpdateBuyingEventByCustomerId(CustomerId string, data *model.BuyingEventRespondModel) error
}