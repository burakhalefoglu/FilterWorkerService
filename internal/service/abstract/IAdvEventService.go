package abstract

import "FilterWorkerService/internal/model"

type IAdvEventService interface {
	AddAdvEvent(data *model.AdvEventRespondModel) (s bool, m string)
	UpdateAdvEvent(modelResponse *model.AdvEventRespondModel) (s bool, m string)
}
