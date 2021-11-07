package abstract

import "FilterWorkerService/internal/model"

type IAdvEventDal interface {
	Add(data *model.AdvEventRespondModel) error
	GetByCustomerId(CustomerId string, CollectionName string)(*model.BuyingEventRespondModel, error)
	UpdateByCustomerId(CustomerId string, data *model.BuyingEventRespondModel) error
}