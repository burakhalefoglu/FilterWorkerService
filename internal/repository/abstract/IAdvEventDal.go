package abstract

import "FilterWorkerService/internal/model"

type IAdvEventDal interface {
	Add(data *model.AdvEventRespondModel) error
	GetAdvEventById(ClientId string) (*model.AdvEventRespondModel, error)
	UpdateAdvEventById(ClientId string, data *model.AdvEventRespondModel) error
}
