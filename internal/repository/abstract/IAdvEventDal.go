package abstract

import "FilterWorkerService/internal/model"

type IAdvEventDal interface {
	Add(data *model.AdvEventRespondModel) error
	GetAdvEventByCustomerId(CustomerId string, CollectionName string)(*model.BuyingEventRespondModel, error)
	UpdateAdvEventByCustomerId(CustomerId string, data *model.BuyingEventRespondModel) error
}