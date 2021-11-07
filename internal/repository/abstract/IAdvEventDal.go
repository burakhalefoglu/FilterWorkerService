package abstract

import "FilterWorkerService/internal/model"

type IAdvEventDal interface {
	Add(data *model.AdvEventRespondModel) error
	GetAdvEventByCustomerId(CustomerId string)(*model.AdvEventRespondModel, error)
	UpdateAdvEventByCustomerId(CustomerId string, data *model.AdvEventRespondModel) error
}
