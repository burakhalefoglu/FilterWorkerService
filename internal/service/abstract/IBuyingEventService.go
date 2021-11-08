package abstract

import "FilterWorkerService/internal/model"

type IBuyingEventService interface {
	AddBuyingEvent(data *[]byte) (s bool, m string)
	UpdateBuyingEventByCustomerId(modelResponse *model.BuyingEventRespondModel) (s bool, m string)
}
