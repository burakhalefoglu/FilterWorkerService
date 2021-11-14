package abstract

import "FilterWorkerService/internal/model"

type IBuyingEventDal interface {
	Add(data *model.BuyingEventRespondModel) error
	GetBuyingEventById(ClientId string) (*model.BuyingEventRespondModel, error)
	UpdateBuyingEventById(ClientId string, data *model.BuyingEventRespondModel) error
}