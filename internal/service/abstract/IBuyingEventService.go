package abstract

import "FilterWorkerService/internal/model"

type IBuyingEventService interface {
	ConvertRawModelToResponseModel(data *[]byte) (buying *model.BuyingEventRespondModel, s bool, m string)
}
