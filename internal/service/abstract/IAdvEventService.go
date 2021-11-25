package abstract

import "FilterWorkerService/internal/model"

type IAdvEventService interface {
	ConvertRawModelToResponseModel(data *[]byte) (adv *model.AdvEventRespondModel, s bool, m string)
}
